package game

import "github.com/go-numb/go-game-poker-texasholdem/game/hand"

type Phese uint8

const (
	DEAL     Phese = iota
	PREFLOP        // 1st
	FLOP           // 2nd
	TURN           // 3rd
	RIVER          // 4th
	SHOWDOWN       // 5th showdown
)

// Deal テーブルディーラーが参加者にカードを配る
func (p *Table) Deal() {
	p.Phese = DEAL
	// TODO: 配る順番
	for i := 0; i < len(p.Users); i++ {
		p.Users[i].Hand = hand.New(p.TableID)
	}
}

// Preflop 参加者がチップを賭ける
func (p *Table) Preflop() {
	p.Phese = PREFLOP

	// TODO: 各ユーザーがアクションを行う

}

// Flop ボードにカード5枚を秘匿し、内3枚を開示する
func (p *Table) Flop() {
	p.Phese = FLOP

	p.Board.Flop = hand.NewFlop(p.TableID)
}

// Turn 4枚目を開示し、ユーザーはアクションを行う
func (p *Table) Turn() {
	p.Phese = TURN

	// TODO: ボードの4枚目のカードを開く

	// TODO: ユーザがアクションを行う

}

// River 5枚目を開示し、ユーザーはアクションを行う
func (p *Table) River() {
	p.Phese = RIVER

	// TODO: ボードの4枚目のカードを開く

	// TODO: ユーザがアクションを行う

}

// Showdown 手役を評価し、チップを移動する
func (p *Table) Showdown() {
	p.Phese = SHOWDOWN

	// TODO: 評価する

	// TODO: チップを移動する

}
