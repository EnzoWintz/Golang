//TD1
//Exo1
/*
package main

import "fmt"

func main() {
	fmt.Printf("Bad ")
	fmt.Println("formating")
}
*/
/*
package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0"

func main() {
	var version = flag.Bool("version", false, "")
	flag.Parse()
	if *version {
		fmt.Println("version:", VERSION)
	} else {
		fmt.Println("pas de version")
	}
}*/
/*
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

	ByteToByte1 := bytes.Compare(img1, img2)
	ByteToByte2 := bytes.Compare(img1, img3)
	ByteToByte3 := bytes.Compare(img2, img3)

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
*/
/*
package main

import "fmt"

func main() {
	ex := Example{age: 10}
	modifAge(ex, 45)
	fmt.Println(ex)

}

type Example struct {
	age int
}

func modifAge(ex Example, age int) {
	ex.age = age
}
*/
//TD2
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var users = make(map[string]User)

func main() {
	data, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	var fileUsers []User
	err = json.Unmarshal(data, &fileUsers)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range fileUsers {
		users[u.ID] = u
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	valeurID := r.FormValue("id")

	if valeurID == "" {

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	user, found := users[valeurID]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {

		return
	}
}

type User struct {
	Login    string `json:"UserName"`
	Password string
	ID       string `json:"UserId"`
}
