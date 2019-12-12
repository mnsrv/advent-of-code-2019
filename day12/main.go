package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	steps := 1000
	for step := 1; step <= steps; step++ {
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
	}

	// calculate energy
	energy := 0.0
	for _, moon := range moons {
		pot := math.Abs(float64(moon.pos.x)) + math.Abs(float64(moon.pos.y)) + math.Abs(float64(moon.pos.z))
		kin := math.Abs(float64(moon.vel.x)) + math.Abs(float64(moon.vel.y)) + math.Abs(float64(moon.vel.z))
		energy += pot * kin
	}
	fmt.Println("part one:", energy) // 8310
}
