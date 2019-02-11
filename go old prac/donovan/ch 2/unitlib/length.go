package unitlib

//FToM function
func FToM(f Feet) Meters { return Meters(f / 3.2808) }

//MToF function
func MToF(m Meters) Feet { return Feet(m * 3.2808) }
