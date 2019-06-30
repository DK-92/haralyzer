package main

import (
	"log"
	"path/filepath"

	"github.com/DK-92/harviewer/screens/model"

	"github.com/DK-92/harviewer/screens/logic"

	_ "github.com/andlabs/ui/winmanifest"
	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

type ApplicationWindow struct {
	*walk.MainWindow
	tabWidget    *walk.TabWidget
	prevFilePath string
}

func main() {
	mw := new(ApplicationWindow)
	var openAction *walk.Action

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Haralyzer",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						OnTriggered: mw.openAction_Triggered,
					},
					Separator{},
					Action{
						Text:        "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: mw.aboutAction_Triggered,
					},
				},
			},
		},
		MinSize: Size{320, 240},
		Size:    Size{800, 600},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			TabWidget{
				AssignTo: &mw.tabWidget,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

func (mw *ApplicationWindow) openAction_Triggered() {
	if err := mw.openFile(); err != nil {
		log.Print(err)
	}
}

func (mw *ApplicationWindow) openFile() error {
	fileDialog := new(walk.FileDialog)

	fileDialog.FilePath = mw.prevFilePath
	fileDialog.Filter = "Image Files (*.har)|*.har"
	fileDialog.Title = "Select a file"

	if ok, err := fileDialog.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.prevFilePath = fileDialog.FilePath

	position := logic.LoadHarFile(mw.MainWindow, fileDialog.FilePath)

	if position != -1 {
		page, _ := walk.NewTabPage()
		page.SetTitle(filepath.Base(fileDialog.FilePath))

		mw.tabWidget.Pages().Add(page)
		mw.tabWidget.SetCurrentIndex(mw.tabWidget.Pages().Len() - 1)

		model.MakeFileContentTab(page, position, filepath.Base(fileDialog.FilePath))
	}

	return nil
}

func (mw *ApplicationWindow) aboutAction_Triggered() {
	walk.MsgBox(mw, "About", "Standalone HAR viewer for Windows.\n\nVersion 1.0 released on 20-06-2019.", walk.MsgBoxIconInformation)
}
