package cards

import "github.com/graphql-go/graphql"

// Setup 저장소에 저장되어 있는 카드 정보와 리졸버를 통해 스키마 생성
func Setup() (graphql.Schema, error) {
	cardType := CardType()

	fields := graphql.Fields{
		"cards": &graphql.Field{
			Type: graphql.NewList(cardType),
			Args: graphql.FieldConfigArgument{
				"suit": &graphql.ArgumentConfig{
					Description: "Filter cards by card suit (Spades, Diamonds, Clubs, Hearts)",
					Type:        graphql.String,
				},
				"value": &graphql.ArgumentConfig{
					Description: "Filter cards by card value (A-K)",
					Type:        graphql.String,
				},
			},
			Resolve: CardResolve,
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
