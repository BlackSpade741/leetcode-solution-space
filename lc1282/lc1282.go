package lc1282

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	// keep track of current groups in a map of size -> list of people
	// every time we add one, check if it's done
	groups := make([][]int, 0)
	currentGroups := make(map[int][]int)
	
	for i, groupSize := range(groupSizes) {
		fmt.Println(groups)
		fmt.Println(currentGroups)
		group := currentGroups[groupSize]

		group = append(group, i)
		currentGroups[groupSize] = group

		fmt.Printf("Adding %d to group, group now: %v\n", i, group)
		if len(group) == groupSize {
			fmt.Printf("Group %v is at size, add to groups and remove from currentGroups\n", group)

			groups = append(groups, group)

			delete(currentGroups, groupSize)
		}
	}

	return groups
}

func Solution() {
	fmt.Println("Leetcode Question 1282")

	groupSizes := []int{3,3,3,3,3,1,3}

	fmt.Println(groupThePeople(groupSizes))
}
