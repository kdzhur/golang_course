package game

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"the_game/internal/player"
)

const GAME_NAME string = "The New World"

func LoadGame(p *player.Player) {

	g := Game{
		name:   GAME_NAME,
		Player: p,
	}
	fmt.Printf("Welcome to the %v game. Your name is %v and you are seeking for adventure!\n", g.name, g.Player.Name)

	var locationList []location

	jsonFile, err := os.Open("assets/script.json")
	if err != nil {
		log.Fatalln("failed to open file assets/script.json")
	}

	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	if err := json.Unmarshal(bytes, &locationList); err != nil {
		log.Fatalln("Failed to unmarshal: ", err)
	}

	var input int
	var event_id int

	for _, location := range locationList {
		fmt.Println(location.Description)
		for event_id = 0; event_id < len(location.Events); event_id++ {
			fmt.Println(location.Events[event_id].Description)
			if location.Events[event_id].The_end {
				os.Exit(0)
			}
			for _, choice := range location.Events[event_id].Choices {
				fmt.Println(choice.Description)
			}
			if len(location.Events[event_id].Choices) != 0 {
				fmt.Scan(&input)

				if location.Events[event_id].Choices[input-1].Reference != 0 {
					for index := 0; index < len(location.Events); index++ {
						if location.Events[event_id].Choices[input-1].Reference == location.Events[index].Id {
							event_id = index - 1
							break
						}
					}
				}
			}
		}
	}
}
