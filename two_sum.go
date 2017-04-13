package leetcode

func twoSum(nums []int, target int) []int {
	valueMap := map[int]int{}
	for index, value := range nums {
		if pairIndex, exist := valueMap[target-value]; exist {
			return []int{index, pairIndex}
		}
		valueMap[value] = index
	}
	return []int{-1, -1}
}
