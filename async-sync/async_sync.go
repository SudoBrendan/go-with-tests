package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// start async tasks
	chDoThingOne := doThingOne()
	chDoThingTwo := doThingTwo()

	// wait for all async tasks to complete
	for chDoThingOne != nil || chDoThingTwo != nil {
		select {
		case err := <-chDoThingOne:
			finalizeChannel(chDoThingOne, "Thing One", err)
			chDoThingOne = nil
		case err := <-chDoThingTwo:
			finalizeChannel(chDoThingTwo, "Thing Two", err)
			chDoThingTwo = nil
		}
	}
}

func doThingOne() chan error {
	ch := make(chan error)
	go func() {
		fmt.Println("Doing Thing One...")
		time.Sleep(1 * time.Second)
		ch <- errors.New("THING ONE FAILED")
		//close(ch)
	}()
	return ch
}

func doThingTwo() chan error {
	ch := make(chan error)
	go func() {
		fmt.Println("Doing Thing Two...")
		time.Sleep(3 * time.Second)
		//close(ch)
		ch <- nil
	}()
	return ch
}

func finalizeChannel(ch chan error, name string, err error) {
	if ch != nil {
		msg := fmt.Sprintf("%q completed ", name)
		if err != nil {
			msg += fmt.Sprintf("with error: %v", err)
		} else {
			msg += "successfully!"
		}
		fmt.Println(msg)
	}
}
