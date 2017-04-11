package os
import(
	"os"
)

func CheckFilesExist(fileName string)(bool,error){
	_,err:=os.Stat(fileName)
	if err!=nil{
		if os.IsNotExist(err){
			return false,nil
		}else{
			return true,err
		}
	}
	return true,nil
}
