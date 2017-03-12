package router
import(
	"net/http"
	"log"
	"io"
	"io/ioutil"
//	"mime/multipart"
	"os"
	tmpl"lib/template"
)

const(
UPLOAD_DIR="./uploads/"
)


func Upload(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		err:=tmpl.RenderHTML(w,"photo/upload",nil)
		if err!=nil{
			log.Fatal("Upload,render:photo/upload",err)
			http.Error(w,err.Error(),500)
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

func View(w http.ResponseWriter,r *http.Request){

	name:=r.FormValue("name")
	if name==""{
		http.NotFound(w,r)
		log.Println("name is nil")
		return
	}
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

func isExist(path string)bool{
	_,err:=os.Stat(path)
	if err==nil{
		return true
	}
	return os.IsExist(err)
}


func ListView(w http.ResponseWriter,r *http.Request){
	log.Println("router/photoWeb.go/ListView:ready to Auth")
	u,ok:=Auth(w,r)
	if !ok{
		log.Println("ListView:Auth fail")
		return
	}

	log.Println("Listview is request by "+u.Name)

	fileInfoSlice,err:=ioutil.ReadDir(UPLOAD_DIR)
	if err!=nil{
		log.Println("ListView,ioutil.ReadDir:",err.Error())
		return
	}

	data:=make(map[string]interface{})
	fileNameSlice:=make([]string,0)
	for _,fileInfo:=range fileInfoSlice{
		fileNameSlice=append(fileNameSlice,fileInfo.Name())
	}
	data["imagename"]=fileNameSlice
	err=tmpl.RenderHTML(w,"photo/list",data)
	if err!=nil{
		log.Println("ListView:render list.html",err.Error())
		return
	}
	log.Println(http.StatusOK,"ListView")
	return
}
