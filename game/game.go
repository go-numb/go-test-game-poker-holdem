package game

import (
	"time"

	"github.com/go-numb/go-game-poker-texasholdem/game/user"
)

// Central 管理監督（当日区切りなど
type Central struct {
	// 運営総額
	Pool float64
	// Deposit お客様からお預かりしたチップ総額
	Deposit float64
	// テーブル管理
	TableID []string

	CreatedAt time.Time
	EndedAt   time.Time
}

func New(initialMoney float64) *Central {
	return &Central{
		Pool: initialMoney,

		TableID:   make([]string, 0),
		CreatedAt: time.Now(),
		EndedAt:   time.Now(),
	}
}

func (p *Central) AddTable(tableID string) {
	p.TableID = append(p.TableID, tableID)
}

// DepositToChip ユーザーからのデポジットを受け、チップに変え、ユーザに渡す
func (p *Central) DepositToChip(u *user.User, deposit float64) {
	// デポジットを受領し
	p.Deposit += deposit
	// チップに変換
	chip := p.toChip(deposit)
	// ユーザーに付与
	u.Chip = chip
}

// toChip USDなどからチップに変換する関数
func (p *Central) toChip(deposit float64) uint32 {
	return uint32(deposit)
}
