package main

func main() {
}

func anagrams(s []string) [][]string {
	var res [][]string
	added := make(map[string]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := added[s[i]]; ok {
			continue
		} else {
			res = append(res, []string{s[i]})
			added[s[i]] = struct{}{}
		}
		for j := i + 1; j < len(s); j++ {
			if isAnagram(s[i], s[j]) {
				res[i] = append(res[i], s[j])
				added[s[j]] = struct{}{}
			}
		}
	}

	return res
}

func isAnagram(s1 string, s2 string) bool {
	var m1 = countLetters(s1)
	var m2 = countLetters(s2)

	if len(m1) != len(m2) {
		return false
	}

	for s, c := range m1 {
		if m2[s] != c {
			return false
		}
	}

	return true
}

func countLetters(s string) map[byte]int {
	var m = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 1
		} else {
			m[s[i]]++
		}
	}

	return m
}
