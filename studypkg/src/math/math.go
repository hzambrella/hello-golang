package math
import (
	"fmt"
//	"log"
	"math"
)
//"math"
//"math/rand"
//"math/rand"
// IEEE 754
//search math
// part of math
// "math"
func ExpMath(){
	// 10^e
	fmt.Println("math.Pow10(9)",math.Pow10(9))
	// x^y
	fmt.Println("math.Pow",math.Pow(2,2))
	// <=
	fmt.Println("math.Floor(11021.021)",math.Floor(11021.021))
	//>=
	fmt.Println("math.Ceil(-11021.021)",math.Ceil(-11021.021))
	// li fang gen
	fmt.Println("math.Cbrt(-11021.021)",math.Cbrt(-11021.021))
	// ping fang gen  sqrt
	// max(x-y,0)
	fmt.Println("math.Dim(1,2)",math.Dim(1,2))
	// e^2
	fmt.Println("Exp",math.Exp(2))
	// sqrt(a*a+b*b)
	fmt.Println("hypot",math.Hypot(3,4))
	// inf 
	fmt.Println("inf",math.Inf(-1))
	//nan
	fmt.Println("isNan",math.IsNaN(100))
	fmt.Println("Nan",math.NaN())
	// max ,min
	fmt.Println("max",math.Max(1,2))
	// Mod  yushu
	fmt.Println("mod",math.Mod(4,4))
	// remain yushu
	fmt.Println("remainder",math.Remainder(4,4))
	// -  or -0
	fmt.Println("signbit",math.Signbit(-0))
	// zheng shu bu  fen
	fmt.Println("Trunc",math.Trunc(12.3))
}

func SearchRand(){
}
