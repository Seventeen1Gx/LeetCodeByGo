package LeetCode

func VowelStrings(words []string, left, right int) int {
	var ans int
	for i := left; i < right+1; i++ {
		if isVowelString(words[i]) {
			ans++
		}
	}
	return ans
}

func isVowelString(word string) bool {
	n := len(word)
	if n <= 0 {
		return false
	}
	if (word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u') &&
		(word[n-1] == 'a' || word[n-1] == 'e' || word[n-1] == 'i' || word[n-1] == 'o' || word[n-1] == 'u') {
		return true
	}
	return false
}
