package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/kk-no/go-lishogi"
)

func main() {
	token := os.Getenv("LISHOGI_ACCESS_TOKEN")

	cli := lishogi.NewClient(lishogi.SetAccessToken(token))
	results := make(map[string]int, 100)

	for i := 1; i < 100; i++ {
		s, err := cli.Tournament.GetStanding("TOURNAMENT_ID", strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err)
		}
		if len(s.Players) == 0 {
			break // finished.
		}
		for _, player := range s.Players {
			for _, score := range player.Sheet.Scores {
				v := reflect.ValueOf(score)
				switch v.Kind() {
				case reflect.Int:
					if i := v.Int(); i <= 1 {
						continue
					}
					results[player.Team]++
				case reflect.Float32, reflect.Float64:
					if f := v.Float(); f <= 1 {
						continue
					}
					results[player.Team]++
				case reflect.Slice:
					results[player.Team]++
				default:
					fmt.Println("unknown type", score, v.Kind())
				}
			}
		}
	}
	fmt.Println(results)
}
