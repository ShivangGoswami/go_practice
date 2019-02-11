package main
import "fmt"
func main(){
	defer fmt.Println("world")
	fun()
	fmt.Println("hello")
}
func fun(){
	ret
}