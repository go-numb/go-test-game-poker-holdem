package board

import (
	"fmt"

	"github.com/go-numb/go-game-poker-texasholdem/game/hand"
	"github.com/google/uuid"
)

type Board struct {
	ID string

	Flop *hand.Flop
}

func New() *Board {
	return &Board{
		ID:   fmt.Sprintf("BOARD-%s", uuid.NewString()),
		Flop: nil,
	}
}

func (p Board) String() string {
	return fmt.Sprintf(
		`
		[BOARD DATA]
		ID:		%s

		%s
		`,
		p.ID,
		p.Flop,
	)
}
