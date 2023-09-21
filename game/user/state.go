package user

type State uint8

const (
	Inactive State = iota
	Active
	Leave
	Watch
)

func (p State) String() string {
	switch p {
	case Active:
		return "ACTIVE"
	case Leave:
		return "LEAVE"
	case Watch:
		return "WATCH"
	}

	return "INACTIVE"
}
