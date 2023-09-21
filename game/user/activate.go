package user

type Activate uint8

const (
	Hold Activate = iota
	Fold
	Check
	Call
	Rise
	Allin
)

func (p Activate) String() string {
	switch p {
	case Fold:
		return "FOLD"
	case Check:
		return "CHECK"
	case Call:
		return "CALL"
	case Rise:
		return "RISE"
	case Allin:
		return "ALL-IN"
	}

	return "HOLD"
}
