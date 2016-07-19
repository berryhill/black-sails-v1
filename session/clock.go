package session

type Clock struct {
	IsPlaying 		bool       `json:"play"`
	Bpm 			int        `json:"bpm"`
}

func NewClock() *Clock {
	c := new(Clock)
	c.IsPlaying = false
	c.Bpm = 90

	return c
}

func (c *Clock) Work() {
	go func() {
		if c.IsPlaying == true {
			//TODO implement
		}
	}()
}

func (c *Clock)Play () {
	c.IsPlaying = true
}

func (c *Clock) Stop() {
	c.IsPlaying =false
}

func (c *Clock) SetBpm(bpm int) {
	c.Bpm = bpm
}

func (c *Clock) GetBpm() int {
	return c.Bpm
}

