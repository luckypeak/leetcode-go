package leetcode_go


func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	cnt := make([]int, 26)
	for _, c := range s{
		cnt[c-'a'] ++
	}

	for _, c := range t{
		cnt[c-'a'] --
	}

	for _, c:=range cnt{
		if c != 0{
			return false
		}
	}
	return true
}

