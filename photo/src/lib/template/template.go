package template
import(
	"log"
	"html/template"
	"net/http"
	"path"
	"strings"
	"io/ioutil"
)
var Tmpl map[string]*template.Template

func init(){
	Tmpl=make(map[string]*template.Template)
	fileInfoSlice,err:=ioutil.ReadDir("./public")

	if err!=nil{
		log.Fatal("lib/template.go:init():ReadDir ",err)
	}

	for _,fileInfo:=range fileInfoSlice {
		ext:=path.Ext(fileInfo.Name());

			if ext!=".html"{
				continue
			}

			tmplName:="photo/"+strings.TrimSuffix(fileInfo.Name(),ext)
		Tmpl[tmplName]=template.Must(template.ParseFiles("./public/"+fileInfo.Name()))
		log.Println("lib/template.go:succeed to load template:"+tmplName)
	}
}

// 老方法，缺点是每次调用都要要读取和渲染模板,现在采用下面的方法:RenderHTML
func oldRenderHTML(method string,w http.ResponseWriter,tmpfile string,data map[string]interface{})error{
	t,err:=template.ParseFiles(tmpfile)
	if err!=nil{
		return err
	}
	if err=t.Execute(w,data);err!=nil{
		return err
	}
	return nil
}


func RenderHTML(w http.ResponseWriter,tmpl string,data map[string]interface{})error{
	if err:=Tmpl[tmpl].Execute(w,data);err!=nil{
		return err
	}
	return nil
}
