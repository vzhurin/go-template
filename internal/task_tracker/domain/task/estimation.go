package task

type Estimation struct {
	estimation uint64
}

func NewEstimation(e uint64) Estimation {
	return Estimation{estimation: e}
}

func (e Estimation) Uint64() uint64 {
	return e.estimation
}
