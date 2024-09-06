package Array

func GcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return GcdOfStrings(str1[len(str2):], str2)
}
