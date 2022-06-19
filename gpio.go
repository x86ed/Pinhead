package main

import (
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

func Left(d bool) {
	offset1 := rpi.J8p31
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
	if err != nil {
		panic(err)
	}
	defer func() {
		l1.Reconfigure(gpiod.AsInput)
		l1.Close()
	}()
	out := 1
	if d {
		out = 0
	}
	l1.SetValue(out)
	time.Sleep(time.Millisecond * 17)
}

func Right(d bool) {
	offset1 := rpi.J8p37
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	l1.Reconfigure(gpiod.AsInput)
	// 	l1.Close()
	// }()
	out := 1
	if d {
		out = 0
	}
	l1.SetValue(out)
	time.Sleep(time.Millisecond * 17)
}

func Start() {
	offset1 := rpi.J8p35
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
	if err != nil {
		panic(err)
	}
	defer func() {
		l1.Reconfigure(gpiod.AsInput)
		l1.Close()
	}()
	l1.SetValue(1)
	time.Sleep(time.Millisecond * 100)
	l1.SetValue(0)
}

func Launch() {
	offset1 := rpi.J8p33
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
	if err != nil {
		panic(err)
	}
	defer func() {
		l1.Reconfigure(gpiod.AsInput)
		l1.Close()
	}()
	l1.SetValue(1)
	time.Sleep(time.Millisecond * 100)
	l1.SetValue(0)
}
