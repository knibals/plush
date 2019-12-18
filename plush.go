package plush

// Gimme can give you a kiss or a cuddle
func Gimme(what string) (t string) {
	switch what {
	case "kiss":
		t = "😘"
		break
	case "hug":
		t = "🤗"
		break
	default:
		t = "waaaaaaaat!!!"
		break
	}
	return t
}
