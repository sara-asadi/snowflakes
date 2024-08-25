package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	snowflakes := prepareData()
	fmt.Println("start")
	dt := time.Now()
	processData(snowflakes)
	dt2 := time.Now()
	fmt.Println(dt2.Second() - dt.Second())
}

func prepareData() [][6]int {
	var snowflakes [][6]int

	for i := 0; i < 1_000_000; i++ {
		snowflake := make([]int, 6)
		for j := 0; j < 6; j++ {
			snowflake[j] = rand.Intn(9_999) + 1
		}
		snowflakes = append(snowflakes, [6]int(snowflake))
	}

	return snowflakes
}

func processData(snowflakes [][6]int) {
	categorizedData := make(map[string][][6]int)

	var wrong int = 0
	for _, snowflake := range snowflakes {
		key := makeHash(snowflake)

		val, ok := categorizedData[key]

		if ok {
			if !checkCorrectness(val[0], snowflake) {
				wrong++
			} else {
				fmt.Println(val[0], snowflake)
				return
			}
		} else {
			categorizedData[key] = [][6]int{snowflake}
		}
	}
	// fmt.Println(wrong)
}

func checkCorrectness(snowflake1 [6]int, snowflake2 [6]int) bool {

	seenPoint := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if snowflake1[i] != snowflake2[j] {
				continue
			} else {
				seenPoint = 1
				for seenPoint < 6 {
					if snowflake1[(i+seenPoint)%6] != snowflake2[(j+seenPoint)%6] {
						seenPoint = 0
						break
					} else {
						seenPoint++
					}
				}
				if seenPoint == 6 {
					return true
				}
			}
		}
	}

	for i := 0; i < 6; i++ {
		for j := 5; j >= 0; j-- {
			if snowflake1[i] != snowflake2[j] {
				continue
			} else {
				seenPoint = 1
				for seenPoint < 6 {
					if j-seenPoint >= 0 {
						if snowflake1[(i+seenPoint)%6] != snowflake2[(j-seenPoint)] {
							seenPoint = 0
							break
						} else {
							seenPoint++
						}
					} else {
						if snowflake1[(i+seenPoint)%6] != snowflake2[6+(j-seenPoint)] {
							seenPoint = 0
							break
						} else {
							seenPoint++
						}
					}
				}
				if seenPoint == 6 {
					return true
				}
			}
		}
	}
	return false
}

func makeHash(snowflake [6]int) string {
	var sum int = 0
	var multiplication int = 1
	neighborMultiplication := 1

	neighborSum := 0

	a := 0
	b := 0

	for key, val := range snowflake {
		sum += val
		multiplication *= val

		if key == 0 {
			neighborMultiplication *= (val + snowflake[5] + snowflake[key+1])
			neighborSum += (val * snowflake[5] * snowflake[key+1])

			a += ((snowflake[5] + snowflake[key+1]) * val)
			b += ((snowflake[5] * snowflake[key+1]) + val)

		} else if key == 5 {
			neighborMultiplication *= (val + snowflake[key-1] + snowflake[0])
			neighborSum += (val * snowflake[key-1] * snowflake[0])

			a += ((snowflake[key-1] * snowflake[0]) * val)
			b += ((snowflake[key-1] * snowflake[0]) + val)

		} else {

			neighborMultiplication *= (val + snowflake[key-1] + snowflake[key+1])
			neighborSum += (val * snowflake[key-1] * snowflake[key+1])

			a += ((snowflake[key-1] + snowflake[key+1]) * val)
			b += ((snowflake[key-1] * snowflake[key+1]) + val)
		}
	}
	key := strconv.Itoa(sum) + "_" + strconv.Itoa(multiplication) + "_" + strconv.Itoa(neighborMultiplication) + "_" + strconv.Itoa(neighborSum) + "_" + strconv.Itoa(a) + "_" + strconv.Itoa(b)

	return key
}
