package router
import(
	"net/http"
	"log"
	"io"
	"io/ioutil"
//	"mime/multipart"
	"os"
	"html/template"
	"lib"
)


func SayHello(w http.ResponseWriter,r *http.Request){
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

func Upload(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		err:=lib.RenderHTML(w,"photo/upload",nil)
		if err!=nil{
			log.Fatal("Upload,render:photo/upload",err)
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
		http.NotFound(w,"name is nil")
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
	fileInfoSlice,err:=ioutil.ReadDir(UPLOAD_DIR)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			log.Fatal("ListView,ioutil.ReadDir:",err)
			return
		}

		data:=make(map[string]interface{})
		fileNameSlice:=make([]string,0)
		for _,fileInfo:=range fileInfoSlice{
			fileNameSlice=append(fileNameSlice,fileInfo.Name())
		}
		data["imagename"]=fileNameSlice
		err=lib.RenderHTML(w,"photo/list",data)
		if err!=nil{
			log.Fatal("ListView:render list.html",err)
			return
		}
		log.Println(http.StatusOK,"ListView")
	return
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

