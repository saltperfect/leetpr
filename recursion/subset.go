package recursion

import "fmt"

func subset(input, output string, answer *[]string) {
	if len(input) == 0 {
		*answer = append(*answer, output)
		return
	}
	nextWith := output + string(input[0])
	nextWithout := output
	remain := input[1:]
	subset( remain, nextWith, answer)
	subset(remain, nextWithout, answer)
}

func wrapperSubset(input string) int {
	var answers []string
	subset(input, "", &answers)
	fmt.Println(answers)
	return len(answers)
}
