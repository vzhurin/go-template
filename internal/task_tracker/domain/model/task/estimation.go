package task

type Estimation struct {
	estimation uint64
}

func NewEstimation(e uint64) Estimation {
	return Estimation{estimation: e}
}
