package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
	z int
}
type velocity struct {
	x int
	y int
	z int
}
type moon struct {
	pos position
	vel velocity
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := strings.Split(string(data), "\n")

	moons := make([]moon, 0, 4)
	for _, input := range inputStrings {
		moon := moon{vel: velocity{x: 0, y: 0, z: 0}}
		positionArray := strings.Split(input[1:len(input)-1], ", ")
		for _, coord := range positionArray {
			key := coord[:1]
			value, _ := strconv.Atoi(coord[2:])
			switch key {
			case "x":
				moon.pos.x = value
			case "y":
				moon.pos.y = value
			case "z":
				moon.pos.z = value
			}
		}
		moons = append(moons, moon)
	}

	initialMoons := make([]moon, 4)
	copy(initialMoons, moons)
	xRepeated := 0
	yRepeated := 0
	zRepeated := 0
	steps := 0
	for {
		steps++
		// calculate velocity
		for i := range moons {
			for i2 := range moons {
				if i == i2 {
					continue
				}
				moon := moons[i]
				moon2 := moons[i2]
				if moon.pos.x < moon2.pos.x {
					moons[i].vel.x = moon.vel.x + 1
				} else if moon.pos.x > moon2.pos.x {
					moons[i].vel.x = moon.vel.x - 1
				}
				if moon.pos.y < moon2.pos.y {
					moons[i].vel.y = moon.vel.y + 1
				} else if moon.pos.y > moon2.pos.y {
					moons[i].vel.y = moon.vel.y - 1
				}
				if moon.pos.z < moon2.pos.z {
					moons[i].vel.z = moon.vel.z + 1
				} else if moon.pos.z > moon2.pos.z {
					moons[i].vel.z = moon.vel.z - 1
				}
			}
		}
		// calculate position
		for i := range moons {
			moons[i].pos.x = moons[i].pos.x + moons[i].vel.x
			moons[i].pos.y = moons[i].pos.y + moons[i].vel.y
			moons[i].pos.z = moons[i].pos.z + moons[i].vel.z
		}

		if moons[0].pos.x == initialMoons[0].pos.x && moons[0].vel.x == initialMoons[0].vel.x &&
			moons[1].pos.x == initialMoons[1].pos.x && moons[1].vel.x == initialMoons[1].vel.x &&
			moons[2].pos.x == initialMoons[2].pos.x && moons[2].vel.x == initialMoons[2].vel.x &&
			xRepeated == 0 {
			xRepeated = steps
		}
		if moons[0].pos.y == initialMoons[0].pos.y && moons[0].vel.y == initialMoons[0].vel.y &&
			moons[1].pos.y == initialMoons[1].pos.y && moons[1].vel.y == initialMoons[1].vel.y &&
			moons[2].pos.y == initialMoons[2].pos.y && moons[2].vel.y == initialMoons[2].vel.y &&
			yRepeated == 0 {
			yRepeated = steps
		}
		if moons[0].pos.z == initialMoons[0].pos.z && moons[0].vel.z == initialMoons[0].vel.z &&
			moons[1].pos.z == initialMoons[1].pos.z && moons[1].vel.z == initialMoons[1].vel.z &&
			moons[2].pos.z == initialMoons[2].pos.z && moons[2].vel.z == initialMoons[2].vel.z &&
			zRepeated == 0 {
			zRepeated = steps
		}

		if xRepeated != 0 && yRepeated != 0 && zRepeated != 0 {
			fmt.Println("part two:", LCM(xRepeated, yRepeated, zRepeated)) // 319290382980408
			break
		}
	}

	// // calculate energy
	// energy := 0.0
	// for _, moon := range moons {
	// 	pot := math.Abs(float64(moon.pos.x)) + math.Abs(float64(moon.pos.y)) + math.Abs(float64(moon.pos.z))
	// 	kin := math.Abs(float64(moon.vel.x)) + math.Abs(float64(moon.vel.y)) + math.Abs(float64(moon.vel.z))
	// 	energy += pot * kin
	// }
	// fmt.Println("part one:", energy) // 8310
}

// GCD – greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM – find Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
