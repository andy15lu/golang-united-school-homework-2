package square

import "math"
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type side int;

func CalcSquare(sideLen float64, sidesNum side) float64 {
	var square float64;
	switch (sidesNum) {
	case 0:
		square = sideLen*math.Pi
	case 3:
		square = math.Sqrt(3)*sideLen*sideLen/4
	case 4:
		square = sideLen*sideLen
//	default:
//		return (float64)sidesNum*sideLen*sideLen/( 4*math.Tan( 180/(float64)sidesNum ) )
	}
	return square 
}
