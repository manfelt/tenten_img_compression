package main

import (
	"fmt"
)

type Window struct {
	title string
	x, y int
	width, height int
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}

// returns an output window handle that can be passed on
func createWindow(title string, x int, y int, width int, height int) (*Window, error) {	
	fmt.Println(title, x, y, width, height)
	return &Window{title: title, x: x, y: y, width: width, height: height}, nil
}

func printWindow(w *Window) {
	fmt.Println(w.title)
	for _ = range w.width {
		fmt.Printf("_")
	}
	fmt.Printf("\n")
	for _ = range w.height {
		fmt.Printf("|")
		for _ = range (w.width-2) {
			fmt.Printf(" ")
		}
		fmt.Println("|")
	}
	for _ = range w.width {
		fmt.Printf("-")
	}
}

// TODO create window: WIP
// TODO decompression


func main() {


	o := []string{" "}
	x := []string{"█"}

	//	blank := []byte(o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o)

	a := [10]int{ 103,  38, 100, 384,  -8, -24, -24, -24, -24, -24}

	testWindow, _ := createWindow("Sprite",4,7,32,32)

	fmt.Println("Testwindow:",testWindow)
	fmt.Println("asd", a, o, x)
	printWindow(testWindow)
	//	fmt.Println("█","█","█"," "," "," ","█","█","█","█","█","█"," "," "," ","█","█","█","█","█","█","█","█","█","█","█","█","█","█","█",)
}
