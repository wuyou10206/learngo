package norepeating

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr2("我爱你"))
}
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr3(s string) int {
	start := 0
	maxLength := 0
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	for i, ch := range []rune(s) {
		lastI := lastOccurred[ch]
		if lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
