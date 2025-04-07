package main

import "unicode"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isLetterOfNum(rune(s[l])) {
			l++
		} else if !isLetterOfNum(rune(s[r])) {
			r--
		} else if toLowerCase(rune(s[l])) == toLowerCase(rune(s[r])) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome2(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		} else if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		} else if unicode.ToLower(rune(s[l])) == unicode.ToLower(rune(s[r])) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isLetterOfNum(ch rune) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9')
}

func toLowerCase(ch rune) rune {
	if 'A' <= ch && 'Z' >= ch {
		return ch + 32
	}
	return ch
}
