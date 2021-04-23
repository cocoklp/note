算法分类：
===
比较排序：
=====
    通过比较决定元素的相对次序，时间复杂度无法突破O(nlog(n))
    交换排序：冒泡排序、快速排序
    插入排序：简单插入排序、希尔插入排序
    选择排序：简单选择排序、堆排序
    归并排序：二路归并排序、多路归并排序

非比较排序：
=====
    线性时间运行
    计数排序、桶排序、基数排序

稳定：a在b前，a==b 排序后a仍在b前

冒泡排序：
=======
    比较每一对相邻元素，前一个大于后一个则交换，以此循环最后一个是最大的数
    对于剩余元素做相同处理
    O(n2)
```
func bubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		needBreak := true
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				needBreak = false
			}
		}
		if needBreak {
			break
		}
	}
}
```

选择排序：
=======
    第一次遍历n-1个数，找到最大的/最小的与第一个元素交换
    第二次遍历n-2个数，找到最大的/最小的与第二个交换
    以此类推，在未排序的序列中找到最大/小元素，放到序列起始位置，然后在剩余的序列里继续
    O(n2)
```	
func selectSort(data []int) {
	if len(data) <= 1 {
		return
	}
	min := data[0]
	index := 0
	for i, v := range data {
		if min > v {
			min = v
			index = i
		}
	}
	data[0], data[index] = data[index], data[0]
	selectSort(data[1:])
}
```

插入排序：
=======
    构建有序序列，对于未排序的数据，在已排序的序列中从后向前扫描，找到位置后插入
```
func insertSort(data []int) {
	for i := 1; i < len(data); i++ {
		preIndex := i - 1
		cur := data[i]
		for preIndex >= 0 && data[preIndex] > cur {
			data[preIndex+1] = data[preIndex]
			preIndex--
		}
		data[preIndex+1] = cur
	}
}
```

希尔排序：
    改进版的插入排序，优先比较距离较远的元素，又叫缩小增量排序
```
func shellSort(data []int) {
	for gap := len(data) / 2; gap > 0; gap = gap >> 1 {
		for i := 0; i < len(data)-gap; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
			}
		}
	}
}
```
归并排序：
    分治。子序列有序，子序列段间有序
    长度为n的序列分成长度为n/2的子序列，对子序列归并排序，排好的子序列合并
```
// 递归
func mergeSort1(data []int) {
	fmt.Println(data)
	if len(data) < 2 {
		return
	}
	length := len(data)
	left := data[:length/2]
	right := data[length/2:]
	mergeSort1(left)
	mergeSort1(right)
	data = merge1(left, right)
}

func merge1(left, right []int) []int {
	fmt.Println(left, right)
	lindex, rindex := 0, 0
	rst := make([]int, 0, len(left)+len(right))
	for lindex < len(left) && rindex < len(right) {
		if left[lindex] < right[rindex] {
			rst = append(rst, left[lindex])
			lindex++
		} else {
			rst = append(rst, right[rindex])
			rindex++
		}
	}
	rst = append(rst, left[:lindex]...)
	rst = append(rst, right[:rindex]...)
	return rst

}

func mergeSort(data []int, first, end int) {
	if first < end {
		mid := (first + end) / 2
		mergeSort(data, first, mid)
		mergeSort(data, mid+1, end)
		merge(data, first, mid, end)
	}
}

func merge(data []int, first, mid, end int) {
	l1 := make([]int, mid-first+1)
	l2 := make([]int, end-mid)
	copy(l1, data[first:mid+1])
	copy(l2, data[mid+1:end+1])
	fmt.Println(l1, l2, data, first, mid, end)
	i, j, k := 0, 0, first
	for i < mid-first+1 && j < end-mid {
		if l1[i] < l2[j] {
			data[k] = l1[i]
			i++
		} else {
			data[k] = l2[j]
			j++
		}
		k++
	}
	if i < len(l1) {
		for ij := i; ij < len(l1); ij++ {
			data[k] = l1[ij]
			k++
		}
	}
	if j < len(l2) {
		for ij := j; ij < len(l2); ij++ {
			data[k] = l2[ij]
			k++
		}
	}
}
```
快速排序：
    1st：找到轴，元素a[0]
    2nd：i从左向右，找到第一个大于轴的值，j从右向左，找到第一个小于轴的
    3rd：交换i j 处的值
    4th: i j相遇，交换轴和i的值
    5th：左右递归 
先从数列中取出一个数作为key值；
将比这个数小的数全部放在它的左边，大于或等于它的数全部放在它的右边；
对左右两个小数列重复第二步，直至各区间只有1个数。
ref
    https://www.cnblogs.com/onepixel/p/7674659.html
```

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	privot := data[0]
	less := make([]int, 0, len(data))
	greater := make([]int, 0, len(data))
	for i := 1; i < len(data); i++ {
		if data[i] < privot {
			less = append(less, data[i])
		} else {
			greater = append(greater, data[i])
		}
	}
	left := quickSort(less)
	right := quickSort(greater)
	res := append(left, privot)
	res = append(res, right...)
	return res
}

```