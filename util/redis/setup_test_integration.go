//go:build integration_test
// +build integration_test

package redis

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func SetupRedisTest(ctx context.Context) (client *redis.Client, cleanUp func() error, err error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	runCfg := &dockertest.RunOptions{
		Repository: "redis",
		Tag:        "latest",
		Env:        []string{},
		Cmd: []string{
			"redis-server", "--port", "6379",
		},
	}

	resource, err := pool.RunWithOptions(runCfg, func(hostConfig *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		hostConfig.AutoRemove = true
		hostConfig.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	cleanUp = resource.Close

	// resource.Expire(360)
	handleInterrupt(pool, resource)

	// addr := resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	addr := net.JoinHostPort(resource.GetBoundIP("6379/tcp"), resource.GetPort("6379/tcp"))

	if err = pool.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr: addr,
		})

		return client.Ping(ctx).Err()
	}); err != nil {
		log.Fatalf("Cannot connect to redis with address %s due to %+v", addr, err)
	}

	return client, cleanUp, nil
}

func handleInterrupt(pool *dockertest.Pool, container *dockertest.Resource) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := pool.Purge(container); err != nil {
			log.Fatalf("Could not purge container: %s", err)
		}
		os.Exit(0)
	}()
}
