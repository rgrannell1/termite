package main

type SummingState struct {
	entries   Filleable
	transform Transformation
	values    float64
}

func Summing(transform Transformation) *SummingState {
	return &SummingState{
		Counting(), transform, 0,
	}
}

func (state *SummingState) Fill(datum float64) {
	state.values += state.transform(datum)
}
