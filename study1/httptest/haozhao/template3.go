package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Haha   bool
	Emails []string
	Jobs   []*Job
}
type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}.
The age is {{.Age|dayago}}.
{{if .Haha}}hahahahhaa{{else}}wawawawawaawa{{end}}
{{range.Emails}}An email is {{.}}
{{end}}
{{with.Jobs}}
{{range.}}An employer is {{.Employer}}
and the role is {{.Role}}
{{end}}
{{end}}
`

func main() {
	job1 := Job{Employer: "lol", Role: "biubiubiu"}
	job2 := Job{Employer: "Box", Role: "duang"}
	person := Person{
		Name:   "haozhao",
		Age:    50,
		Emails: []string{"44143143@qq.com", "6666@qq.com"},
		Jobs:   []*Job{&job1, &job2},
		Haha:   false,
	}
	var t = template.Must(template.New("Person template").Funcs(template.FuncMap{"dayago": dayago}).Parse(templ))
	_ = t.Execute(os.Stdout, person)
}

func dayago(Age int) int {
	return Age / 4
}
