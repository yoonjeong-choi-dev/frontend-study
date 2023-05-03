package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"graphql/cards"
	"log"
)

func main() {
	// 스키마 가져오기
	schema, err := cards.Setup()
	if err != nil {
		panic(err)
	}

	// 요청 쿼리 설정
	queryKeys := []string{
		"value - A",
		"value - 1",
		"suit - Spades",
	}
	queries := []string{
		`
			{
				cards(value: "A") {
					value
					suit
				}
			}
		`,
		`
			{
				cards(value: "1") {
					value
					suit
				}
			}
		`,
		`
			{
				cards(suit: "Spades") {
					value
					suit
				}
			}
		`,
	}

	for i, query := range queries {
		params := graphql.Params{Schema: schema, RequestString: query}
		res := graphql.Do(params)
		if len(res.Errors) > 0 {
			log.Fatalf("failed to query graphql: %+v\n", err)
		}

		data, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Response with %s :\n%s \n", queryKeys[i], data)
	}

}
