package cards

import (
	"github.com/graphql-go/graphql"
	"strings"
)

// CardResolve graphql 요청에 대한 응답 반환
func CardResolve(p graphql.ResolveParams) (interface{}, error) {
	var ret []Card

	suit, suitOk := p.Args["suit"].(string)
	suit = strings.ToLower(suit)

	value, valueOk := p.Args["value"].(string)
	value = strings.ToLower(value)

	// 저장소의 카드 정보들 중에 요청 정보와 일치하는 카드들만 필터링
	for _, card := range cards {
		if suitOk && suit != strings.ToLower(card.Suit) {
			continue
		}
		if valueOk && value != strings.ToLower(card.Value) {
			continue
		}

		ret = append(ret, card)
	}

	return ret, nil
}
