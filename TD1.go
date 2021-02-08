//TD1
//exo1
/*
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
*/
//exo2
/*
package main
import "fmt"
func main() {
fmt.Printf( "Bad ") ;
fmt.Println("formating" )
}
*/

//exo3
package main

import (
    "flag"
    "fmt"
)

const VERSION = "1.0"

func main() {
    var version = flag.Bool("version", false, "Version prog")
    flag.Parse()
    if *version {
        fmt.Println(VERSION)
    } else {
        fmt.Println("pas de version")
    }
}