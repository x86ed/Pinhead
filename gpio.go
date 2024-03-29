package main

import (
	"fmt"
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

func Initials(abc string) {
	fmt.Println(abc)

	var lastChar rune

	// truncate strings longer than length 3
	abc = abc[:3]

	// lowercase starts at rune 97 (a) and ends at 122 (z)
	// uppercase starts at rune 65 (A) and ends at 90 (Z)
	for _, c := range abc {

		// TODO: try to attempt to coerce accented characters?

		// ensure uppercase, if possible
		if c >= 'a' && c <= 'z' {
			c -= 32
		}

		// not a valid rune
		if c-'A' < 0 || c-'A' > 25 {
			// TODO: panic here?
			// set to 'A' rune for now
			c = 'A'
		}

		// rune init zero value is 0
		if lastChar == 0 {
			lastChar = 'A'
		}

		offset := int(c - lastChar)

		fmt.Println(offset)

		if offset >= 0 {
			for i := 0; i < offset; i++ {
				RightClick()
			}
		} else if offset < 0 {
			for i := 0; i < -offset; i++ {
				LeftClick()
			}
		}

		lastChar = c
		Start()
	}

	// last start submits
	Start()
}

var leftState = 0
var rightState = 0

func Left(d bool) {
	if d == true && leftState == 2 {
		return
	}
	if d == false && leftState == 1 {
		return
	}
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
	if d == true {
		leftState = 2
		fmt.Println("⇦")
	} else {
		leftState = 1
		fmt.Println("⬅")
	}
}

func LeftClick() {
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
	l1.SetValue(1)
	time.Sleep(time.Millisecond * 100)
	l1.SetValue(0)
	fmt.Println("⬅")
}

func Right(d bool) {
	if d == true && rightState == 2 {
		return
	}
	if d == false && rightState == 1 {
		return
	}
	offset1 := rpi.J8p37
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.WithPullUp, gpiod.AsOutput(v1))
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
	if d == true {
		rightState = 2
		fmt.Println("⇨")
	} else {
		rightState = 1
		fmt.Println("⮕")
	}
}

func RightClick() {
	offset1 := rpi.J8p37
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.WithPullUp, gpiod.AsOutput(v1))
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
	fmt.Println("⮕")
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
	fmt.Println("🙋")
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
	fmt.Println("⬆")
}
