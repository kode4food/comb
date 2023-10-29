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

func Bind[In, Handoff, Out any](
	l Comp[In, Handoff], b Binder[Handoff, Out],
) Comp[In, Out] {
	return func(i In) (Out, error) {
		res, err := l(i)
		if err == nil {
			return b(res)(res)
		}
		var zero Out
		return zero, err
	}
}

func Compose[In, Handoff, Out any](
	l Comp[In, Handoff], r Comp[Handoff, Out],
) Comp[In, Out] {
	return Bind(l, func(_ Handoff) Comp[Handoff, Out] {
		return func(h Handoff) (Out, error) {
			return r(h)
		}
	})
}
