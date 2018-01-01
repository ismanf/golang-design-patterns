package composition

import (
	"fmt"
)

type Driver interface {
	Drive()
}

type Cleaner interface{
	Clean()
} 

type Killer interface {
	Kill()
}

type Barker interface {
	Bark()
}

type Dog struct {}
func (d *Dog) Bark()  {
	fmt.Println("Wow!")
}

type RobotDog struct {}
func (r *RobotDog) Drive()  {
	fmt.Println("Driving...")
}
func (r *RobotDog) Bark()  {
	fmt.Println("Wow!")
}

type CleanerRobotDog struct {}
func (r *CleanerRobotDog) Drive()  {
	fmt.Println("Driving...")
}
func (r *CleanerRobotDog) Bark()  {
	fmt.Println("Wow!")
}
func (r *CleanerRobotDog) Clean()  {
	fmt.Println("Cleaning...")
}

type KillerRobotDog struct {}
func (r *KillerRobotDog) Drive()  {
	fmt.Println("Driving...")
}
func (r *KillerRobotDog) Bark()  {
	fmt.Println("Wow!")
}
func (r *KillerRobotDog) Kill()  {
	fmt.Println("You gonna die!")
}