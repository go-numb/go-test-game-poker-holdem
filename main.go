package main

import (
	"fmt"
	"time"

	"github.com/go-numb/go-game-poker-texasholdem/game"
	"github.com/go-numb/go-game-poker-texasholdem/game/user"
	"github.com/rs/zerolog/log"
)

func main() {
	var (
		initialMoney float64 = 100_000
	)

	central := game.New(initialMoney)

	// Tableを作成する
	table1 := central.NewTable()

	// Userが来店
	var (
		user1           = user.New()
		deposit float64 = 1_000
	)
	// ユーザーからのデポジットを受け、チップに変え、ユーザに渡す
	central.DepositToChip(user1, deposit)

	// 参加者がテーブルに座る
	// [ERROR]チップが不足していればエラーを返す
	// [ERROR]テーブル参加者の上限ならばエラーを返す
	if err := table1.Accepts(user1); err != nil {
		log.Err(err)
	}

	// PHESE: DEAL 参加者にカードを配る
	table1.Deal()
	fmt.Println("PHESE: DEAL")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

	// PHESE: PREFLOP 参加者がチップを賭けるアクションを選択する
	table1.Preflop()
	fmt.Println("PHESE: PREFLOP")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

	// PHESE: FLOP ボードにカード5枚を秘匿し、内3枚を開示する
	table1.Flop()
	fmt.Println("PHESE: FLOP")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

	// PHESE: TURN 4枚目を開示し、ユーザーがアクションを行う
	table1.Turn()
	fmt.Println("PHESE: TURN")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

	// PHESE: RIVER 5枚目を開示し、ユーザーはアクションを行う
	table1.River()
	fmt.Println("PHESE: RIVER")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

	// PHESE: SHOWDOWN 手役を評価し、チップを移動する
	table1.Showdown()
	fmt.Println("PHESE: SHOWDOWN")
	fmt.Printf("central:\n%+v\n", central)
	fmt.Printf("table:\n%+v\n", table1)
	time.Sleep(time.Second * 3)

}
