package user

import (
	"fmt"

	"github.com/go-numb/go-game-poker-texasholdem/game/hand"
	"github.com/google/uuid"
)

type User struct {
	ID string

	// Status テーブルへの状態
	Status State

	// Activate ゲームへの参加状態
	Activate Activate

	// Chip 保有チップ総量
	Chip uint32

	// Hand 現在のカード状態であり、未知のカードである0値を含む
	Hand *hand.Hand
}

func New() *User {
	return &User{
		ID:       fmt.Sprintf("USER-%s", uuid.NewString()),
		Status:   State(0),
		Activate: Activate(0),
		Chip:     0,
		Hand:     nil,
	}
}

func (p User) String() string {
	return fmt.Sprintf(
		`
		[USER DATA]
		UUID:		%s
		Status:		%s
		Activate:	%s

		Chip:		%d
		%s
		`,
		p.ID,
		p.Status,
		p.Activate,
		p.Chip,
		p.Hand,
	)
}
