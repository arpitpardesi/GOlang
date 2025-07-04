package main
import (
   "os"
   "text/template"
)
func main() {
       const tmpl = `{{range .}}{{if .Active}}Name: {{.Name}}, Age: {{.Age}}
{{end}}{{end}}`

   t, err := template.New("test").Parse(tmpl)
   if err != nil {
      panic(err)
   }

   type Person struct {
      Name   string
      Age    int
      Active bool
   }
   people := []Person{
      {Name: "Revathi", Age: 25, Active: true},
      {Name: "Tapas", Age: 35, Active: false},
      {Name: "Vivek", Age: 27, Active: true},
   }

   err = t.Execute(os.Stdout, people)
   if err != nil {
      panic(err)
   }
} 