package comb

type (
	Comb[In, Out any]     func(In) (Out, error)
	MustComb[In, Out any] func(In) Out
	Binder[In, Out any]   func(In) Comb[In, Out]
)

func (c Comb[In, Out]) Then(next Comb[Out, Out]) Comb[In, Out] {
	return Compose(c, next)
}

func (c Comb[In, Out]) Bind(b Binder[Out, Out]) Comb[In, Out] {
	return Bind(c, b)
}

func (c Comb[In, Out]) Must() MustComb[In, Out] {
	return func(in In) Out {
		res, err := c(in)
		if err != nil {
			panic(err)
		}
		return res
	}
}

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

func Bind[In, Handoff, Out any](
	l Comb[In, Handoff], b Binder[Handoff, Out],
) Comb[In, Out] {
	return Compose(l, func(in Handoff) (Out, error) {
		return b(in)(in)
	})
}
