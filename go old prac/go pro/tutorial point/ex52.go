package main
import "fmt"
func main(){
	var countryCapitalMap map[string] string
	countryCapitalMap=make(map[string]string)
	
	countryCapitalMap["France"]="Paris"
	countryCapitalMap["Italy"]="Rome"
	countryCapitalMap["Japan"]="Tokyo"
	countryCapitalMap["India"]="New Delhi"
	
	for country:=range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
	capital,ok:=countryCapitalMap["UnitedStates"]
	//capital,ok:=countryCapitalMap["France"]
	
	if(ok){
			fmt.Println("Capital of United States is",capital)
		}else{
			fmt.Println("Capital of United States is not present")
		}
	//fmt.Println(ok) //false
	//fmt.Println(capital) //empty
}