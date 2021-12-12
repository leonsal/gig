package main

import (
	"os"

	"github.com/leonsal/gig"
	dlg "github.com/leonsal/gig/dialog"
)

func init() {

	registerDemo(catWidgets, "Dialogs", NewDialogsDemo)
}

func NewDialogsDemo() gig.IWidget {

	msgDialog := dlg.NewMsgDialog("MsgDialog", "Question requesting user confirmation ....", "Yes", "No", "Retry").
		SetOnClick(func(d *dlg.MsgDialog, idx int) {
			switch idx {
			case 0:
				log.Info("MsgDialog:", "Yes clicked")
				d.Close()
			case 1:
				log.Info("MsgDialog:", "No clicked")
			case 2:
				log.Info("MsgDialog:", "Retry clicked")
				d.Close()
			}
		})

	inputDialog := dlg.NewInputDialog("InputDialog", "Prompt:", "OK", "Maybe", "Cancel").
		SetOnEnter(func(d *dlg.InputDialog) {
			log.Infof("Input:", d.Text())
		}).
		SetText("initial text").
		SetOnClick(func(d *dlg.InputDialog, idx int) {
			switch idx {
			case 0:
				log.Info("InputDialog:", "OK:", d.Text())
				d.Close()
			case 1:
				log.Info("InputDialog:", "Maybe clicked")
			case 2:
				log.Info("InputDialog:", "Cancel clicked")
				d.Close()
			}
		})

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileDialog, err := dlg.NewFileDialog("FileDialog", path, "")
	if err != nil {
		panic(err)
	}
	//fileDialog.SetCanClose(true)
	fileDialog.AddPath("/")
	fileDialog.ShowNewFolder("New folder")
	fileDialog.SetOnSelect(func(d *dlg.FileDialog) {
		log.Info("FileDialog selected:", d.Filepath())
		fileDialog.Close()
	}).SetOnCancel(func(d *dlg.FileDialog) {
		fileDialog.Close()
	})

	return gig.NewGroup().AddChildren(
		msgDialog,
		inputDialog,
		fileDialog,
		gig.NewButton("MsgDialog").
			SetOnClick(func(b *gig.Button) {
				msgDialog.Open()
			}),
		gig.NewButton("InputDialog").
			SetOnClick(func(b *gig.Button) {
				inputDialog.Open()
			}),
		gig.NewButton("FileDialog").
			SetOnClick(func(b *gig.Button) {
				fileDialog.SetFile("")
				fileDialog.Open()
			}),
	)
}
