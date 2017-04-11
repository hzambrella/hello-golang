package os
import(
	"os"
	"fmt"
)

func OsStudy(){
	fmt.Println("haha")
	// os.Getenv
	fmt.Println("os.Getenv:",os.Getenv("GOPATH"))
}
