package proxy

import "errors"

type Human struct {
	Gender string
	Age    int
}

type Bar interface {
	Welcome(Human) error
}

type OpenBar struct {
	Visitors []Human
}

func (b *OpenBar) Welcome(h Human) error {
	b.Visitors = append(b.Visitors, h)
	return nil
}

type BarProxy struct {
	Bar
}

func (b *BarProxy) Welcome(h Human) error {
	if h.Gender == "female" && h.Age < 20 {
		return errors.New("Females under 20 not allowed.")
	}

	if h.Gender == "male" && h.Age < 18 {
		return errors.New("Males under 18 not allowed.")
	}

	return b.Bar.Welcome(h)
}
