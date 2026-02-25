package main

func main() {

	r := New(100).Add(10.4).Add(20.5).Sub(10.4).Mul(2).Div(2).Get()

	println(r)
}

type Calc struct {
	Data float64
}

func New(d float64) *Calc {
	return &Calc{d}
}

func (c *Calc) Add(d float64) ICalc {
	c.Data += d
	return c
}

func (c *Calc) Sub(d float64) ICalc {
	c.Data -= d
	return c
}

func (c *Calc) Mul(d float64) ICalc {
	c.Data *= d
	return c
}
func (c *Calc) Div(d float64) ICalc {
	c.Data /= d
	return c
}
func (c *Calc) Get() float64 {
	return c.Data
}

type ICalc interface {
	Add(float64) ICalc
	Sub(float64) ICalc
	Mul(float64) ICalc
	Div(float64) ICalc
	Get() float64
}
