package main

import "fmt"

var keys = map[int]string{
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
    0: "zero",
}

var array [4] int
var str [4]string

func main() {

var digits = []int{0,1,2,3,4,5,6,7,8,9}

for i := range digits {
    fmt.Println(i)
    array[0] = i
    fmt.Println(array)

    //Loop through array and convert output to strings. 
    for k := 0; k < len(array); k++ {
      str[k] = keys[array[k]]
    }

    fmt.Println(str)
  }
}
