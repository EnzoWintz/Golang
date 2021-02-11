package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go maFonction(&wg)
	wg.Add(1)
	go maFonction2(&wg)
	wg.Add(1)
	go maFonction3(&wg)

	wg.Wait()
	fmt.Println("Fin du programme")
}
func maFonction(wg *sync.WaitGroup) {
	defer wg.Done() //-1 pour le add
	fmt.Println("j'ai fini !")
}

func maFonction2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("j'ai fini deux fois !")
}

func maFonction3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("j'ai fini trois fois !")
}
