package recursion

import "fmt"

func permute(input, output string, answers *[]string) {
	if len(input) == 0 {
		fmt.Println(output)
		*answers = append(*answers, output)
		return
	}

	for i := 0; i < len(input); i++ {
		next := output + string(input[i])
		remain := input[:i] + input[i+1:]
		permute(remain, next, answers)
	}
}

func wrapper(input string) int {
	var answers []string
	permute(input, "", &answers)
	fmt.Println(answers)
	return len(answers)
}
