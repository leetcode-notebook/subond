package set
/*
219 存在重复元素II
https://leetcode.cn/problems/contains-duplicate-ii/
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	res := false
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i+1; j <= k+i && j < size; j++ {
			if nums[j] == nums[i] {
				if j - i <= k {
					return true
				} else {
					break
				}
			}
		}
	}
	return res
}