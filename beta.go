// based on git@github.com:leesper/go_rng.git
package rand

import (
	"fmt"
)

func _beta(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	// according to "Numerical Recipes", all distinct
	// gamma variates must have different random seeds
	// MK : as we use a pool to back Gamma this should ok
	x := gamma(alpha, 1, random)
	y := gamma(beta, 1, random)
	return x / (x + y)
}

func calculateBeta(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	if !(alpha > 0.0) {
		panic(fmt.Sprintf("Invalid parameter alpha: %.2f", alpha))
	}
	if !(beta > 0.0) {
		panic(fmt.Sprintf("Invalid parameter beta: %.2f", beta))
	}
	return _beta(alpha, beta, random)
}

// Beta returns a random number of beta distribution (alpha > 0.0 and beta > 0.0)
// Uses package's built-in float64 randomization routine.
func Beta(alpha, beta float64) float64 {
	return calculateBeta(alpha, beta, Float64)
}

// Beta returns a random number of beta distribution (alpha > 0.0 and beta > 0.0)
// Uses externally provided float64 randomization routine.
func BetaEx(alpha, beta float64, random GetRandomFloat64Fn) float64 {
	return calculateBeta(alpha, beta, random)
}
