package main
import "fmt"
type Animal struct {
   Name string
}
func (a *Animal) Speak() {
   fmt.Println(a.Name, "makes a sound")
}
type Dog struct {
   Animal
   Breed  string
}
func (d *Dog) Speak() {
   fmt.Println(d.Name, "barks")
}
func main() {
   d := Dog{
      Animal: Animal{Name: "Rex"},
      Breed:  "German Shepherd",
   }
   d.Speak()
   }