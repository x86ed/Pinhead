package main

import (
	"fmt"
)

func Initials(abc string) {
	//TODO write initial entry function
	fmt.Println(abc)

	// var lastChar rune

	// // truncate strings longer than length 3
	// abc = abc[:3]

	// // lowercase starts at rune 97 (a) and ends at 122 (z)
	// // uppercase starts at rune 65 (A) and ends at 90 (Z)
	// for _, c := range abc {

	// 	// TODO: try to attempt to coerce accented characters?

	// 	// ensure uppercase, if possible
	// 	if c >= 'a' && c <= 'z' {
	// 		c -= 32
	// 	}

	// 	// not a valid rune
	// 	if c-'A' < 0 || c-'A' > 25 {
	// 		// TODO: panic here?
	// 		// set to 'A' rune for now
	// 		c = 'A'
	// 	}

	// 	// rune init zero value is 0
	// 	if lastChar == 0 {
	// 		lastChar = 'A'
	// 	}

	// 	offset := int(c - lastChar)

	// 	fmt.Println(offset)

	// 	if offset >= 0 {
	// 		for i := 0; i < offset; i++ {
	// 			Right(true) // no need for debounce here?
	// 			Right(false)
	// 		}
	// 	} else if offset < 0 {
	// 		for i := 0; i < -offset; i++ {
	// 			Left(true) // what about here?
	// 			Left(false)
	// 		}
	// 	}

	// 	lastChar = c
	// 	Start()
	// }

	// // last start submits
	// Start()
}

// var leftState = 0
// var rightState = 0

// func Left(d bool) {
// 	if d == true && leftState == 2 {
// 		return
// 	}
// 	if d == false && leftState == 1 {
// 		return
// 	}
// 	offset1 := rpi.J8p31
// 	v1 := 0
// 	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		l1.Reconfigure(gpiod.AsInput)
// 		l1.Close()
// 	}()
// 	out := 1
// 	if d {
// 		out = 0
// 	}
// 	l1.SetValue(out)
// 	time.Sleep(time.Millisecond * 17)
// 	if d == true {
// 		leftState = 2
// 	} else {
// 		leftState = 1
// 	}
// }

// func Right(d bool) {
// 	if d == true && rightState == 2 {
// 		return
// 	}
// 	if d == false && rightState == 1 {
// 		return
// 	}
// 	offset1 := rpi.J8p37
// 	v1 := 0
// 	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.WithPullUp, gpiod.AsOutput(v1))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		l1.Reconfigure(gpiod.AsInput)
// 		l1.Close()
// 	}()
// 	out := 1
// 	if d {
// 		out = 0
// 	}
// 	l1.SetValue(out)
// 	time.Sleep(time.Millisecond * 17)
// 	if d == true {
// 		rightState = 2
// 	} else {
// 		rightState = 1
// 	}
// }

// func Start() {
// 	offset1 := rpi.J8p35
// 	v1 := 0
// 	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		l1.Reconfigure(gpiod.AsInput)
// 		l1.Close()
// 	}()
// 	l1.SetValue(1)
// 	time.Sleep(time.Millisecond * 100)
// 	l1.SetValue(0)
// }

// func Launch() {
// 	offset1 := rpi.J8p33
// 	v1 := 0
// 	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		l1.Reconfigure(gpiod.AsInput)
// 		l1.Close()
// 	}()
// 	l1.SetValue(1)
// 	time.Sleep(time.Millisecond * 100)
// 	l1.SetValue(0)
// }
