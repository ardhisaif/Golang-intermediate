package src

import "fmt"

var searchResult string = ""

func SearchData() {
	var data = []string{"apel", "jeruk", "pepaya", "mangga", "anggur"}
	input := "mangga"
	// input := "durian"
	fmt.Println("mulai")
	wg.Add(1)
	go search(data, input)
	wg.Wait()
	
	if searchResult == "" {
		fmt.Printf("buah %s tidak ditemukan \n", input)
	}
	fmt.Println("selesai")
}

func search(fruits []string, input string) {
	var ch = make(chan string, 1)
	for _, fruit := range fruits {
		if fruit == input {
			wg.Add(2)
			mt.Lock()
			go add(ch, fruit)
			mt.Lock()
			go read(ch)
		}
	}
	// wg.Wait()
	fmt.Print(searchResult)
	wg.Done()
}

func add(ch chan string, fruit string) {
	defer func() {
		mt.Unlock()
		close(ch)
		wg.Done()
	}()

	ch <- fruit
}

func read(ch chan string) {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()

	chn := <-ch 
	if chn != "" {
		searchResult += chn
		fmt.Printf("yeay buah %s ditemukan \n", chn)
	}

}
