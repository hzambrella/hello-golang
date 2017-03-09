package router
import(
	"net/http"
	"fmt"
	"log"
	"io"
	"io/ioutil"
//	"mime/multipart"
	"os"
)

func SayHelloOld(w http.ResponseWriter,r *http.Request){
	userName:="haozhao"//TODO:real username
	helloMes:="hello "+userName+"!"
	_,err:=io.WriteString(w,helloMes)
	if err!=nil{
		log.Fatal("SayHello:",err)
		return
	}
	log.Println(http.StatusOK,r.Method,"SayHello")
	return
}

const(
	upload="<html><head><title>yellow picture</title></head><body><form method=\"POST\" action=\"/upload\" "+
	" enctype=\"multipart/form-data\">"+
	"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
	"<input type=\"submit\" value=\"Upload\" />"+
	"</form><body></html>"
	UPLOAD_DIR="./uploads"
)
func UploadOld(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		_,err:=io.WriteString(w,upload)
		if err!=nil{
			log.Fatal("Upload",err)
			return
		}
	log.Println(http.StatusOK,r.Method,"Upload")
	return
	}

	if r.Method=="POST"{
		f,h,err:=r.FormFile("image")
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			log.Fatal("Upload,FormFile(image):",err)
			return
		}

		filename:=h.Filename
		defer f.Close()
		file,err:=os.Create(UPLOAD_DIR+"/"+filename)
		defer file.Close()
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			log.Fatal("Upload,os.Create:",err)
			return
		}

		_,err=io.Copy(file,f)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			log.Fatal("Upload,io.Copy:",err)
			return
		}

		http.Redirect(w,r,"/view?name="+filename,http.StatusFound)
		return
	}
}

func ViewOld(w http.ResponseWriter,r *http.Request){
	name:=r.FormValue("name")
	imageName:=UPLOAD_DIR+"/"+name
	if exist:=isExist(imageName);exist!=true{
		http.NotFound(w,r)
		log.Println("view"+imageName+" not found")
		return
	}
	w.Header().Set("Content_Type","image")
	http.ServeFile(w,r,imageName)
	log.Println(http.StatusOK,r.Method,"View")
	return
}

func isExistOld(path string)bool{
	_,err:=os.Stat(path)
	if err==nil{
		return true
	}
	return os.IsExist(err)
}

const(
	listHtmlAll="<html><head><title>yellow picture listview</title></head><body>%s<body></html>"
)

func ListViewOld(w http.ResponseWriter,r *http.Request){
	fileInfoSlice,err:=ioutil.ReadDir(UPLOAD_DIR)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			log.Fatal("ListView,ioutil.ReadDir:",err)
			return
		}

		var listHtml string
		for _,fileInfo:=range fileInfoSlice {
			imgName := fileInfo.Name()
			listHtml += "<li><a href=\"/view?name="+imgName+"\">imgname</a></li>"
		}
		listHtmlEnd:=fmt.Sprintf(listHtmlAll,"<ol>"+listHtml+"</ol>")
	io.WriteString(w, listHtmlEnd)
	return
}


