package main

import (
	"strconv"
	"strings"
)

func intArrArrEqualNoOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int, len(a))
	for _, arr := range a {
		m[intsToStr(arr)]++
	}
	for _, arr := range b {
		m[intsToStr(arr)]--
	}
	for _, diff := range m {
		if diff != 0 {
			return false
		}
	}
	return true
}

func intsToStr(arr []int) string {
	var sb strings.Builder
	for _, n := range arr {
		sb.WriteRune('#')
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}
