package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromPath(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type 'NAME wins' to record a win")
	poker.NewCLI(
		store,
		os.Stdin,
		poker.BlindAlerterFunc(poker.StdOutAlerter),
	).PlayPoker()
}
