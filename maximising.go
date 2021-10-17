package main

import "math"

type MaximisingState struct {
	entries   Filleable
	transform Transformation
	maxSize   float64
	value     interface{}
}

func Maximising(transform Transformation) *MaximisingState {
	return &MaximisingState{
		Counting(), transform, math.Inf(-1), nil,
	}
}

func (state *MaximisingState) Fill(datum interface{}) {
	state.entries.Fill(datum)

	proj := state.transform(datum)

	if proj > state.maxSize {
		state.maxSize = proj
		state.value = datum
	}
}
