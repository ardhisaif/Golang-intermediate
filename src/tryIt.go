package src

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Sum(){
	fmt.Println("masuk")
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int, 2)
	wg.Add(2)
	go sum(a[:len(a)/2], chn)
	go sum(a[len(a)/2:], chn)
	// receive result from chanel and print
	var result1 = <- chn
	fmt.Println(result1)
	var result2 = <- chn
	fmt.Println(result2)

	fmt.Println("result : ", result1 + result2)
	wg.Wait()
	fmt.Println("selesai")
}

func sum(d []int, ch chan int) {
	defer func() {
		wg.Done()
	}()
	var result int = 0
	for _, v := range d {
		// hitung
		result += v
	}
	ch <- result
	// send result to result
}