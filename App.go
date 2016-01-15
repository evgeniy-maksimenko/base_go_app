package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTimer(time.Millisecond * 100)
	go func(){
		for t := range ticker.C{
			fmt.Println("Tick at", t)
		}
	} ()


	time.Sleep(time.Millisecond * 500)

	ticker.Stop()
	fmt.Println("Ticker stopped")

}
