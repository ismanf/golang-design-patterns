package template_p

import (
	"fmt"
	
)

type Greeter interface {
	message() string
}

type GreeterTemplate interface {
	first() string
	third() string
	greet(Greeter) string
}

type PrincesGreeterTmpl struct {}
func (p PrincesGreeterTmpl) first() string {
	return "Welcome"
}
func (p PrincesGreeterTmpl) third() string {
	return "to our palace!"
}
func (p PrincesGreeterTmpl) greet(g Greeter) string {
	return fmt.Sprintf("%s, %s %s", p.first(), g.message(), p.third())
}

type PrincesGreeter struct {
	name string
}
func (p *PrincesGreeter) message() string {
	return fmt.Sprintf("your magesty princes %s", p.name)
}

func Demo()  {
	prGreeter := &PrincesGreeter{"Diana"}
	tmpl := PrincesGreeterTmpl{}

	fmt.Println(tmpl.greet(prGreeter))
}