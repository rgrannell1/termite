package main

type Transformation = func(val interface{}) float64

type Filleable interface {
	Fill(datum interface{})
}

func Identity(val interface{}) float64 {
	return val.(float64)
}

func Compose(left, right Transformation) Transformation {
	composed := func(datum interface{}) float64 {
		return left(right(datum))
	}

	return composed
}
