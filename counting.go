package main

type CountingState struct {
	entries   float64
	transform Transformation
}

func (state *CountingState) Fill(datum interface{}) {
	state.entries += state.transform(datum)
}

func Counting() *CountingState {
	return &CountingState{
		entries:   0,
		transform: Identity,
	}
}
