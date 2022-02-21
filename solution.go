package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sidesNumType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const SidesSquare sidesNumType = 4
const SidesTriangle sidesNumType = 3
const SidesCircle sidesNumType = 0
const Pi float64 = 3.141

func CalcSquare(sideLen float64, sidesNum sidesNumType) float64 {
	var square float64
	if sidesNum == 3 {
		square = ((math.Sqrt(3) * math.Pow(sideLen, 2)) / 4)
	} else if sidesNum == 4 {
		square = (math.Pow(sideLen, 2))
	} else if sidesNum == 0 {
		square = (Pi * math.Pow(sideLen, 2))
	} else {
		square = (0)
	}
	return (square)
}
