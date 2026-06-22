package main

import (
	"fmt"
	"math/rand"
)

func RandGenerator(ch chan int){
	for i := 0; i < 10; i++{
		ch <- rand.Intn(10)
	}
	close(ch)
}

func main(){
	transf := make(chan int)
	go RandGenerator(transf)
	for v := range transf{
		fmt.Println(v)
	}
}
