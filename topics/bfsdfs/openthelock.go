package bfsdfs

import "strconv"

func OpenLock(deadends []string, target string) int {
	// make hash map of dead ends
	deadendsMap := map[string]bool{}
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}

	// make queue
	queue := []string{"0000", "next"}

	// make seen
	seen := map[string]bool{
		"0000": true,
	}

	rotations := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == "next" {
			rotations++
			// may need to check if next item to be popped from queue is next before adding next to the end
			// todo why does this if check need to be here?
			if len(queue) > 0 && queue[0] != "next" {
				queue = append(queue, "next")
			}
		} else if curr == target {
			return rotations
		} else if !deadendsMap[curr] {
			for i := 0; i < 4; i++ {
				for j := -1; j <= 1; j += 2 {
					// get neighbors
					currNumberOfCombination := int(curr[i] - '0')
					newNumberOfCombination := (currNumberOfCombination + j + 10) % 10
					newCombination := curr[:i] + strconv.Itoa(newNumberOfCombination) + curr[i+1:]
					// handle neighbors
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
