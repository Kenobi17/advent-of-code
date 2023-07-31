package one

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func DayOne() {
	caloriesTxt := "./day-1/big-input.txt"

	buffer, err := os.Open(caloriesTxt)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(buffer)
	fileScanner.Split(bufio.ScanLines)

	var calories []string

	for fileScanner.Scan() {
		calories = append(calories, fileScanner.Text())
	}

	buffer.Close()

	var totalCalories []int
	caloriesSum := int64(0)

	for _, c := range calories {
		if c == "" {
			totalCalories = append(totalCalories, int(caloriesSum))
			caloriesSum = 0
		} else {
			itemCalorie, _ := strconv.ParseInt(c, 10, 0)
			caloriesSum = caloriesSum + itemCalorie
		}
	}

	sort.Slice(totalCalories, func(a, b int) bool {
		return totalCalories[a] > totalCalories[b]
	})

	fmt.Println(totalCalories[0], totalCalories[1], totalCalories[2])

	fmt.Println(totalCalories[0] + totalCalories[1] + totalCalories[2])

}
