package main
import "fmt"

func main(){
	countryCapitalMap:=map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
	
	fmt.Println("Original Map")
	for country:=range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
	
	delete(countryCapitalMap,"France");
	fmt.Println("Entry for France is deleted")
	
	fmt.Println("Updated Map")
	for country:=range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}