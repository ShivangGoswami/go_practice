//Package word provides utilities for word games.
package word

//IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	// for i := range s {
	// 	if s[i] != s[len(s)-1-i] {
	// 		return false
	// 	}
	// }
	n := len(s) / 2
	for i := 0; i < n; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
