package sqrt
import (
	"math"
)
// Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:
func Sqrt(x float64) float64 {
	z := 100.0
	delta := 0.000001;
	last := 0.0;
	for {
		last = z;
		z=z-(z*z - x)/(2*z)
		if(math.Abs(z-last) < delta){
			return z
		}
	}
}
