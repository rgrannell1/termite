package main

import "math"

type MinimisingState struct {
	entries   Filleable
	transform Transformation
	minSize   float64
	value     interface{}
}

func Minimising(transform Transformation) *MinimisingState {
	return &MinimisingState{
		Counting(), transform, math.Inf(1), nil,
	}
}

func (state *MinimisingState) Fill(datum interface{}) {
	state.entries.Fill(datum)

	proj := state.transform(datum)

	if proj < state.minSize {
		state.minSize = proj
		state.value = datum
	}
}
