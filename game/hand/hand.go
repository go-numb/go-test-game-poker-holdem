package hand

import (
	"fmt"

	"github.com/go-numb/go-game-poker-texasholdem/game/hand/card"
)

type Hand struct {
	// TableID 参加テーブル
	TableID string

	Cards [2]card.Card

	// 評価指数
	Role        string  // 手役
	Progress    uint8   // 消費ターン
	Score       float32 // 手役スコア
	Probability float32 // 勝率
}

func New(tableID string) *Hand {
	return &Hand{
		TableID:     tableID,
		Cards:       card.NewUser(),
		Role:        "",
		Progress:    0,
		Score:       0,
		Probability: 0,
	}
}

// Update テーブルの更新を得て、各Userの手役の評価を更新する
func (p *Hand) Update(tableCards [5]card.Card) {
	cards := card.New()
	cards.User = p.Cards
	cards.Table = tableCards

	p.Role, p.Progress, p.Score, p.Probability = cards.Evaluation()
}

func (p Hand) String() string {
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
