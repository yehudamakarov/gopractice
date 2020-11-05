package recursion

func Subsets(nums []int) [][]int {
	answer := getSubsets2(nums, len(nums)-1)
	return answer
}

func getSubsets(nums []int, index int) [][]int {
	var allSubsets [][]int
	if len(nums) == index {
		allSubsets = append(allSubsets, []int{})
	} else {
		allSubsets = getSubsets(nums, index+1)
		// fmt.Printf("%v\n", allSubsets)
		item := nums[index]
		var clonesToAddItemTo [][]int
		for _, existingSubset := range allSubsets {
			newSubset := make([]int, len(existingSubset))
			copy(newSubset, existingSubset)
			newSubset = append(newSubset, item)
			clonesToAddItemTo = append(clonesToAddItemTo, newSubset)
		}
		allSubsets = append(allSubsets, clonesToAddItemTo...)
		// fmt.Printf("%v\n", allSubsets)
	}
	return allSubsets
}

func getSubsets2(nums []int, index int) [][]int {
	var allSubsets [][]int
	if index == -1 {
		allSubsets = append(allSubsets, []int{})
	} else {
		allSubsets = getSubsets2(nums, index-1)
		// fmt.Printf("%v\n", allSubsets)
		item := nums[index]
		var clonesToAddItemTo [][]int
		for _, existingSubset := range allSubsets {
			newSubset := make([]int, len(existingSubset))
			copy(newSubset, existingSubset)
			newSubset = append(newSubset, item)
			clonesToAddItemTo = append(clonesToAddItemTo, newSubset)
		}
		allSubsets = append(allSubsets, clonesToAddItemTo...)
		// fmt.Printf("%v\n", allSubsets)
	}
	return allSubsets
}
