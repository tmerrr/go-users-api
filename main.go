package main

func main() {
	user := newUser(
		"John",
		"Wick",
		"john@wick.com",
		"good-dog",
	)
	err := user.save()
	if err != nil {
		panic(err.Error())
	}
}
