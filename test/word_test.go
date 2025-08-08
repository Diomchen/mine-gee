package test

import "testing"

func wordBreak(s string, wordDict []string) bool {
	strLen := len(s)
	for t := 0; t < len(wordDict); t++ {
		i := 0
		for i < strLen {
			isMatch := true
			// 判断
			for j := t; ; {
				v := wordDict[j%len(wordDict)]
				wLen := len(v)
				if i+wLen <= strLen {
					subStr := s[i : i+wLen]
					if subStr == v {
						i += wLen
						break
					}
				}
				j++
				if j%len(wordDict) == t {
					isMatch = false
					break
				}
			}

			if !isMatch {
				break
			} else if i == strLen {
				return true
			}
		}
	}
	return false
}

func TestWordBreak(t *testing.T) {
	s := "catskicatcats"
	wordDict := []string{"cat", "dog", "ski", "cats"}
	if !wordBreak(s, wordDict) {
		t.Error("Failed")
	}
}
