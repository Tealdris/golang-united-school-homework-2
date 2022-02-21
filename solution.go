package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sidesNumType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const SS int = 4
const ST int = 3
const SC int = 0
const Pi float64 = 3.141

func CalcSquare(sideLen float64, sidesNum sidesNumType) float64 {

	if sidesNum == 3 {
		fmt.Println((math.Sqrt(3) * math.Pow(sideLen, 2)) / 4)
	} else if sidesNum == 4 {
		fmt.Println(math.Pow(sideLen, 2))
	} else if sidesNum == 0 {
		fmt.Println(Pi * math.Pow(sideLen, 2))
	} else {
		fmt.Println(0)
	}
	return
}
