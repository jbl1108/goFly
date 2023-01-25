package driver

import (
	"context"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/jbl1108/goFly/restservice"
	"github.com/jbl1108/goFly/util"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	go restservice.Start()
	//	err := startRedis()
	//	if err != nil {
	//		log.Println(err)
	//	}
}

func shutdown() {
	//	err := stopRedis()
	//	if err != nil {
	//		log.Println(err)
	//	}
}

func startRedis() error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "redis",
		ExposedPorts: nat.PortSet{"6379": struct{}{}},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("6379"): {{HostIP: "127.0.0.1", HostPort: "6379"}}},
	}, nil, nil, "redis-test")
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func stopRedis() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	if err := cli.ContainerStop(ctx, "redis-test", nil); err != nil {
		return err
	}

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := cli.ContainerRemove(ctx, "redis-test", removeOptions); err != nil {
		return err
	}
	return nil

}

func TestGoFlyRequest(t *testing.T) {
	var restClient = NewRestClient()
	var m = make(map[string]string)
	result, err := restClient.Request("http://localhost:8000/gofly", m)

	resultMap := result.(map[string]interface{})["response"].([]interface{})[0].(map[string]interface{})

	if err != nil {
		t.Fatal("error:", err.Error())
	}
	if resultMap["code"].(string) != "CDG" {
		t.Error("code error")
	}

}

func TestFlightRequest(t *testing.T) {
	var restClient = NewRestClient()
	result, err1 := restClient.Request("http://localhost:8000/flight", map[string]string{})
	if err1 != nil {
		t.Fatal(err1)
	}
	var parser = NewFlightDataParser()
	parsed, err2 := parser.ParseData(result)
	if err2 != nil {
		t.Error(err2)
	}
	if len(parsed) != 2 {
		t.Error(parsed)
	}
}

func TestRedisString(t *testing.T) {
	conf := util.NewConfig()
	var redisClient = NewRedisDriver(conf)
	redisClient.StoreString("test", "value1")
	value, err := redisClient.FetchString("test")
	if err != nil {
		t.Error(err)
	}
	if value != "value1" {
		t.Errorf("value1 is not right -%s-", value)
	}
}

func TestRedisList(t *testing.T) {
	conf := util.NewConfig()
	var redisClient = NewRedisDriver(conf)
	b := []string{"value1", "value2", "value3"}
	redisClient.StoreList("test", b)
	value, err := redisClient.FetchList("test")
	if err != nil {
		t.Error(err)
	}
	if len(value) != 3 {
		t.Errorf("value1 is not right %s", value)
	}
}
