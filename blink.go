package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

func Blink() {
	offset1 := rpi.J8p31
	v1 := 0
	l1, err := gpiod.RequestLine("gpiochip0", offset1, gpiod.AsOutput(v1))
	if err != nil {
		panic(err)
	}
	offset2 := rpi.J8p33
	v2 := 0
	l2, err := gpiod.RequestLine("gpiochip0", offset2, gpiod.AsOutput(v2))
	if err != nil {
		panic(err)
	}
	offset3 := rpi.J8p35
	v3 := 0
	l3, err := gpiod.RequestLine("gpiochip0", offset3, gpiod.AsOutput(v3))
	if err != nil {
		panic(err)
	}
	offset4 := rpi.J8p37
	v4 := 0
	l4, err := gpiod.RequestLine("gpiochip0", offset4, gpiod.AsOutput(v4))
	if err != nil {
		panic(err)
	}
	// revert line to input on the way out.
	defer func() {
		l1.Reconfigure(gpiod.AsInput)
		l1.Close()
		l2.Reconfigure(gpiod.AsInput)
		l2.Close()
		l3.Reconfigure(gpiod.AsInput)
		l3.Close()
		l4.Reconfigure(gpiod.AsInput)
		l4.Close()
	}()
	values := map[int]string{0: "inactive", 1: "active"}
	fmt.Printf("Set pin %d %s\n", offset1, values[v1])
	fmt.Printf("Set pin %d %s\n", offset2, values[v2])
	fmt.Printf("Set pin %d %s\n", offset3, values[v3])
	fmt.Printf("Set pin %d %s\n", offset4, values[v4])

	// capture exit signals to ensure pin is reverted to input on exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	for {
		select {
		case <-time.After(time.Second):
			v1 ^= 1
			l1.SetValue(v1)
			fmt.Printf("Set pin %d %s\n", offset1, values[v1])
			v2 ^= 1
			l2.SetValue(v2)
			fmt.Printf("Set pin %d %s\n", offset2, values[v2])
			v3 ^= 1
			l3.SetValue(v3)
			fmt.Printf("Set pin %d %s\n", offset3, values[v3])
			v4 ^= 1
			l4.SetValue(v4)
			fmt.Printf("Set pin %d %s\n", offset4, values[v4])
		case <-quit:
			return
		}
	}
}
