package main
import "fmt"
func main(){
	var string="hello world!"
	
	for i,j:=range string{
	fmt.Printf("string[%d]=%c\n",i,j)
	}
}