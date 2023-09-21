package card

// Evaluation 現在のカードを参照し手役、評価、勝率を提示する
func (p *Cards) Evaluation() (role string, progress uint8, score, prob float32) {
	// 進行状況を得る
	return role, uint8(len(p.Table)), score, prob
}
