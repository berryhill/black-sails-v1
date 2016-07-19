package session

type Quantizer struct {
	Fire 			bool
	Length		 	[]int        `json:"length"`
}

func (q *Quantizer) GetLength() []int {
	return q.Length
}

func (q *Quantizer) SetLength(numerator int, denominator int) {
	var length []int
	length = append(length, numerator)
	length = append(length, denominator)
	q.Length = length
}

func (q *Quantizer) GetFire() bool {
	return q.Fire
}

func (q *Quantizer) SetFire(b bool) {
	q.Fire = b
}

