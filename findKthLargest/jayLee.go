package findKthLargest

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelectRandomizedKth(nums, 0, len(nums)-1, len(nums)-k)
}

// 以quickSearch为模型，使用单侧划分的方式获取第k大元素，并使用随机数
func quickSelectRandomizedKth(A []int, p, r int, i int) int {
	// A = [2,8,7,1,3,5,6,4] -> [2,1,3,4,...] -> q = 3
	// i = 2
	q := randomPartition(A, p, r)
	if q == i {
		return A[q]
	}
	if q < i { // 当前q在i的左侧，选择右侧部分递归
		return quickSelectRandomizedKth(A, q+1, r, i)
	}
	return quickSelectRandomizedKth(A, p, q-1, i)
}

// 以quickSearch为模型，使用单侧划分的方式获取第k大元素
func quickSelectKth(A []int, p, r int, i int) int {
	// A = [2,8,7,1,3,5,6,4] -> [2,1,3,4,...] -> q = 3
	// i = 2
	q := partition(A, p, r)
	if q == i {
		return A[q]
	}
	if q < i { // 当前q在i的左侧，选择右侧部分递归
		return quickSelectKth(A, q+1, r, i)
	}
	return quickSelectKth(A, p, q-1, i)
}

// 快速排序的代码
func quickSearch(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSearch(A, p, q-1)
		quickSearch(A, q+1, r)
	}
}

// 采用随机抽样。等概率的从r-p+1个元素中选取一个作为主元。
func randomPartition(A []int, p, r int) int {
	i := rand.Int()%(r-p+1) + p // RANDOM(p, r)的golang实现
	fmt.Println(p, r, "===", i)
	A[i], A[r] = A[r], A[i]
	fmt.Println(A)
	return partition(A, p, r)
}

// 算法的关键部分，实现了对自数组A[p..r]的原址重排
func partition(A []int, p, r int) int {
	x := A[r]  // 主元。围绕它来划分子数组A[p..r]。
	i := p - 1 // q = i+1
	// [2,8,7,1,3,5,6,4] -> [2,1,3,8,7,5,6,4]。此时：i = 2; j = r - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[j], A[i] = A[i], A[j]
		}
	}
	A[i+1], A[r] = A[r], A[i+1] // 由于之前只遍历到r-1，因此需要交换i+1和r的数据，使得q = i+1, A[q] >= A[p..q-1] && A[q] <= A[q+1..r]
	return i + 1                // 返回i+1
}
