package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
)

func parse_intcodes(input string) []int {
    var intcodes = []int{}

    for _, code := range strings.Split(input, ",") {
        x, err := strconv.Atoi(code)

        if err != nil {
            panic(err)
        }

        intcodes = append(intcodes, x)
    }

    return intcodes
}

func p1(input string, noun int, verb int) []int {
    intcodes := parse_intcodes(input)
    size := len(intcodes)

    // As per AoC instructions
    intcodes[1] = noun
    intcodes[2] = verb

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

func p2(input string, expected int) int {
    for noun:= 0; noun <= 99; noun++ {
        for verb:= 0; verb <= 99; verb++ {
            if expected == p1(input, noun, verb)[0] {
                return 100 * noun + verb
            }
        }
    }
    return -1
}

func main() {
    data, err := ioutil.ReadFile("day2.input")

    if err != nil {
        panic(err)
    }

    input := strings.ReplaceAll(string(data), "\n", "")
    fmt.Println("P1:", p1(input, 12, 2)[0])
    fmt.Println("P2:", p2(input, 19690720))
}
