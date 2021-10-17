package main

import "math"

type BinningState struct {
	bins      int64
	low       float64
	high      float64
	transform Transformation
	underflow Filleable
	overflow  Filleable
	nanflow   Filleable
	entries   Filleable
	values    []Filleable
}

func (state *BinningState) Fill(datum interface{}) {
	y := state.transform(datum)
	state.entries.Fill(datum)

	if math.IsNaN(y) {
		state.nanflow.Fill(datum)
	} else if y < state.low {
		state.underflow.Fill(datum)
	} else if y > state.high {
		state.overflow.Fill(datum)
	} else {
		bin := int(math.Floor(float64(state.bins) * (y - state.low) / (state.high - state.low)))
		state.values[bin].Fill(datum)
	}
}

func Binning(bins int64, low, high float64) *BinningState {
	aggs := make([]Filleable, bins)

	for idx := range aggs {
		aggs[idx] = Counting()
	}

	return &BinningState{
		bins:      bins,
		low:       low,
		high:      high,
		transform: Identity,
		underflow: Counting(),
		overflow:  Counting(),
		nanflow:   Counting(),
		entries:   Counting(),
		values:    aggs,
	}
}
