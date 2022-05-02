package arrayUtils

func ArrayStringIn(list []string, found string) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}
func ArrayInt64In(list []int64, found int64) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}
