/*package main

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