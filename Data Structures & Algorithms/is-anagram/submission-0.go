import "maps";

func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false;
	}

	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i, _ := range s {
		sMap[s[i]] = sMap[s[i]] + 1;
		tMap[t[i]] = tMap[t[i]] + 1;
	}

	return maps.Equal(sMap, tMap);
}
