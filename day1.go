package main

import (
	"fmt"
)

func fuel(mass int) int {
    return mass / 3 - 2
}

// ok, I get it, golang is not very keen
// of functional programming
func _fuel_from_fuel(f int, acc int) int {
    new_cost := fuel(f)

	if new_cost <= 0 {
		return acc
    }

	return _fuel_from_fuel(new_cost, acc + new_cost)
}

func fuel_cost(f int) int {
	return _fuel_from_fuel(f, 0)
}

func p1(modules []int) {
    total_fuel := 0
    for _, mass := range modules {
        total_fuel += fuel(mass)
    }
    fmt.Printf("p1: %v\n", total_fuel)
}

func p2(modules []int) {
    total_fuel := 0
    for _, mass := range modules {
		total_fuel += fuel_cost(mass)
	}
    fmt.Printf("p2: %v\n", total_fuel)
}

func main() {
	modules_m := [...]int{143754,
		83242,
		124730,
		62796,
		128187,
		68925,
		60687,
		68800,
		112450,
		70696,
		94653,
		62124,
		82251,
		91514,
		79895,
		82973,
		71678,
		141671,
		88243,
		109553,
		135097,
		78026,
		100048,
		52113,
		109934,
		92274,
		62821,
		138384,
		90112,
		114684,
		137383,
		71727,
		143236,
		79842,
		101187,
		71202,
		131156,
		128805,
		105102,
		71319,
		88615,
		62024,
		126027,
		55321,
		91226,
		75020,
		136689,
		70265,
		97850,
		96536,
		135311,
		64962,
		87137,
		50402,
		70604,
		56879,
		60016,
		98231,
		136635,
		64590,
		143522,
		112152,
		142511,
		95350,
		83483,
		123681,
		123792,
		99044,
		139282,
		96610,
		116844,
		50416,
		110682,
		55137,
		69795,
		100411,
		110119,
		141558,
		90780,
		108063,
		102247,
		85487,
		107174,
		79009,
		131908,
		95164,
		120588,
		62031,
		51070,
		63773,
		128565,
		96458,
		91388,
		54345,
		52840,
		130519,
		51357,
		146851,
		68455,
		102463,
	}
    p1(modules_m[:])
    p2(modules_m[:])
}
