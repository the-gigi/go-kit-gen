package code_analyzer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
	"strconv"
)

var _ = Describe("Simple Parser tests", func() {
	var parser *SimpleParser

	BeforeEach(func() {
		parser = NewSimpleParser()
	})
	It("should parse interface, func()", func() {
		result, err := parser.Parse("../test_data/interface.go")
		Ω(err).Should(BeNil())
		Ω(result).ShouldNot(BeNil())

		Ω(result.Name).Should(Equal("Foo"))
		Ω(result.Methods).Should(HaveLen(3))

		for i, method := range result.Methods {
			Ω(method.Name).Should(Equal("Op" + strconv.Itoa(i + 1)))
			Ω(method.Arguments).Should(HaveLen(1))
			Ω(method.Arguments[0].Name).Should(Equal("r"))
			Ω(method.Arguments[0].Type).Should(Equal(fmt.Sprintf("*Op%dRequest", i + 1)))

			Ω(method.Result).Should(HaveLen(2))
			Ω(method.Result[0].Type).Should(Equal(fmt.Sprintf("*Op%dResponse", i + 1)))
			Ω(method.Result[1].Type).Should(Equal("error"))
		}
	})
})
