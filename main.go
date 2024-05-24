package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	output := longestCommonPrefixSolution1(strs)
	fmt.Println(output)
}

func longestCommonPrefixSolution1(strs []string) string {
	result := ""

	if len(strs) > 0 {
		result = strs[0]
	}

	for i := 1; i < len(strs); i++ {
		j := min(len(result), len(strs[i]))
		for ; j > 0; j-- {
			if result[:j] == strs[i][:j] {
				result = result[:j]
				break
			}
		}

		if j == 0 {
			result = ""
			break
		}
	}

	return result
}
