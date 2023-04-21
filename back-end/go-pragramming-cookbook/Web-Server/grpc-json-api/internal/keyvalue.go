package internal

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"grpc-json-api/keyvalue"
	"sync"
)

type KeyValue struct {
	mutex sync.RWMutex
	kvm   map[string]string
	keyvalue.UnimplementedKeyValueServer
}

func (k *KeyValue) Set(ctx context.Context, r *keyvalue.SetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.Lock()
	k.kvm[r.GetKey()] = r.GetValue()
	k.mutex.Unlock()

	return &keyvalue.KeyValueResponse{
		Value: r.GetValue(),
	}, nil
}

func (k *KeyValue) Get(ctx context.Context, r *keyvalue.GetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	val, ok := k.kvm[r.GetKey()]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "no key")
	}
	return &keyvalue.KeyValueResponse{Value: val}, nil
}

func NewKeyValue() *KeyValue {
	return &KeyValue{
		kvm: make(map[string]string),
	}
}
