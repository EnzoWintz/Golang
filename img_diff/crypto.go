/*
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	img := []string("./image_1.jpg")
	content, err := ioutil.ReadFile("tst",img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}
*/

package main

import (
	"bytes"
    "crypto/sha256"
    "fmt"
    "io/ioutil"
    "log"
)

//const CHAINEDECARACTERES = 

func Hash(path string) []byte {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    Hashimg := sha256.New()
    Hashimg.Write(content)
    return Hashimg.Sum(nil)

	
}

func main() {

    img1 := Hash("image_1.jpg")
    img2 := Hash("image_2.jpg")
    img3 := Hash("image_3.jpg")

    //fmt.Println(img1)
    //fmt.Println(img2)
    //fmt.Println(img3)

	ByteToByte1 := bytes.Compare(img1,img2)
	ByteToByte2 := bytes.Compare(img1,img3)
	ByteToByte3 := bytes.Compare(img2,img3)
	
	if ByteToByte1 != 0 && ByteToByte2 != 0 {
        fmt.Println("image_1 est unique")
    } else {
        fmt.Println("image_1 n'est pas unique")
    }

    if ByteToByte1 != 0 && ByteToByte3 != 0 {
        fmt.Println("image_2 est unique")
    } else {
        fmt.Println("image_2 n'est pas unique")
    }

    if ByteToByte2 != 0 && ByteToByte3 != 0 {
        fmt.Println("image_3 est unique")
    } else {
        fmt.Println("image_3 n'est pas unique")
    }

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

