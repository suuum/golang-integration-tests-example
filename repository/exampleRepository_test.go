package repository_test

import (
	"example/database"
	dockerTestInit "example/init-docker-tests"
	r "example/repository"
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
)

func TestExampleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pg Suite")
}

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
	repo     r.ExampleRepository
)

var _ = BeforeSuite(func() {
	// Init container, open connection, run migrations seed database, init repository
	rand.Seed(time.Now().UnixNano())
	exposedPort := fmt.Sprint(rand.Intn(10000))
	pool, resource = dockerTestInit.InitTestDocker(exposedPort)
	gdb := dockerTestInit.OpenDatabaseConnection(pool, resource, exposedPort)
	database.RunMigrations(gdb)
	dockerTestInit.SeedDatabase(gdb, pool, resource)
	repo = r.CreateExampleRepositoryRepository(gdb)
})

// Scenario: Function fetch first matching element
//
//	When first prop is equal to 1 and secondProp is equal to
//	It should return one row
var _ = Describe("The Example repository", func() {
	Context("Function fetch first matching element", func() {
		When("first prop is equal to 1 and secondProp is equal to 2", func() {
			It("should return one row", func() {
				result := repo.GetEntity(1, 2)

				Expect(result).ShouldNot(BeNil())
				Expect(result.FistProp).Should(BeEquivalentTo(1))
				Expect(result.SecondProp).Should(BeEquivalentTo(2))
			})
		})
	})
})

var _ = AfterSuite(func() {
	// Purge function destroys container
	if err := pool.Purge(resource); err != nil {
		logrus.Error(err)
	}
})
