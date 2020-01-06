// package ctci contains solutions for various problems presented in Cracking the Coding Interview
package ctci

// findSpace returns the index of the first space character in a string, or -1 if there is none
func findSpace(s string) int {
	for i, char := range s {
		if char == ' ' {
			return i
		}
	}
	return -1
}

// urlify replaces all spaces in a string with '%20'
func urlify(s string) string {
	spaceIndex := findSpace(s)
	ret := s
	for spaceIndex != -1 {
		ret = ret[:spaceIndex] + "%20" + ret[spaceIndex+1:]
		spaceIndex = findSpace(ret)
	}
	return ret
}
