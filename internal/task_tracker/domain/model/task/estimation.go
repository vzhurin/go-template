package task

import "strconv"

type Estimation struct {
	estimation uint64
}

func NewEstimation(e uint64) Estimation {
	return Estimation{estimation: e}
}

func (e Estimation) String() string {
	return strconv.FormatUint(e.estimation, 10)
}
