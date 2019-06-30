package logic

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/lxn/walk"

	"github.com/DK-92/harviewer/structs"
)

func LoadHarFile(mw *walk.MainWindow, filePath string) int {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		showFileLoadError(mw)

		return -1
	}

	var harContents structs.MainLog

	err = json.Unmarshal(file, &harContents)
	if err != nil {
		log.Println(err)
		showFileLoadError(mw)

		return -1
	}

	data := structs.GetData()
	data.HarFiles = append(data.HarFiles, harContents)

	return len(data.HarFiles) - 1
}

func showFileLoadError(mw *walk.MainWindow) {
	walk.MsgBox(mw, "Error loading file", "An error occurred when loading the specified har file.", walk.MsgBoxIconError)
}
