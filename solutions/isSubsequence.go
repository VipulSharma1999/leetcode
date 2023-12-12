package main

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	sidx, tidx := 0, 0
	for sidx < ls && tidx < lt {
		if s[sidx] == t[tidx] {
			sidx++
		}
		tidx++
	}
	return sidx == ls
}
