package nextPermutation

func nextPermutationMaxLiu(nums []int) {
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[j] <= nums[i] {
			i--
			j--
		} else if nums[i] < nums[j] {
			break
		}
	}

	// 如果i=-1，说明序列每一位都是递减的，已经是最大序列了，此时仅需要将数组逆转即可得到最小序列
	if i == -1 {
		reverse(nums[i+1:])
		return
	}

	// 如果j是最后一位，直接交换i,j即可
	if j == len(nums)-1 {
		nums[i], nums[j] = nums[j], nums[i]
		return
	}

	stop := j
	v := nums[j]
	jj := j
	for jj < len(nums) {
		if nums[jj] > nums[i] && nums[jj] <= v {
			stop = jj
			v = nums[jj]
		}
		jj++
	}
	nums[i], nums[stop] = nums[stop], nums[i]
	reverseMaxLiu(nums[j:])
}

// 对位交换
func reverseMaxLiu(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
