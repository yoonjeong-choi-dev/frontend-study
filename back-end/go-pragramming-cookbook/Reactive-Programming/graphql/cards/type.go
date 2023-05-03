package cards

import "github.com/graphql-go/graphql"

// CardType graphql 쿼리에 사용하기 위해 저장소에 저장된 card 객체 반환
func CardType() *graphql.Object {
	cardType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Card",
		Description: "A Playing Card",
		Fields: graphql.Fields{
			"value": &graphql.Field{
				Type:        graphql.String,
				Description: "Ace through King",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if card, ok := p.Source.(Card); ok {
						return card.Value, nil
					}
					return nil, nil
				},
			},
			"suit": &graphql.Field{
				Type:        graphql.String,
				Description: "Spades, Diamonds, Clubs, Hearts",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if card, ok := p.Source.(Card); ok {
						return card.Suit, nil
					}
					return nil, nil
				},
			},
		},
	})

	return cardType
}
