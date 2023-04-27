package service_discovery

import "github.com/hashicorp/consul/api"

// Client consul 패키지를 이용한 디스커버리 서비스(서비스 메시)에 등록할 서비스 인터페이스
type Client interface {
	Register(tags []string) error
	Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error)
}

// SimpleClient Client 인터페이스 구현체
type SimpleClient struct {
	client  *api.Client
	address string
	name    string
	port    int
}

func (c *SimpleClient) Register(tags []string) error {
	service := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Port:    c.port,
		Address: c.address,
		Tags:    tags,
	}

	return c.client.Agent().ServiceRegister(service)
}

func (c *SimpleClient) Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.client.Health().Service(service, tag, false, nil)
}

func NewClient(config *api.Config, address, name string, port int) (Client, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	sc := &SimpleClient{
		client:  client,
		name:    name,
		address: address,
		port:    port,
	}

	return sc, nil
}
