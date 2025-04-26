package main

var globalCounter *int32 = new(int32)

type AddNewFileRequest struct {
	Username string   `json:"username"`
	File     FileInfo `json:"file"`
}

type EnableIPForFileRequest struct {
	IP           string   `json:"ip"`
	File         FileInfo `json:"file"`
	ActiveBlocks []bool   `json:"active_blocks"`
}

type FileInfo struct {
	ID        int32  `json:"id,omitempty"`
	Filename  string `json:"filename"`
	Size      int64  `json:"size"`
	BlockSize int32  `json:"blocksize"`
}

type App struct {
	FileInfoMap      map[string]FileInfo
	FileIdLocalIdMap map[int32]string
	FileBlockIPs     map[int32]map[int32][]string
}

type GetBlockRequest struct {
	ID int32  `json:"id,omitempty"` //ID файла, который мы качаем
	IP string `json:"ip"`
}

type GetBlockResponse struct {
	ID        int32  `json:"id,omitempty"` //ID файла, который мы качаем
	IP        string `json:"ip"`
	Index     int32  `json:"index"`
	BlockSize int32  `json:"blocksize"`
}
