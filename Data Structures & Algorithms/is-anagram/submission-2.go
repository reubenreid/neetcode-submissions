func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false;
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for i, char := range s {
		sMap[char]++;
		tMap[rune(t[i])]++;
	}

	for k, v := range sMap {
		if v != tMap[k] {
			return false
		} 
	}

	return true
}
