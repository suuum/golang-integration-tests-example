package repository_test

import (
	"example/repository"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExampleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pg Suite")
}

// Scenario: Function add two values
//    When first value is 10 and second value is 20
//    It should return 30
var _ = Describe("The Example service - example 1", func() {
	var example repository.ExampleRepository = &repository.ExampleRepositoryStruct{}
	Context("Function add two values", func() {
		When("first value is 10 and second value is 20", func() {
			It("should return 30", func() {
				result := example.GetExampleTaxValue(4, 50)

				Expect(result).Should(BeEquivalentTo(107))
			})
		})
	})
})
