package card

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
)

// Cards ゲームテーブル5枚と各参加者であるカードを内包し、ゲーム進行度でCardNaN有り
type Cards struct {
	Table [5]Card
	User  [2]Card
}

type Card struct {
	ID   Number
	Mark Mark
}

type Number uint8

const (
	// NaN 未配のカード
	NumberNaN Number = iota

	Number1
	Number2
	Number3
	Number4
	Number5
	Number6
	Number7
	Number8
	Number9
	Number10
	NumberJ
	NumberQ
	NumberK
	_
	// NumberUnknown equivalent to JOKER
	NumberUnknown
)

type Mark uint8

const (
	MarkNaN Mark = iota

	MarkHeart
	MarkDia
	MarkSpade
	MarkClover
)

func New() *Cards {
	return &Cards{
		Table: [5]Card{},
		User:  [2]Card{},
	}
}

func NewCard() Card {
	return Card{
		ID:   Number(uint8(r.Intn(13) + 1)),
		Mark: Mark(uint8(r.Intn(4) + 1)),
	}
}

// NewFlop ユーザー用の参加カードを配る
func NewFlop() [5]Card {
	return [5]Card{
		NewCard(),
		NewCard(),
		NewCard(),
		NewCard(),
		NewCard(),
	}
}

// NewUser ユーザー用の参加カードを配る
func NewUser() [2]Card {
	return [2]Card{
		NewCard(),
		NewCard(),
	}
}

func (p Card) String() string {
	var (
		number, mark string
	)

	switch p.ID {
	case 11:
		number = "J"
	case 12:
		number = "Q"
	case 13:
		number = "K"
	default:
		number = fmt.Sprintf("%d", p.ID)
	}

	switch p.Mark {
	case MarkHeart:
		mark = "HEART"
	case MarkDia:
		mark = "DIAMOND"
	case MarkSpade:
		mark = "SPADE"
	case MarkClover:
		mark = "CLOVER"
	}

	return fmt.Sprintf("%s-%s", mark, number)
}
