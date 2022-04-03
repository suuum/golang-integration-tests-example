package dockerTestInit

import (
	"fmt"
	"log"
	"time"

	seeds "example/seeds"

	dockertest "github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitTestDocker function initialize docker with postgres image used for integration tests
func InitTestDocker(exposedPort string) (*dockertest.Pool, *dockertest.Resource) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	var passwordEnv = "POSTGRES_PASSWORD=postgres"
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"listen_addresses = '*'",
			fmt.Sprint(passwordEnv),
		},
		ExposedPorts: []string{exposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: exposedPort},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := resource.Expire(120); err != nil { // Tell docker to hard kill the container in 120 seconds
		logrus.Error(err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return pool, resource
}

func OpenDatabaseConnection(pool *dockertest.Pool, resource *dockertest.Resource, exposedPort string) *gorm.DB {
	user := "postgres"
	password := "postgres"
	db := "postgres"
	port := "5432"
	dns := "host=%s port=%s user=%s sslmode=disable password=%s dbname=%s"

	retries := 10
	host := resource.GetBoundIP(fmt.Sprintf("%s/tcp", port))
	gdb, err := gorm.Open(postgres.Open(fmt.Sprintf(dns, host, exposedPort, user, password, db)), &gorm.Config{})
	for err != nil {
		if retries > 1 {
			retries--
			time.Sleep(1 * time.Second)
			gdb, err = gorm.Open(postgres.Open(fmt.Sprintf(dns, host, exposedPort, user, password, db)), &gorm.Config{})
			continue
		}

		if err := pool.Purge(resource); err != nil {
			logrus.Error(err)
		}

		log.Panic("Fatal error in connection: ", err, resource.GetBoundIP("5432/tcp"))
	}

	return gdb
}

func SeedDatabase(gdb *gorm.DB, pool *dockertest.Pool, resource *dockertest.Resource) {
	for _, seed := range seeds.All() {
		if err := seed.Run(gdb); err != nil {
			if err := pool.Purge(resource); err != nil {
				logrus.Error(err)
			}

			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
