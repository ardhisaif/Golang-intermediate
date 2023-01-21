package src

import (
	"fmt"
)

func Deret() {
	var q int = 40
	deret := DeretBilangan{Number: &q}
	fmt.Println("masuk deret")

	wg.Add(4)

	// deret.evenNumber()
	go deret.evenNumber()
	// deret.oddNumber()
	go deret.oddNumber()
	// deret.fibonacciNumber()
	go deret.fibonacciNumber()
	// deret.primeNumber()
	go deret.primeNumber()

	wg.Wait()

	fmt.Println("selesai")
}

type DeretBilangan struct {
	Number *int
}

func (deret *DeretBilangan) evenNumber(){
	defer wg.Done()
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i%2 == 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
}

func (deret *DeretBilangan) oddNumber(){
	defer wg.Done()
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i%2 != 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
}

func (deret *DeretBilangan) primeNumber(){
	defer wg.Done()
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i == 1 {
			continue
		}

		checkPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				checkPrime = false
				break
			}
		}

		if checkPrime {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
}

func (deret *DeretBilangan) fibonacciNumber(){
	defer wg.Done()
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i == 1 {
			arr = append(arr, 0)
		} else if i == 2 || i == 3 {
			arr = append(arr, 1)
		} else {
			num := arr[len(arr)-1] + arr[len(arr)-2]
			if num < *deret.Number {
				arr = append(arr, num)
			}
		}
	}
	fmt.Println(arr)
}
