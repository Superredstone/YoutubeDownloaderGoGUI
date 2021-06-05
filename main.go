package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"runtime"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	ytDownloader "github.com/Superredstone/youtubeDownloaderGo/pkg"
	"github.com/atotto/clipboard"
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(500)), app.Title(topLabel))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			kitchen(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

var (
	urlForm = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	outputNameForm = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	button      = new(widget.Clickable)
	pasteButton = new(widget.Clickable)
	list        = &layout.List{
		Axis: layout.Vertical,
	}
	topLabel = "Youtube Downloader Go GUI"
	urlLabel = "Youtube URL: "

	URL        string
	OutputName string
	Log        string

	ops op.Ops
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func kitchen(gtx layout.Context, th *material.Theme) layout.Dimensions {
	widgets := []layout.Widget{
		//Youtube URL form
		material.H3(th, topLabel).Layout,
		func(gtx C) D {
			e := material.Editor(th, urlForm, urlLabel)
			e.Font.Style = text.Italic
			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}

			URL = e.Editor.Text()

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			})

		},
		//Paste button
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(0))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for pasteButton.Clicked() {
							//When clicked
							clipboardText, err := clipboard.ReadAll()
							if err != nil {
								fmt.Println(err)
							}

							URL = clipboardText
							urlLabel = clipboardText
						}

						dims := material.Button(th, pasteButton, "Paste").Layout(gtx)
						pointer.CursorNameOp{Name: pointer.CursorPointer}.Add(gtx.Ops)

						return dims
					})
				}),
			)
		},
		func(gtx C) D {
			e := material.Editor(th, outputNameForm, "Output file: ")
			e.Font.Style = text.Italic
			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}

			OutputName = e.Editor.Text()

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			})
		},
		//Download Button
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(0))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for button.Clicked() {
							go func() {
								err := DownloadVideo()
								if err != nil {
									fmt.Println("Error")
								}
							}()
						}
						dims := material.Button(th, button, "Download").Layout(gtx)
						pointer.CursorNameOp{Name: pointer.CursorPointer}.Add(gtx.Ops)
						return dims
					})
				}),
			)
		},

		material.H6(th, Log).Layout,

		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx)
		}}

	return list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}

func DownloadVideo() error {
	if URL == "" {
		fmt.Println("Error, no url")
		return nil
	}

	switch runtime.GOOS {
	case "android":
		OutputName = "/storage/self/primary/Download/" + OutputName
	}

	fmt.Println("Downloading " + URL + " into " + OutputName + ".mp4")
	Log = Log + "Downloading " + URL + " into " + OutputName + ".mp4\n"

	err := ytDownloader.Download(URL, OutputName+".mp4")
	if err != nil {
		return err
	}

	if runtime.GOOS == "android" {
		Log = "Video saved in Download/" + OutputName + "\n"
	}

	fmt.Println("Download comlete.")

	return nil
}
