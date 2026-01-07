package main

import (
	"fmt"
	"sync"
	"time"
)



func main(){
	var wg sync.WaitGroup


	for i := 1; i <= 5; i++ {
		wg.Add(1)
	go func(id int) {
		defer wg.Done()
		fmt.Println("Gorutina", id, "Started")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Gorutina", id, "Finished")
	}(i)
	}
	wg.Wait()
	fmt.Println("All gorutina Finished")
}