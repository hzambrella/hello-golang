package graph

import(
	"fmt"
	"path"
//	"bytes"
	"os"
	"os/exec"
)

//graphviz+DOT

//directory is relative directory where  saves .vrg file
//fName is name of file ,like fName.vrg
//drawStr is graphviz language which use to draw graph
// return relative  path  of file and error,like diretory/fName.vrg
func SaveToFileVrg(directory,fName,drawStr string)(string,error){
	var filepath string

	fileName:=fmt.Sprintf("%s.vrg",fName)
	if directory!=""{
		filepath=path.Join(directory,fileName)
	}else{
		filepath=fileName
	}

	f,err:=os.Create(filepath)
	if err!=nil{
		return "",err
	}
	_,err=f.WriteString(drawStr)
	if err!=nil{
		return "",err
	}
	defer f.Close()
	return filepath,nil
}

// execute dot -Tpng -o  /pngdst/fName.vrg  vrgsrc/fName.vrg in cmd 
// vrgsrc is relative directory path where have .vrg file
// pngdst is destination of directory path to save graph,which format .png
// fName is name of .vrg file ,like fName.vrg
// return error and relative path of .png file,like pngdst/fName.png
func DrawVrg(vrgsrc,pngdst,fName string)(string,error){
	var fvrg string
	var fpng string
	if vrgsrc!=""{
		fvrg=path.Join(vrgsrc,fName)
	}
	if pngdst!=""{
		fpng=path.Join(pngdst,fName)
	}

	s:=fmt.Sprintf("dot -Tpng -o %s.png %s.vrg",fpng,fvrg)
	fmt.Println(s)
	cmd:=exec.Command(`/bin/sh`,`-c`,s)
	//var out bytes.Buffer
	//cmd.Stdout=&out
	if err:=cmd.Run();err!=nil{
		fmt.Println("errors,",err.Error())
		return "",err
	}
	return fpng,nil
}
