package main

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(n1, n2 int) int {
	n := minn(n1, n2)
	for n > 0 {
		if n1%n == 0 && n2%n == 0 {
			break
		}
		n--
	}
	return n
}

func minn(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
