package session

type Stem struct {
	Index 				int
	Grid 				*Grid           	`json:"grid"`
	IsPlaying 			bool            	`json:"is_playing"`
	Loop 				bool           		`json:"loop"`
}

func NewStem(index int) *Stem {
	s := new(Stem)
	s.Index = index

	return s
}

func (s *Stem) OpenFile(location string) error {
	//TODO implement

	return nil
}

func (s *Stem) Load() error {
	//TODO implement

	return nil
}

func (s *Stem) Play() {
	//TODO implement
}

func (s *Stem) Stop() {
	//TODO implement
}

func (s *Stem) MovePlayPosition(Bars int, Steps int) {
	//TODO implement
}

func (s *Stem) LoopOn(bars int) {
	//TODO implement
}

func (s *Stem) LoopOff() {
	//TODO implement
}

