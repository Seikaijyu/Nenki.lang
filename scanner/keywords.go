package scanner

var keywords []string = []string{
	"",
}

func FindKeyWord(find string) bool {
	for _, value := range keywords {
		if find == value {
			return true
		}
	}
	return false
}
