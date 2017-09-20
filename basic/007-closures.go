package main

func nextStep() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := nextStep()
	println(nextInt())
	println(nextInt())
	println(nextInt())

	println("new instance")

	newNextInt := nextStep()
	println(newNextInt())
	println(newNextInt())

	i := 1
	nextX := func() int {
		i *= 2
		return i
	}
	println("next power")
	println(nextX())
	println(nextX())
	println(nextX())
}
