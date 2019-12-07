package util

func ConvertByteSliceToIntSlice(slice []byte) []int {
	res := make([]int, len(slice))

	for idx, v := range slice {
		res[idx] = int(v)
	}
	return res
}

func ConvertIntSliceToByteSlice(slice []int) []byte {
	res := make([]byte, len(slice))

	for idx, v := range slice {
		res[idx] = byte(v)
	}
	return res
}

func TestBit(value uint, bit uint) bool {
	return (value)&bit != 0
}
