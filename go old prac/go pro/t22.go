package main
import (
	"fmt"
	"math"
)

func pow(x,n,lim float64) float64{
	if v:=math.Pow(x,n);v<lim {
		return v
	}
	return lim
	//return v //Scope of variable v is inside the block only can't be used outside!
}

func main(){
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}