package main

func lengthOfLongestSubstring(s string) int {
	result := 0

	runes := []rune(s)

	runeMap := make(map[rune]int)

	left := 0

	for right, ch := range runes {
		if ind, ok := runeMap[ch]; ok && ind >= left {
			left = ind + 1
		}
		runeMap[ch] = right
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}
