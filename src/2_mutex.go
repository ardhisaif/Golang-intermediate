package src

import (
	"fmt"
	"sync"

)

var mt = sync.Mutex{}

var data = []string{"apel", "jeruk", "pepaya", "mangga", "anggur"}

func Mutex() {

	fmt.Println("mulai")
	var ch = make(chan string, len(data))
	

	for i := 0; i < len(data); i++ {
		wg.Add(2)
		mt.Lock()
		go index(i, ch)
		mt.Lock()
		go showData(ch)
	}

	wg.Wait()

	fmt.Println("selesai")

}

func index(i int, ch chan string) {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()

	ch <- data[i]
}

func showData(ch chan string){
	chn := <- ch
	fmt.Println("buah", chn)

	mt.Unlock()
	wg.Done()
}

