package eval

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {

	}
}
