package transform

func ArrayStringUniq(arr []string) []string {
	set := make(map[string]struct{}, 0)
	for _, s := range arr {
		set[s] = struct{}{}
	}
	var newStr = make([]string, 0)
	for k := range set {
		newStr = append(newStr, k)
	}
	return newStr
}
