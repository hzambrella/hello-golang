package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	muban_eng := `{{.Count}} items
	are made of {{.Material}}
	`
	muban_chn := "{{.Material}}做了{{.Count}}个项目\n"
	//建立一个模板的名称是china，模板的内容是muban_chn字符串
	tmpl := template.New("china")
	tmpl, _ = tmpl.Parse(muban_chn)
	//建立一个模板的名称是english，模板的内容是muban_eng字符串
	tmpl = tmpl.New("english")
	tmpl, _ = tmpl.Parse(muban_eng)
	//将struct与模板合成，用名字是china的模板进行合成，结果放到os.Stdout里，内容为“wool做了17个项目”
	_ = tmpl.ExecuteTemplate(os.Stdout, "china", sweaters)
	//将struct与模板合成，用名字是china的模板进行合成，结果放到os.Stdout里，内容为“17 items are made of wool”
	_ = tmpl.ExecuteTemplate(os.Stdout, "english", sweaters)

	tmpl = template.New("english")
	fmt.Println(tmpl.Name()) //打印出english
	tmpl = tmpl.New("china")
	fmt.Println(tmpl.Name()) //打印出china
	//	tmpl = tmpl.Lookup("english") //必须要有返回，否则不生效
	//	fmt.Println(tmpl.Name())      //打印出english
}
