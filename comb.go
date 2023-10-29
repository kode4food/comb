package comb

type (
	Comp[In, Out any]     func(In) (Out, error)
	MustComp[In, Out any] func(In) Out
	Binder[In, Out any]   func(In) Comp[In, Out]
)

func (c Comp[In, Out]) Then(next Comp[Out, Out]) Comp[In, Out] {
	return Compose(c, next)
}

func (c Comp[In, Out]) Bind(b Binder[Out, Out]) Comp[In, Out] {
	return Bind(c, b)
}

func (c Comp[In, Out]) Must() MustComp[In, Out] {
	return func(in In) Out {
		res, err := c(in)
		if err != nil {
			panic(err)
		}
		return res
	}
}

func Compose[In, Handoff, Out any](
	l Comp[In, Handoff], r Comp[Handoff, Out],
) Comp[In, Out] {
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
	l Comp[In, Handoff], b Binder[Handoff, Out],
) Comp[In, Out] {
	return Compose(l, func(in Handoff) (Out, error) {
		return b(in)(in)
	})
}
