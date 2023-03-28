package gosayhello

func SayHello(name string) string {
	return "Hello" + name + "How are you today?	;)"
}

func Age(age int) string {
	if age <= 1 {
		return "Anda di kategorikan Bayi..!"
	} else if age <= 5 {
		return "Anda di kategorikan Balita..!"
	} else if age <= 12 {
		return "Anda di kategorikan Anak-anak..!"
	} else if age <= 18 {
		return "Anda di kategorikan Remaja..!"
	} else if age <= 60 {
		return "Anda di kategorikan Dewasa..!"
	} else {
		return "Anda di kategorikan Manula..!"
	}
}
