package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ModifyGiftFile(srcPath string, size int64, blockSize int32, srcBlockNum int32) error {
	dir, filename := filepath.Split(srcPath)
	giftPath := filepath.Join(dir, "gift_"+filename)

	giftFile, err := os.OpenFile(giftPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть или создать файл: %v", err)
	}
	defer giftFile.Close()

	fileInfo, err := giftFile.Stat()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	if fileInfo.Size() < size {
		if err := giftFile.Truncate(size); err != nil {
			return fmt.Errorf("не удалось установить размер файла: %v", err)
		}
	}

	blockData, err := ReadBlock(srcPath, srcBlockNum, blockSize)
	if err != nil {
		return fmt.Errorf("ошибка чтения блока: %v", err)
	}

	dstOffset := int64(srcBlockNum) * int64(blockSize)

	if _, err := giftFile.WriteAt(blockData, dstOffset); err != nil {
		return fmt.Errorf("ошибка записи блока: %v", err)
	}

	return nil
}

func ReadBlock(filePath string, blockNum, blockSize int32) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	offset := int64(blockNum) * int64(blockSize)

	_, err = file.Seek(offset, 0)
	if err != nil {
		return nil, fmt.Errorf("ошибка при seek: %v", err)
	}

	buffer := make([]byte, blockSize)
	bytesRead, err := file.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении: %v", err)
	}

	return buffer[:bytesRead], nil
}

func main() {
	// путь, размер файла, размер блока, номер блока для перезаписи в файл
	err := ModifyGiftFile(
		"example.bin",
		64,
		8,
		2,
	)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Операция перезаписи выполнена")
}
