package card

import (
	"fmt"
	"os"
	"testing"
)

const (
	TESTCOUNT      = 13000
	TESTRESULTFILE = "./test.csv"
)

func TestCards(t *testing.T) {
	count := TESTCOUNT
	cards := make([]Card, count)

	for i := 0; i < count; i++ {
		cards[i] = NewCard()
	}

	f, _ := os.Create(TESTRESULTFILE)
	defer f.Close()

	for i, card := range cards {
		f.WriteString(fmt.Sprintf("%d, %d, %d\n", i, card.ID, card.Mark))
	}

}
