package service_test

import (
	"testing"

	"example/repository/mocks"
	"example/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestExampleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pg Suite")
}

type TestVerification struct {
	A      int
	B      int
	Result int
}

// Scenario: Function add two values
//    When first value is 10 and second value is 20
//    It should return 30
var _ = Describe("The Example service - example 1", func() {
	var example service.ExampleService = &service.ExampleServiceStruct{}
	Context("Function add two values", func() {
		When("first value is 10 and second value is 20", func() {
			It("should return 30", func() {
				result := example.Add(10, 20)

				Expect(result).Should(BeEquivalentTo(30))
			})
		})
	})
})

var _ = Describe("The Example service  - example 2", func() {
	var example service.ExampleService = &service.ExampleServiceStruct{}
	Context("Function add two values with array verification", func() {
		testCases := []TestVerification{
			{A: 10, B: 20, Result: 30},
			{A: 20, B: 20, Result: 40},
			{A: 30, B: 20, Result: 50},
			{A: 40, B: 50, Result: 90},
			{A: 10, B: 60, Result: 70},
			{A: 10, B: 80, Result: 90},
		}

		When("values to add are set in array", func() {
			It("should return expected result from array", func() {
				for _, testCase := range testCases {
					result := example.Add(testCase.A, testCase.B)
					Expect(result).Should(BeEquivalentTo(testCase.Result))
				}

			})
		})
	})
})

var _ = Describe("The Example service  - example 3", func() {
	var example service.ExampleService = &service.ExampleServiceStruct{}

	Context("Function add two values with tax value from DB", func() {
		repositoryMock := new(mocks.ExampleRepository)
		repositoryMock.On("GetExampleTaxValue", mock.Anything, mock.Anything).Return(7)

		// if you want to check exact paramters then you should replace them with real values
		// repositoryMock.On("GetExampleTaxValue", 1, 2).Return(7)
		example = service.InitExampleService(repositoryMock)

		When("first value have value 50, second 50 and tax rate is 7% ", func() {
			It("should return sum of two values with tax from DB", func() {
				result := example.AddWithTaxValueFromDB(50, 50)

				Expect(result).Should(BeEquivalentTo(107))
			})
		})
	})
})
