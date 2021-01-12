package main

import (
	"fmt"
	"math"
	"sort"
)

// There are 2N people a company is planning to interview. The cost of
// flying the i-th person to city A is costs[i][0], and the cost of flying
// the i-th person to city B is costs[i][1].
//
// Return the minimum cost to fly every person to a city such that exactly N
// people arrive in each city.
//
// Example:
//
// Input: [[10,20],[30,200],[400,50],[30,20]]
//
// Output: 110
//
// Explanation:
//
// The first person goes to city A for a cost of 10.
// The second person goes to city A for a cost of 30.
// The third person goes to city B for a cost of 50.
// The fourth person goes to city B for a cost of 20.
//
// The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people
// interviewing in each city.
func main() {
	costs := [][]int{
		{10, 20},
		{30, 100},
		{29, 200},
		{10, 50},
	}

	fmt.Println(findMinCost(costs))
}

func findMinCost(costs [][]int) int {
	findDifference(costs)

	sortDescByDifference(costs)

	return getTotalMinimumCost(costs)
}

func sortDescByDifference(costs [][]int) {
	// sort slice desc by difference
	sort.Slice(costs[:], func(i, j int) bool {
		for range costs[i] {
			if costs[i][2] == costs[j][2] {
				continue
			}
			return costs[i][2] > costs[j][2]
		}
		return false
	})
}

func findDifference(costs [][]int) {
	// add a new entry to each one indicating the difference between costs
	for i, j := range costs {
		dif := int(math.Abs(float64(j[0] - j[1])))
		j = append(j, dif)
		costs[i] = j
	}
}

func getTotalMinimumCost(costs [][]int) int {
	cityACount := 0
	cityBCount := 0
	total := 0
	perCity := len(costs) / 2
	for _, cost := range costs {
		if (cost[0] <= cost[1] && cityACount < perCity) || (cityBCount == perCity) {
			total += cost[0]
			cityACount++
		} else {
			total += cost[1]
			cityBCount++
		}
	}
	return total
}