## Minimal example (without GUI)

```go
package main

import "C"
import "fmt"

func init(){
  // injection func
  fmt.Println("Injected")
}

func main(){}
```
To build the **DLL** you need to make function **init()**

Everything in this function **will be** executed after **DLL** injection.

## Minimal example with GUI
#### I'm using library `github.com/AllenDang/giu`, basically easier imgui bindings for Go.

```go
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
```
You can also run this script without building **DLL**, for example if you want to test it, but build and inject dll is too long

## Full example with GUI and button
#### Using syscall and windows libraries for requesting message box.
```go
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
```
##
More **examples** you can find **[here](https://github.com/AllenDang/giu/tree/master/examples)**.
# Building
You can build **DLL** only on Windows, which is obvious
##
**Format:**
`
go build -o <filename> -buildmode=c-shared <filepath>
`

**Example:**
`
go build -o dllmain.dll -buildmode=c-shared main.go
`
