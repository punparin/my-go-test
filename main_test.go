package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/mock"
	mocks "test/mocks"
)

var _ = Describe("main", func() {
	var (
		mockTest	*mocks.Test
	)

	BeforeEach(func() {
		mockTest = &mocks.Test{}
	})

	Context("Test B", func() {
		It("is 6", func() {
			mockTest.On("A",
				mock.AnythingOfType("int")).Return(
				4,
			)

			n := b(mockTest.A, 3)
			Expect(n).Should(BeEquivalentTo(6))
		})
	})
})
