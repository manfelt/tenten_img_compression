package main

import (
	//"bufio"
	//"encoding/binary"
	"fmt"
	"os"
	"strconv"
)


const (
	SP = 32
	COMMA = 44
	HYPHEN = 45	
)

const (
	ZERO = iota + 48
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
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

func printWindow(w *Window, a [99]string) {
	fmt.Println(w.title)
	const o = " "
	for _ = range w.width {
		fmt.Printf("_")
	}
	fmt.Printf("\n")
	//for _ = range w.height {
	//fmt.Printf("|")
	for i := range 10 {
		fmt.Println("|",a[i],"|")
		fmt.Println("|",a[i],"|")
	} 
	//fmt.Printf("%s %4s",o,o)
	//}
	for _ = range w.width {
		fmt.Printf("-")
	}
}

func refreshSprite(w *Window, spriteArray []string) {
	a := [99]string{"","","","","","","","","",""}
	for i := 0; i < 100; i += 10 {
		for j := 0; j < 10; j++ {
			a[j] += spriteArray[i+j]
			a[j] += spriteArray[i+j]
			a[j] += spriteArray[i+j]
		}
	}
	printWindow(w,a)
	//for i := range 10 {
	//printWindow(w,a[i])
	//fmt.Println("|",a[i],"|")
	//fmt.Println("|",a[i],"|")

	//fmt.Printf("|%s|",a[i])
	//}

	//fmt.Println("arr: ",a)
}

func toInt(data []byte) []int {
	intArr := []int{}

	//for i := range data {
	//fmt.Printf("%v ",data[i])
	//switch data {
	//case COMMA:
	//fmt.Println(COMMA)
	//}
	//}


	return intArr
}

// TODO decompression WIP


func main() {

	const o = " "
	const x = "█"

	blank := []string{o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o}
	filled := []string{x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x}
	testArr := []string{x,x,x,x,x,x,x,o,o,o,o,x,x,x,x,x,o,o,o,o,o,o,x,x,x,x,x,o,o,o,o,o,o,x,x,o,o,x,x,o,o,o,o,o,x,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o,o}

	dat, err := os.ReadFile("./data/399.tt")
	check(err)
	fmt.Print(string(dat))

	// 32 SP
	// 44 ,
	// 45 -
	// 48-58 0-9

	s, err := strconv.Atoi(string(dat[0]))
	if err!= nil {
		fmt.Println("Error: ASCII to integer failed")
	}

	// TODO ASCII to int array function

	for i := range dat {
		switch dat[i] {
		case COMMA:
			fmt.Println(dat[i])	
		case SP:
			fmt.Println("space")
		case HYPHEN:
			fmt.Println("hyphen")
		default:
			if dat[i]>=ZERO {
				fmt.Println("number")
			}
			fmt.Println("Unexpected character: ", dat[i])
		}
	}
	fmt.Println("val of int s:", s)

	//number := binary.LittleEndian.Uint16(dat)
	fmt.Printf("Parsed to int: %v\n", dat[21])

	//var data = int(f) 	
	//fmt.Println(data)


	compressed_a := [10]int{ 103,  38, 100, 384,  -8, -24, -24, -24, -24, -24}

	testWindow, _ := createWindow("Sprite",4,7,34,32)

	fmt.Println("Testwindow:",testWindow)
	fmt.Println("asd", compressed_a, o, x)
	fmt.Println("blank: ",blank)
	fmt.Println("filled: ",filled)
	fmt.Println("test array: ",testArr)
	//printWindow(testWindow)
	fmt.Printf("\n")

	refreshSprite(testWindow, testArr)

	//	fmt.Println(testArr[1+2])
	//	fmt.Println("█","█","█"," "," "," ","█","█","█","█","█","█"," "," "," ","█","█","█","█","█","█","█","█","█","█","█","█","█","█","█",)
}
