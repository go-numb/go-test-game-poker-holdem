package game

import "time"

type Bet struct {
	ID string
	// UserID ベットしたUserのID
	UserID string

	// Amount ベット額: 単位チップ故に整数
	Amount uint32

	// isStandby Pool化などの処理の有無
	isStandby bool

	BetAt time.Time
}

// IsStandby ベット待機状態を返す
func (p *Bet) IsStandby() bool {
	return p.isStandby
}
