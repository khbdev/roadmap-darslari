package main

import (
	"fmt"
	"sync"
)


func main(){
	ch := make(chan int, 5)


	var wg sync.WaitGroup
	wg.Add(5)


	for i := 1; i <= 5; i++ {
		go func(n int) {
		  defer wg.Done()
		  ch <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}