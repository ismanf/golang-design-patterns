package prototye

import "fmt"
import "errors"

const (
	WHITE = 1
	GREEN = 2
	BLACK = 3
)

type ProductInfoGetter interface {
	GetInfo() string
}

type Tshirt struct {
	Size string 
	Color string
}

func (t Tshirt) GetInfo() string {
	return fmt.Sprintf("Tshirt size: '%s', Color: '%s'.", t.Size, t.Color)
}

var whiteTshirt *Tshirt = &Tshirt{ "M", "white" }
var greenTshirt *Tshirt = &Tshirt{ "L", "green" }
var blackTshirt *Tshirt = &Tshirt{ "S", "black" }

func GetTshirClone(t int) (ProductInfoGetter, error) {
	switch t {
		case WHITE:
			return whiteTshirt, nil
		case GREEN:
			return greenTshirt, nil
		case BLACK:
			return blackTshirt, nil
		default:
		return nil, errors.New(fmt.Sprintf("Prototype id %d not recognized.", t))
	}
}







