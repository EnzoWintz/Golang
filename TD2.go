/*
//package main

import "fmt"

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
*/

/*
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
    return fmt.Sprintf("%v.ck%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
*/
/*
package main

import (
    "fmt"
    "strconv"
	"strings"
)

type IPAddr [4]byte

/*func (ip IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}*/

/*
func (ip IPAddr) String() string {
    s := make([]string, len(ip))
    for i, val := range ip {
        s[i] = strconv.Itoa(int(val))
    }
    return strings.Join(s, ".")
}


func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},

    }

    //With FMT
    for name, ip := range hosts {
        fmt.Printf("%v:%v\n", name, ip)
    }


    //With StrConv

}
*/
/*
package main

import "fmt"
//import "builtin"
import "time"


type MyError struct {
    When time.Time
    What string
}

func (err MyError) Error() string {
        return fmt.Sprintf("Error at: %v this happened: %s", err.When, err.What)
}

func run() error {
    return MyError{
        When: time.Now(),
        What: "test",
}
}

func main() {
    err := run()
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println("no error")


}
*/

/*
package main

import (
    "fmt"

)
type User struct {
    Login string
    Password string
    age int
    }

func main(){

Paul := User{"Paul", "pass123",14}

fmt.Println(Paul)

}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[{"userName": "matm", "Password": "123456"}, {"userName": "fake44", "Password": "azerty"}]`)
	type Animal struct {
        userName string `json:"userName"`
        Password string `json:"Password"`
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf(jsonBlob)
}
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	UserId   string
	UserName string
	Password string
}

// GLOBAL VAR
var usersMap = make(map[string]User)

func getJson(u *[]User, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &u)
}
func getUserById(res http.ResponseWriter, req *http.Request) {
	// query id -> ?id=something
	id := req.FormValue("id")
	if id != "" {
		fmt.Printf("Your id: %v\n", id)
		if findelem, ok := usersMap[id]; ok {

			// Set my response to a Json format
			res.Header().Set(
				"Content-Type",charset
				"application/json;=utf-8",
			)
			// return response 200 -> OK
			res.WriteHeader(http.StatusOK)

			// Convert the data to a []byte
			json_data, err := json.Marshal(findelem)
			if err != nil {
				log.Fatal(err)
			}
			// return my data
			_, err = res.Write(json_data)
			return
		}
		// 	return 404  -> not found
		res.WriteHeader(http.StatusNotFound)
	}
}

//////////////////////////////
//MAIN
/////////////////////////////
func main() {
	var (
		users []User
	)
	getJson(&users, "./users.json")
	for _, user := range users {
		usersMap[user.UserId] = user
	}
	http.HandleFunc("/", getUserById)
	http.ListenAndServe(":80", nil)
}
