package commonUtils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJsonFile(jsonPath string, createDirIfNotExist bool) (b []byte, err error) {
	// quick use tip
	// var b []byte = readJsonFile(trackListFn, true)
	// json.Unmarshal(b, &t)

	var jsonDir string
	jsonDir, _ = filepath.Split(jsonPath)
	if jsonDir == "" {
		jsonDir = "./"
	}
	_, err = os.Stat(jsonDir)
	if err != nil {
		if os.IsNotExist(err) {
			if createDirIfNotExist {
				os.MkdirAll(jsonDir, os.ModePerm)
			}
		} else {
			return
		}
	} else {
		var jsonFile *os.File
		jsonFile, err = os.Open(jsonPath)
		defer jsonFile.Close()
		if err != nil {
			if os.IsNotExist(err) {
				return
			} else {
				return
			}
		}
		b, _ = ioutil.ReadAll(jsonFile)
	}
	return
}

func SaveToJson(b []byte, jsonPath string) (err error) {
	// quick use tip
	// b, err := json.Marshal(*v)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// SaveToJson(b, varJsonPath)
	err = ioutil.WriteFile(jsonPath, b, 0644)
	if err != nil {
		return
	}
	return
}
