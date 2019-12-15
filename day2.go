package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
)

func p1(input string) []int {
    var intcodes = []int{}
    for _, code := range strings.Split(input, ",") {
        x, err := strconv.Atoi(code)

        if err != nil {
            panic(err)
        }

        intcodes = append(intcodes, x)
    }

    size := len(intcodes)

    // As per AoC instructions
    intcodes[1] = 12
    intcodes[2] = 2

    for i := 0; i < size - 4; i+=4 {
        opcode := intcodes[i]
        arg1 := intcodes[i+1]
        arg2 := intcodes[i+2]
        out_pos := intcodes[i+3]

        switch opcode {
        case 1:
            intcodes[out_pos] = intcodes[arg1] + intcodes[arg2]
        case 2:
            intcodes[out_pos] = intcodes[arg1] * intcodes[arg2]
        case 99:
            break
        }
    }
    return intcodes
}

func main() {
    data, err := ioutil.ReadFile("day2.input")

    if err != nil {
        panic(err)
    }

    input := strings.ReplaceAll(string(data), "\n", "")
    fmt.Println("P1:", p1(input)[0])
}
