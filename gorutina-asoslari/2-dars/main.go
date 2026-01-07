package main

import "fmt"


func main(){
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()


	for  val := range ch {
		fmt.Println(val)
	}
}

