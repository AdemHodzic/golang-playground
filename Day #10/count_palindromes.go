package main 

func CountPalindromes(str string) int {

	current := len(str)

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if isPalindrome(string(str[i:j])) == true{
				current++
			}
		}
	}

	return current

}

func isPalindrome(str string) bool {
	mid := len(str) / 2
	last := len(str) - 1

	for i := 0; i < mid; i++ {
		if str[i] != str[last - i] {
			return false
		}
	}

	return true
}