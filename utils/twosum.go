package utils

func TwoSum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for idx, val := range nums {
		if j, ok := resultMap[target-val]; ok {
			return []int{j, idx}
		}
		resultMap[val] = idx
	}
	return nil
}
