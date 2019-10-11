package main

import (
	"fmt"
	"sync"
	"C"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//i := i
		go func () {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

//export Hello
func Hello()  {
	fmt.Println("Hello")
}

//go build --buildmode=c-archive main.go

