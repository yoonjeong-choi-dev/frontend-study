package storage

import "context"

type Item struct {
	Name   string `bson:"name"`
	Price  int64  `bson:"price"`
	OnSale bool   `bson:"onSale"`
}

// ItemStorage Item 정보를 저장하는 데이터 저장소의 인터페이스
// 해당 인터페이스 구현체가 사용하는 DB 정보는 몰라도 된다
// 또한 Mock DB 구현체를 통한 유닛 테스트가 가능
type ItemStorage interface {
	GetByName(context.Context, string) (*Item, error)
	Save(context.Context, *Item) error
}
