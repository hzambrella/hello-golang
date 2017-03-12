package code
import (
	"log"
	"encoding/base64"
	"encoding/json"
)

func Encode(v interface{})(string,error){
	jData,err:=json.Marshal(v)
	log.Println("lib/code.go:Encode ",jData)
	if err!=nil{
		return "",err
	}

	str:=base64.RawURLEncoding.EncodeToString(jData)
	log.Println("lib/code.go:Encode,base64.Encode",str)
	return str,nil
}

func Decode(src string,v interface{})error{
	data,err:=base64.RawURLEncoding.DecodeString(src)
	if err!=nil{
		return err
	}
	log.Println("lib/code.go:base64.Decode ",data)

	if err:=json.Unmarshal(data,v);err!=nil{
		return err
	}

	return nil
}
