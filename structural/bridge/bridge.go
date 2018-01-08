package bridge

import "fmt"
import "errors"

// Bridge interface
type Drawer interface {
	DrawRectangle(w, h int) 
}

type GreenRectangleDrawer struct {}
func (d GreenRectangleDrawer) DrawRectangle(w, h int) {
	fmt.Printf("Green rectangle W: %d, H: %d", w, h)
}

type RedRectangleDrawer struct {}
func (d RedRectangleDrawer) DrawRectangle(w, h int) {
	fmt.Printf("Red rectangle W: %d, H: %d", w, h)
}

type Shape interface {
	Draw() error
}

type Rectangle struct {
	Width, Height int
	drawer Drawer
}

func (d Rectangle) Draw () error {
	if d.drawer == nil {
		return errors.New("Drawer not initialized.")
	}

	d.drawer.DrawRectangle(d.Width, d.Height)
	return nil
}