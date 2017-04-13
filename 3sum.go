package leetcode

import "sort"

func threeSum(nums []int) [][]int {
    rst := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i ++ {
        if i > 0 && nums[i] == nums[i -1 ] {
            continue
        }

        target := -nums[i]
        start := i + 1
        end := len(nums) - 1
        for start < end {
            tmp := nums[start] + nums[end]
            if tmp == target {
                // Avoid duplicate nums[start], nums[end] values
                if start == i + 1 || nums[start] != nums[start-1] {
                    rst = append(rst, []int{nums[i], nums[start], nums[end]})
                }
                start ++;
                end --;
            } else if tmp < target {
                start ++;
            } else {
                end --;
            }
        }
    }
    return rst
}
