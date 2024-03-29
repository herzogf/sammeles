package common

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

// heavily inspired by https://blog.devgenius.io/tooling-go-microservices-with-consul-service-discovery-and-kv-store-2c52bcdf4fd4
type ConsulClient struct {
	*api.Client
}

func NewClient(addr string) (*ConsulClient, error) {

	conf := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(conf)
	if err != nil {
		log.Println("error initiating new consul client: ", err)
		return &ConsulClient{}, err
	}

	return &ConsulClient{
		client,
	}, nil
}

func (c *ConsulClient) RegisterType(typeIdentifier TypeIdentifier, port int) error {
	hostname := getHostname()

	check := &api.AgentServiceCheck{
		Interval: "10s",
		Timeout:  "1s",
		DeregisterCriticalServiceAfter: "1m",
		HTTP:     fmt.Sprintf("http://%s:%d/health", hostname, port),
	}

	name := fmt.Sprintf("type/%s/%s/%d", typeIdentifier.Group, typeIdentifier.Type, typeIdentifier.SchemaVersion)
	id := name + "/" + hostname
	serviceDefinition := &api.AgentServiceRegistration{
		ID:    id,
		Name:  name,
		Address: hostname,
		Port: port,
		Tags:  []string{"sammeles", "type"},
		Meta: map[string]string{
			"kind": "type",
			"group": typeIdentifier.Group,
			"type": typeIdentifier.Type,
			"schemaVersion": fmt.Sprintf("%d", typeIdentifier.SchemaVersion),
		},
		Check: check,
	}
	if err := c.Agent().ServiceRegister(serviceDefinition); err != nil {
		log.Fatal("error registering service: ", err)
	}

	log.Println("registered service: ", serviceDefinition)

	return nil
}

func RegisterOneAndOnlyType(typeIdentifier TypeIdentifier, port int) {
	consulClient, err := NewClient(os.Getenv("CONSUL_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	if err := consulClient.RegisterType(typeIdentifier, port); err != nil {
		log.Fatal(err)
	}
}

func getHostname() (string) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

// the following function deregisters the type/service from consul
func (c *ConsulClient) DeregisterType(typeIdentifier TypeIdentifier) error {
	hostname := getHostname()
	id := fmt.Sprintf("type/%s/%s/%d/%s", typeIdentifier.Group, typeIdentifier.Type, typeIdentifier.SchemaVersion, hostname)
	if err := c.Agent().ServiceDeregister(id); err != nil {
		log.Fatal("error deregistering service: ", err)
	}

	log.Println("deregistered service: ", id)

	return nil
}

func DeregisterOneAndOnlyType(typeIdentifier TypeIdentifier) {
	consulClient, err := NewClient(os.Getenv("CONSUL_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	if err := consulClient.DeregisterType(typeIdentifier); err != nil {
		log.Fatal(err)
	}

	log.Printf("deregistered type '%s' from consul\n", typeIdentifier.Type)
}

// find a service based on the type identifier in ConsulClient
func (c *ConsulClient) FindTypeServiceUrl(typeIdentifier TypeIdentifier) (string, error) {
	name := fmt.Sprintf("type/%s/%s/%d", typeIdentifier.Group, typeIdentifier.Type, typeIdentifier.SchemaVersion)

	filterString := fmt.Sprintf("Service == \"%s\"", name)
	services := c.findServicesWithFilter(filterString)

	for _, service := range services {
		address := service.Address
		port := service.Port
		url := fmt.Sprintf("http://%s:%v/api/types/%s/%s", address, port, typeIdentifier.Group, typeIdentifier.Type)
	
		return url, nil
	}

	return "", fmt.Errorf("no services found for type '%s'", name)
}

func FindTypeServiceUrl(typeIdentifier TypeIdentifier) (string, error) {
	consulClient, err := NewClient(os.Getenv("CONSUL_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	return consulClient.FindTypeServiceUrl(typeIdentifier)
}

func (c *ConsulClient) findServicesWithFilter(filter string) map[string]*api.AgentService {
	services, err := c.Agent().ServicesWithFilter(filter)
	if err != nil {
		log.Fatal("error finding services in consul: ", err)
	}

	return services
}