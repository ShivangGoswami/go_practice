package unitlib

//KgToPound function
func KgToPound(kg Kilograms) Pound { return Pound(kg * 2.2046) }

//PoundToKg function
func PoundToKg(p Pound) Kilograms { return Kilograms(p / 2.2046) }
