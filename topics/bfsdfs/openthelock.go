package bfsdfs

import "strconv"

func OpenLock(deadends []string, target string) int {
	deadendsMap := map[string]bool{}
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}

	queue := []string{"0000", "next"}

	seen := map[string]bool{
		"0000": true,
	}

	rotations := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == "next" {
			rotations++
			if len(queue) > 0 {
				queue = append(queue, "next")
			}
		} else if curr == target {
			return rotations
		} else if !deadendsMap[curr] {
			for i := 0; i < 4; i++ {
				for upOrDown := -1; upOrDown <= 1; upOrDown += 2 {
					currCombinationNumber := int(curr[i] - '0')
					newCombinationNumber := (currCombinationNumber + upOrDown + 10) % 10
					newCombination := curr[:i] + strconv.Itoa(newCombinationNumber) + curr[i+1:]
					if !seen[newCombination] {
						seen[newCombination] = true
						queue = append(queue, newCombination)
					}
				}
			}
		}
	}

	return -1
}
