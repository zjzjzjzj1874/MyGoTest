package leetcode

// 17.电话号码组合
var numLetterMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return numLetterMap[digits]
	} else if len(digits) == 2 {
		cache, ok := numLetterMap[digits]
		if ok {
			return cache
		}
		result := twoStringListsCombinations(numLetterMap[string(digits[0])], numLetterMap[string(digits[1])])
		numLetterMap[digits] = result
		return result
	} else {
		return twoStringListsCombinations(letterCombinations(digits[:2]), letterCombinations(digits[2:]))
	}
}

func twoStringListsCombinations(l1, l2 []string) []string {
	var result []string
	for _, v1 := range l1 {
		for _, v2 := range l2 {
			result = append(result, v1+v2)
		}
	}
	return result
}
