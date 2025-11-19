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

func toInt(dat []byte) [10]int {
        intArr := [10]int{}
        n := 0
        tmp := ""

        // 32 SP
        // 44 ,
        // 45 -
        // 48-58 0-9

        for i := range dat {
                switch dat[i] {

                case COMMA:
                        // increment array index.
                        fmt.Println(dat[i])
                        intArr[n], _ = strconv.Atoi(tmp)
                        /* if err!= nil {
                                fmt.Println("Error: ASCII to integer failed")
                        } */
                        tmp = ""
                        n++
                case SP:
                        // nothing happens at space, error in the future, maybe?
                        fmt.Println("space")
                case HYPHEN:
                        // negative number
                        fmt.Println("hyphen")
                        tmp = tmp+string(dat[i])

                default:
                        if dat[i]>=ZERO {
                                fmt.Println("number", dat[i])
                                tmp = tmp+string(dat[i])
                        } else {
                                // might have to do some error handling around here.
                                fmt.Println("Unexpected character: ", dat[i])
                        }
                }
                if tmp != "" {
                        // get that last number
                        intArr[n], _ = strconv.Atoi(tmp)
                }
        }

        // Expect n amount of commas
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

        s, err := strconv.Atoi(string(dat[0]))

        if err!= nil {
                fmt.Println("Error: ASCII to integer failed")
        }

        intArr := toInt(dat)
        fmt.Println(intArr)

        fmt.Println("val of int s:", s)

        //number := binary.LittleEndian.Uint16(dat)
        fmt.Printf("Parsed to int: %v\n", dat[21])

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

}
                        
