package leetcode_go

import "testing"

func TestValidAnagram(t *testing.T) {
	t.Logf("res is: %v", isAnagram("s", "t"))
	t.Logf("res is: %v", isAnagram("", "t"))
	t.Logf("res is: %v", isAnagram("s", "s"))
	t.Logf("res is: %v", isAnagram("anagram", "nagaram"))
	t.Logf("res is: %v", isAnagram("rat", "car"))

}
