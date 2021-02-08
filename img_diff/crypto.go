
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	img := []byte(";/image_1.jpg")
	content, err := ioutil.ReadFile("tst",img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}


/*
package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    sum := sha256.Sum256([]byte("./image_1.jpg"))
	sum2 := sha256.Sum256([]byte("./image_2.jpg"))
	sum3 := sha256.Sum256([]byte("./image_3.jpg"))

    fmt.Printf("Hachage de l'image 1 : %x \n\n", sum)
	fmt.Printf("Hachage de l'image 2 : %x \n\n", sum2)
	fmt.Printf("Hachage de l'image 3 : %x \n", sum3)
}*/

