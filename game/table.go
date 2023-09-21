package game

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-numb/go-game-poker-texasholdem/game/board"
	"github.com/go-numb/go-game-poker-texasholdem/game/user"
	"github.com/google/uuid"
)

const (
	TABLEMAXIMUMUSERS = 8
)

// Table テーブル
type Table struct {
	TableID string
	// N テーブルが立って以降のゲーム回数
	N uint8

	// Phese ベッティング ラウンドフェーズ
	Phese Phese

	// Pot 賭けチップ総額/ゲームあたり
	Pot   uint32
	Board *board.Board
	// Users テーブルの参加者
	Users []*user.User
	// Bets ベッド履歴
	Bets []*Bet

	CreatedAt time.Time
}

func (p *Central) NewTable() *Table {
	tableID := fmt.Sprintf("TABLE-%s", uuid.NewString())
	p.AddTable(tableID)

	return &Table{
		TableID:   tableID,
		N:         0,
		Phese:     DEAL,
		Pot:       0,
		Board:     board.New(),
		Users:     make([]*user.User, 0),
		Bets:      make([]*Bet, 0),
		CreatedAt: time.Now(),
	}
}

// Accepts ユーザーがテーブルに着席する
func (p *Table) Accepts(u *user.User) error {
	// チップの有無をチェック
	if 10 > u.Chip {
		return errors.New(ErrInsufficientDeposit)
	}

	// テーブル参加者上限を確認する
	if TABLEMAXIMUMUSERS <= len(p.Users) {
		return errors.New(ErrLimitUsersNumber)
	}

	p.Users = append(p.Users, u)
	return nil
}

func (p *Table) AddBet(userID string, amount uint32) {
	t := time.Now()
	p.Bets = append(p.Bets, &Bet{
		ID:        fmt.Sprintf("BET-%d", t.UnixNano()),
		UserID:    userID,
		Amount:    amount,
		isStandby: true,
		BetAt:     t,
	})
}
