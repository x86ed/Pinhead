package main

import (
	"fmt"
)

func Initials(abc string) {
	//TODO write initial entry function
	fmt.Println(abc)
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
