package code_analyzer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser tests", func() {
	It("should parse function", func() {
		parser := Parser{}
		code, err := parser.Parse("test.go")
		Ω(err).Should(BeNil())
		Ω(code).Should(BeNil())
	})
})
