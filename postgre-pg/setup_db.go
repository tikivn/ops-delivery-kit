//go:build integration_test
// +build integration_test

package postgre_pg

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"

	"github.com/go-pg/pg/v10"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func SetupDBTest() (DockerDBConn *pg.DB, closeFunc func() error, err error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	cfg := struct {
		Address  string `json:"addr"`
		Database string `json:"db"`
		Username string `json:"user"`
		Password string `json:"password"`
	}{
		Database: "linehaul_test",
		Username: "linehaul",
		Password: "123456",
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	runCfg := &dockertest.RunOptions{
		Repository: "postgis/postgis",
		Tag:        "12-2.5-alpine",
		Env: []string{
			"POSTGRES_USER=" + cfg.Username,
			"POSTGRES_PASSWORD=" + cfg.Password,
			"POSTGRES_DB=" + cfg.Database,
			"listen_addresses = '*'",
		},
	}

	resource, err := pool.RunWithOptions(runCfg, func(hostConfig *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		hostConfig.AutoRemove = true
		hostConfig.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	closeFunc = resource.Close

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	resource.Expire(360)
	handleInterrupt(pool, resource)

	cfg.Address = resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	cfg.Address = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))

	if err := pool.Retry(func() error {
		DockerDBConn = pg.Connect(&pg.Options{
			Addr:     cfg.Address,
			Database: cfg.Database,
			User:     cfg.Username,
			Password: cfg.Password,
		})

		_, err := DockerDBConn.ExecContext(context.Background(), "SELECT 1")
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return DockerDBConn, closeFunc, err
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
