package Array

func MergeAlternately(word1 string, word2 string) string {
	var idx_word1, idx_word2 int
	var merged_word string
	for idx_word1 < len(word1) && idx_word2 < len(word2) {
		merged_word = merged_word + string(word1[idx_word1]) + string(word2[idx_word2])
		idx_word1++
		idx_word2++
	}

	if idx_word1 < len(word1) {
		merged_word += string(word1[idx_word1:])
	}
	if idx_word2 < len(word2) {
		merged_word += string(word2[idx_word2:])
	}
	return merged_word
}
