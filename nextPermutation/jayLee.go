package nextPermutation

func nextPermutation(nums []int) {
	// 从低位向高位遍历，找到升序的最后一位和开始降序的第一位(i,j)，从[j, len(nums)-1]中找出比nums[i]大并且与nums[i]距离最近的元素
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

	// 遍历[j, len-1]，找出比num[i]大的最少的那个数的索引
	closest := j
	v := nums[j]
	jj := j
	for jj < len(nums) {
		if nums[jj] > nums[i] && nums[jj] <= v { // nums[jj] <= v 注意：这里一定要是 '<=' 而不是 '<'
			closest = jj
			v = nums[jj]
		}
		jj++
	}
	nums[i], nums[closest] = nums[closest], nums[i] // 得到了新的[j, len-1]序列，由于该序列是大到小排序的，因此还不是我们想要的最终序列。
	reverse(nums[j:])                               // 注意：需要再将[j, len-1]按从小到大排序，就得到了最终的序列；
}

// 对位交换
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
