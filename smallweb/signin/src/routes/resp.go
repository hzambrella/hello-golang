package routes

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type tempStore map[string]*template.Template //TODO:读取public下的模板文件，并进行缓存
type H map[string]interface{}

func init() {
	//TODO:读取public下的模板文件，并进行缓存
}

func String(w http.ResponseWriter, code int, result string) {
	defer func() {
		if err := recover(); err != nil {
			logl.Error(err.(error))
		}
	}()

	w.WriteHeader(code)
	w.Header().Set("codelog", strconv.Itoa(code))
	_, err := io.WriteString(w, result)
	if err != nil {
		panic(err)
	}
}

func JSON(w http.ResponseWriter, code int, result H) {
	defer func() {
		if err := recover(); err != nil {
			logl.Error(err.(error))
		}
	}()

	w.WriteHeader(code)
	w.Header().Set("codelog", strconv.Itoa(code))
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}

func Render(w http.ResponseWriter, code int, tempName string, result H) {
	defer func() {
		if err := recover(); err != nil {
			logl.Error(err.(error))
		}
	}()

	w.WriteHeader(code)
	w.Header().Set("codelog", strconv.Itoa(code))

	//TODO:读取public下的模板文件，并进行缓存
	// 这种方式每次调用方法都要加载模板，效率低
	t := template.Must(template.ParseFiles(tempName))
	err := t.Execute(w, result)
	if err != nil {
		panic(err)
	}
	return

}
