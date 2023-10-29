package comb

type (
	// Comb is a combinator function that takes an input and returns either an
	// output or an error.
	Comb[In, Out any] func(In) (Out, error)

	// MustComb is a function type that takes an input and returns an output,
	// panicking if there's an error.
	MustComb[In, Out any] func(In) Out

	// Binder is a function type that takes an input and returns a Comb that
	// depends on the provided input.
	Binder[In, Out any] func(In) Comb[In, Out]
)

// Then composes two Comb functions, where the output of the first becomes the
// input to the second.
func (c Comb[In, Out]) Then(next Comb[Out, Out]) Comb[In, Out] {
	return Compose(c, next)
}

// Bind combines a Comb function with a Binder function to create a new Comb
// function.
func (c Comb[In, Out]) Bind(b Binder[Out, Out]) Comb[In, Out] {
	return Bind(c, b)
}

// Must converts a Comb function into a MustComb function, panicking if an
// error occurs.
func (c Comb[In, Out]) Must() MustComb[In, Out] {
	return func(in In) Out {
		res, err := c(in)
		if err != nil {
			panic(err)
		}
		return res
	}
}

// Compose combines two Comb functions into a single function, passing the
// output of the first as the input to the second.
func Compose[In, Handoff, Out any](
	l Comb[In, Handoff], r Comb[Handoff, Out],
) Comb[In, Out] {
	return func(in In) (Out, error) {
		res, err := l(in)
		if err == nil {
			return r(res)
		}
		var zero Out
		return zero, err
	}
}

// Bind combines a Comb function and a Binder function into a new Comb
// function, utilizing the Binder to transform the output by instantiating a
// Comb depending on the Binder's input.
func Bind[In, Handoff, Out any](
	l Comb[In, Handoff], b Binder[Handoff, Out],
) Comb[In, Out] {
	return Compose(l, func(in Handoff) (Out, error) {
		return b(in)(in)
	})
}
