package plush

// Gimme can give you a kiss or a cuddle
func Gimme(what string) (t string) {
	// if(what == "kiss")
	switch what {
	case "kiss":
		t = "jjj"
		break
	case "hug":
		t = "mouah"
		break
	default:
		t = "waaaaaaaat"
		break
	}
	return t
}
