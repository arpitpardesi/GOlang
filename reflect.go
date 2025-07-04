package main
import (
   "fmt"
   "reflect"
)
func main() {
   var x float64 = 3.14
   t := reflect.TypeOf(x)
   fmt.Println("Type:", t) 
   v := reflect.ValueOf(x)
   fmt.Println("Value:", v)
   fmt.Println("Kind:", v.Kind())
   if v.Kind() == reflect.Float64 {
      fmt.Println("x is a float64")
   }
   p := reflect.ValueOf(&x)
   fmt.Println("Pointer to x:", p)
   e := p.Elem() //
   fmt.Println("Element of pointer:", e)
   if e.CanSet() {
      e.SetFloat(2.71) 
      fmt.Println("Modified value of x:", x) 
   }
}