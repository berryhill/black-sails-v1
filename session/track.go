// Maybe call it bus?
package session

type Track struct {
	Stems 				[]*Stem            `json:"stems"`
}

func NewTrack() *Track {
	t := new(Track)
	for k := 0; k < 4; k++ {
		s := NewStem(k)
		t.Stems = append(t.Stems, s)
	}

	return t
}

func (t *Track) Play() error {
	for k := 0; k < 4; k++ {
		go t.Stems[k].Play()
	}

	return nil
}

func (t *Track) Stop() error {
	for k := 0; k < 4; k++ {
		go t.Stems[k].Stop()
	}

	return nil
}

func (t *Track) LoopOn(bars int) error {
	for k := 0; k < 4; k++ {
		go t.Stems[k].LoopOn(bars)
	}

	return nil
}

func (t *Track) LoopOff() error {
	for k := 0; k < 4; k++ {
		go t.Stems[k].LoopOff()
	}

	return nil
}