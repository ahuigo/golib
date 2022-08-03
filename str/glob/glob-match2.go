package main
/* refer: https://golangbyexample.com/wildcard-matching-golang/#Dynamic_Program_Solution
other method:
    1. convert glob to regex '?a*b' -> '.a.*b'
    2. 
*/
import "fmt"
func main() {
    output := isMatch("adcb", "a*b")
	fmt.Println(output)

    /*
	output = isMatch("aa", "aa")
	fmt.Println(output)

	output = isMatch("aaaa", "*")
	fmt.Println(output)

	output = isMatch("ab", "a?")
	fmt.Println(output)

	output = isMatch("aa", "a")
	fmt.Println(output)

	output = isMatch("mississippi", "m??*ss*?i*pi")
	fmt.Println(output)

	output = isMatch("acdcb", "a*c?b")
	fmt.Println(output)

    */
}

func isMatch(s string, p string) bool {

	runeInput := []rune(s)
	runePattern := []rune(p)

	lenInput := len(runeInput)
	lenPattern := len(runePattern)

	isMatchingMatrix := make([][]bool, lenInput+1)

	for i := range isMatchingMatrix {
		isMatchingMatrix[i] = make([]bool, lenPattern+1)
	}

	isMatchingMatrix[0][0] = true
	for i := 1; i < lenInput; i++ {
		isMatchingMatrix[i][0] = false
	}

	if lenPattern > 0 {
		if runePattern[0] == '*' {
			isMatchingMatrix[0][1] = true
		}
	}

	for j := 2; j <= lenPattern; j++ {
		if runePattern[j-1] == '*' {
			isMatchingMatrix[0][j] = isMatchingMatrix[0][j-1]
		}

	}

    for i:=0;i<=lenInput;i++{
        fmt.Println(isMatchingMatrix[i])
        //for j := 0; j <= lenPattern; j++ { }
    }

	for i := 1; i <= lenInput; i++ {
		for j := 1; j <= lenPattern; j++ {

			if runePattern[j-1] == '*' {
				isMatchingMatrix[i][j] = isMatchingMatrix[i-1][j] || isMatchingMatrix[i][j-1]
			}

			if runePattern[j-1] == '?' || runeInput[i-1] == runePattern[j-1] {
				isMatchingMatrix[i][j] = isMatchingMatrix[i-1][j-1]
			}
		}
	}

	return isMatchingMatrix[lenInput][lenPattern]
}
