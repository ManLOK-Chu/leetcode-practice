package main

import (
	"fmt"
	"math"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
//算法的时间复杂度应该为 O(log (m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//二分查找
	var length1, length2 = len(nums1), len(nums2)
	if length1 > length2 { // 为了保证第一个数组比第二个数组小(或者相等)
		nums2, nums1 = nums1, nums2
		length1, length2 = length2, length1
	}
	// 分割线左边的所有元素需要满足的个数 m + (n - m + 1) / 2;
	// 两个数组长度之和为偶数时，当在长度之和上+1时，由于整除是向下取整，所以不会改变结果
	// 两个数组长度之和为奇数时，按照分割线的左边比右边多一个元素的要求，此时在长度之和上+1，就会被2整除，会在原来的数
	//的基础上+1，于是多出来的那个1就是左边比右边多出来的一个元素
	var totalLeft = (length1 + length2 + 1) / 2
	// 在 nums1 的区间 [0, length1] 里查找恰当的分割线，
	// 使得 nums1[i - 1] <= nums2[j] && nums2[j - 1] <= nums1[i]
	var left, right = 0, length1
	// nums1[i - 1] <= nums2[j]
	//  此处要求第一个数组中分割线的左边的值 不大于(小于等于) 第二个数组中分割线的右边的值
	// nums2[j - 1] <= nums1[i]
	//  此处要求第二个数组中分割线的左边的值 不大于(小于等于) 第一个数组中分割线的右边的值
	// 循环条件结束的条件为指针重合，即分割线已找到
	for left < right {
		// 二分查找，此处为取第一个数组中左右指针下标的中位数，决定起始位置
		// 此处+1首先是为了不出现死循环，即left永远小于right的情况
		// left和right最小差距是1，此时下面的计算结果如果不加1会出现i一直=left的情况，而+1之后i才会=right
		// 于是在left=i的时候可以破坏循环条件，其次下标+1还会保证下标不会越界，因为+1之后向上取整，保证了
		// i不会取到0值，即i-1不会小于0
		// 此时i也代表着在一个数组中左边的元素的个数
		i := left + (right-left+1)/2
		// 第一个数组中左边的元素个数确定后，用左边元素的总和-第一个数组中元素的总和=第二个元素中左边的元素的总和
		// 此时j就是第二个元素中左边的元素的个数
		j := totalLeft - i
		// 此处用了nums1[i - 1] <= nums2[j]的取反，当第一个数组中分割线的左边的值大于第二个数组中分割线的右边的值
		// 说明又指针应该左移，即-1
		if nums1[i-1] > nums2[j] {
			// 下一轮搜索的区间 [left, i - 1]
			right = i - 1
			// 此时说明条件满足，应当将左指针右移到i的位置，至于为什么是右移，请看i的定义
		} else {
			// 下一轮搜索的区间 [i, right]
			left = i
		}
	}
	// 退出循环时left一定等于right，所以此时等于left和right都可以
	// 为什么left一定不会大于right?因为left=i。
	// 此时i代表分割线在第一个数组中所在的位置
	// nums1[i]为第一个数组中分割线右边的第一个值
	// nums[i-1]即第一个数组中分割线左边的第一个值
	i := left
	// 此时j代表分割线在第二个数组中的位置
	// nums2[j]为第一个数组中分割线右边的第一个值
	// nums2[j-1]即第一个数组中分割线左边的第一个值
	j := totalLeft - i
	// 当i=0时，说明第一个数组分割线左边没有值，为了不影响
	// nums1[i - 1] <= nums2[j] 和 Math.max(nums1LeftMax, nums2LeftMax)
	// 的判断，所以将它设置为int的最小值
	var nums1LeftMax, nums1RightMin, nums2LeftMax, nums2RightMin int
	if i == 0 {
		nums1LeftMax = math.MinInt
	} else {
		nums1LeftMax = nums1[i-1]
	}
	// 等i=第一个数组的长度时，说明第一个数组分割线右边没有值，为了不影响
	// nums2[j - 1] <= nums1[i] 和 Math.min(nums1RightMin, nums2RightMin)
	// 的判断，所以将它设置为int的最大值
	if i == length1 {
		nums1RightMin = math.MaxInt
	} else {
		nums1RightMin = nums1[i]
	}
	// 当j=0时，说明第二个数组分割线左边没有值，为了不影响
	// nums2[j - 1] <= nums1[i] 和 Math.max(nums1LeftMax, nums2LeftMax)
	// 的判断，所以将它设置为int的最小值
	if j == 0 {
		nums2LeftMax = math.MinInt
	} else {
		nums2LeftMax = nums2[j-1]
	}
	// 等j=第二个数组的长度时，说明第二个数组分割线右边没有值，为了不影响
	// nums1[i - 1] <= nums2[j] 和 Math.min(nums1RightMin, nums2RightMin)
	// 的判断，所以将它设置为int的最大值
	if j == length2 {
		nums2RightMin = math.MaxInt
	} else {
		nums2RightMin = nums2[j]
	}
	// 如果两个数组的长度之和为奇数，直接返回两个数组在分割线左边的最大值即可
	if ((length1 + length2) % 2) == 1 {
		return float64(max(nums1LeftMax, nums2LeftMax))
	} else {
		// 如果两个数组的长度之和为偶数，返回的是两个数组在左边的最大值和两个数组在右边的最小值的和的二分之一
		// 此处不能被向下取整，所以要强制转换为double类型
		return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//归并排序: 时间复杂度O(m+n)
//空间复杂度O（m+n)
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	var i, j = 0, 0
//	var length1, length2 = len(nums1), len(nums2)
//	var n = length2 + length1
//	var nums = make([]int, 0, n)
//	//归并排序
//	for i < length1 && j < length2 {
//		var x, y = nums1[i], nums2[j]
//		if x < y {
//			nums = append(nums, x)
//			i++
//		} else {
//			nums = append(nums, y)
//			j++
//		}
//	}
//	if i == length1 && j < length2 { //nums1已完成遍历
//		nums = append(nums, nums2[j:]...)
//	} else if j == length2 && i < length1 { //nums2已完成遍历
//		nums = append(nums, nums1[i:]...)
//	}
//	if n%2 == 1 {
//		return float64(nums[n/2])
//	}
//	return float64(nums[n/2]+nums[n/2-1]) / 2.0
//}

//快排: 时间复杂度O( (m+n)log(m+n) )
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	var nums = append(nums1, nums2...)
//	sort.Ints(nums)
//	var n = len(nums)
//	if n%2 == 1 {
//		return float64(nums[n/2])
//	}
//	return float64(nums[n/2]+nums[n/2-1]) / 2.0
//}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
