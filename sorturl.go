package main

import (
        "fmt"
        "./utils/"
       )

func main() {
    str := utils.Convert_10_to_62(1200104500000000)
    fmt.Println(str)
    fmt.Println(utils.Convert_62_to_10(str))
}
