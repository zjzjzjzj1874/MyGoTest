package arithmetic

// 选择排序(由小到大:也可以由大到小)
func SelectionSort(arr []int) []int {
	for i := 0;i < len(arr); i ++{
		minIndex := i
		for j := i + 1;j < len(arr);j ++{
			if arr[j] <= arr[minIndex]{
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
	return arr
}
// 冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}