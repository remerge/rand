// based on git@github.com:leesper/go_rng.git
package rand

import (
	"fmt"
	"math"
)

// Function to provide a random float64 value
type GetRandomFloat64Fn func() float64

// Gamma returns a random number of gamma distribution (alpha > 0.0 and beta > 0.0).
// Uses package's built-in float64 randomization routine.
func Gamma(alpha, beta float64) float64 {
	return calculateGamma(alpha, beta, Float64)
}

// Gamma returns a random number of gamma distribution (alpha > 0.0 and beta > 0.0).
// Uses externally provided float64 randomization routine.
func GammaEx(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	return calculateGamma(alpha, beta, random)
}

func calculateGamma(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	if !(alpha > 0.0) || !(beta > 0.0) {
		panic(fmt.Sprintf("Invalid parameter alpha %.2f beta %.2f", alpha, beta))
	}
	return gamma(alpha, beta, random)
}

func gamma(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	var MAGIC_CONST float64 = 4 * math.Exp(-0.5) / math.Sqrt(2.0)
	if alpha > 1.0 {
		// Use R.C.H Cheng "The generation of Gamma variables with
		// non-integral shape parameters", Applied Statistics, (1977), 26, No. 1, p71-74

		ainv := math.Sqrt(2.0*alpha - 1.0)
		bbb := alpha - math.Log(4.0)
		ccc := alpha + ainv

		for {
			u1 := random()
			if !(1e-7 < u1 && u1 < .9999999) {
				continue
			}
			u2 := 1.0 - random()
			v := math.Log(u1/(1.0-u1)) / ainv
			x := alpha * math.Exp(v)
			z := u1 * u1 * u2
			r := bbb + ccc*v - x
			if r+MAGIC_CONST-4.5*z >= 0.0 || r >= math.Log(z) {
				return x * beta
			}
		}
	} else if alpha == 1.0 {
		u := random()
		for u <= 1e-7 {
			u = random()
		}
		return -math.Log(u) * beta
	} else { // alpha between 0.0 and 1.0 (exclusive)
		// Uses Algorithm of Statistical Computing - kennedy & Gentle
		var x float64
		for {
			u := random()
			b := (math.E + alpha) / math.E
			p := b * u
			if p <= 1.0 {
				x = math.Pow(p, 1.0/alpha)
			} else {
				x = -math.Log((b - p) / alpha)
			}
			u1 := random()
			if p > 1.0 {
				if u1 <= math.Pow(x, alpha-1.0) {
					break
				}
			} else if u1 <= math.Exp(-x) {
				break
			}
		}
		return x * beta
	}
}
