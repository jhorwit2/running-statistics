package runningstats

import "math"

// This code is a port of http://www.johndcook.com/blog/skewness_kurtosis/

// RunningStats struct
type RunningStats struct {
	n              int
	m1, m2, m3, m4 float64
	min, max       float64
}

// NewRunningStats creates a new running stat object
func NewRunningStats() *RunningStats {
	return &RunningStats{}
}

// Clear the running statistic for new data
func (r *RunningStats) Clear() {
	r.n = 0
	r.m1 = 0.0
	r.m2 = 0.0
	r.m3 = 0.0
	r.m4 = 0.0
	r.min = 0.0
	r.max = 0.0
}

// AddAll adds an array of data to the running stats
func (r *RunningStats) AddAll(data []float64) {
	for _, x := range data {
		r.Add(x)
	}
}

// Add value to RunningStats regression
func (r *RunningStats) Add(x float64) {
	if x > r.max {
		r.max = x
	} else if x < r.min {
		r.min = x
	}

	n1 := float64(r.n)
	r.n++
	delta := x - r.m1
	deltaN := delta / n1
	deltaN2 := deltaN * deltaN
	term1 := delta * deltaN * n1
	r.m1 += deltaN
	r.m4 += term1*deltaN2*float64(r.n*r.n-3*r.n+3) + 6*deltaN2*r.m2 - 4*deltaN*r.m3
	r.m3 += term1*deltaN*(n1-2) - 3*deltaN*r.m2
	r.m2 += term1
}

// Len returns the number of observations
func (r *RunningStats) Len() int {
	return r.n
}

// Min returns the number of observations
func (r *RunningStats) Min() float64 {
	return r.min
}

// Max returns the number of observations
func (r *RunningStats) Max() float64 {
	return r.max
}

// Mean returns the mean of the observations
func (r *RunningStats) Mean() float64 {
	return r.m1
}

// Var returns the variance of the observations
func (r *RunningStats) Var() float64 {
	return r.m2 / float64(r.n-1)
}

// Stddev returns the standard deviation of the observations
func (r *RunningStats) Stddev() float64 {
	return math.Sqrt(r.Var())
}

// Skewness returns the skewness value for the observations
func (r *RunningStats) Skewness() float64 {
	return math.Sqrt(float64(r.n)) * r.m3 / math.Pow(r.m2, 1.5)
}

// Kurtosis returns the kurtosis value for the observations
func (r *RunningStats) Kurtosis() float64 {
	return float64(r.n)*r.m4/(r.m2*r.m2) - 3.0
}
