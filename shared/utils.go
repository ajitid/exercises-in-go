package shared

import "github.com/ajitid/litter"

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func PreMain() {
	litter.Config.HidePrivateFields = false
	litter.Config.StripPackageNames = true
	litter.Config.ShowRune = true
}
