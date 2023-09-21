package hand

import (
	"fmt"

	"github.com/go-numb/go-game-poker-texasholdem/game/hand/card"
)

type Flop struct {
	// TableID 参加テーブル
	TableID string

	Cards [5]card.Card

	// 評価指数
	Role        string  // 手役
	Progress    uint8   // 消費ターン
	Score       float32 // 手役スコア
	Probability float32 // 勝率
}

func NewFlop(tableID string) *Flop {
	return &Flop{
		TableID:     tableID,
		Cards:       card.NewFlop(),
		Role:        "",
		Progress:    0,
		Score:       0,
		Probability: 0,
	}
}

func (p Flop) String() string {
	return fmt.Sprintf(
		`
		TableID:	%s
		Cards:		%s

		Role:		%s

		Progress:	%d
		Score:		%.2f
		Prob:		%.2f%%
	`,
		p.TableID,
		p.Cards,
		p.Role,
		p.Progress,
		p.Score,
		p.Probability*100,
	)
}
