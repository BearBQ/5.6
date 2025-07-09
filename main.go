package main

import (
	"fmt"
	"time"
)

func getPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()

	panic("OOOOps")

}
func main() {
	go getPanic()

	time.Sleep(3 * time.Second)
	fmt.Println("finish")
}
