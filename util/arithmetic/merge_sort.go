// 归并排序 -- 归并排序的主要思想是分治法。 -- 递归调用合并(稳定的排序)
// https://zhuanlan.zhihu.com/p/36075856
// https://www.jianshu.com/p/33cffa1ce613
package arithmetic

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}

// 这里是合并两个有序的数组切片 ==>从小到大进行皇城PK
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
