package session

type Session struct {
	Clock 				*Clock
	Track 				*Track
}

func NewSession() *Session {
	s := new(Session)
	s.Clock = NewClock()
	s.Track = NewTrack()

	return s
}

