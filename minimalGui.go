package main

import "C"
import "github.com/AllenDang/giu"

func init(){
	wnd := giu.NewMasterWindow("Hello world", 400, 200, giu.MasterWindowFlagsNotResizable)
	wnd.Run(func(){
		// some code
	})
}

func main(){}
