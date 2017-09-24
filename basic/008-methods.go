package main

type calc struct {
	x, y int
}

func (c *calc) add() int {
	return c.x + c.y
}

func (c *calc) sub() int {
	return c.x - c.y
}

func (c *calc) multi() int {
	return c.x * c.y
}

func (c *calc) div() int {
	return int(c.x / c.y)
}

func main() {
	calculator := calc{}
	calculator.x = 10
	calculator.y = 2

	// add
	println(calculator.add())
	println(calculator.sub())
	println(calculator.multi())
	println(calculator.div())
}
