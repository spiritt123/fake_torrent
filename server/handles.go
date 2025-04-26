package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"slices"
	"sync/atomic"
)

func (app *App) GetFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(app.FileInfoMap)
}

func (app *App) AddNewFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req AddNewFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := req.Username + "_" + req.File.Filename
	_, ok := app.FileInfoMap[key]
	if ok {
		json.NewEncoder(w).Encode(map[string]string{"status": "Файл существует"})
		return
	}

	// добавили файл в список файлов
	req.File.ID = atomic.AddInt32(globalCounter, 1)
	app.FileInfoMap[key] = req.File
	app.FileIdLocalIdMap[req.File.ID] = key

	json.NewEncoder(w).Encode(map[string]string{"message": "String added successfully"})
}

func (app *App) EnableIPForFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req EnableIPForFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, isExist := range req.ActiveBlocks {
		if isExist {
			if app.FileBlockIPs[req.File.ID] == nil {
				app.FileBlockIPs[req.File.ID] = make(map[int32][]string)
			}
			if !slices.Contains(app.FileBlockIPs[req.File.ID][int32(i)], req.IP) {
				app.FileBlockIPs[req.File.ID][int32(i)] = append(app.FileBlockIPs[req.File.ID][int32(i)], req.IP)
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "IP update successfully"})
}

func (app *App) GetBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req GetBlockRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// алгоритм выборки
	file := app.FileBlockIPs[req.ID]

	minBlockIpCount := []int32{}
	min := 4294967295 // MAX_INT32
	for i, v := range file {
		if slices.Contains(v, req.IP) {
			continue
		}
		if len(v) < min {
			min = len(v)
			minBlockIpCount = []int32{}
		} else if len(v) == min {
			minBlockIpCount = append(minBlockIpCount, i)
		}
	}

	choice := rand.Intn(len(minBlockIpCount))

	index := minBlockIpCount[choice]
	choiceIp := rand.Intn(len(file[index]))

	var res GetBlockResponse
	res.BlockSize = app.FileInfoMap[app.FileIdLocalIdMap[req.ID]].BlockSize
	res.ID = req.ID
	res.Index = index
	res.IP = file[index][choiceIp]

	json.NewEncoder(w).Encode(res)
}
