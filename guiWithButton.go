package main

import "C"
import (
    "golang.org/x/sys/windows"
    "github.com/AllenDang/giu"
    "syscall"
)

func init(){
	wnd := giu.NewMasterWindow("Hello world", 200, 50, giu.MasterWindowFlagsNotResizable)
	wnd.Run(func(){
		giu.SingleWindow().Layout(
			giu.Button("Example Button").OnClick(func(){
				requestMessage("Hello from DLL Gui", "Title")
			}),
		)
	})
}

func requestMessage(message, title string) {
	windows.MessageBox(
		windows.HWND(0),
		syscall.StringToUTF16Ptr(message),
		syscall.StringToUTF16Ptr(title),
		windows.MB_OK,
	)
}

func main(){}
