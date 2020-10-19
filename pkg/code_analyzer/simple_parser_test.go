package code_analyzer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Parser tests", func() {
	var parser *SimpleParser

	BeforeEach(func() {
		parser = NewSimpleParser()
	})
	FIt("should parse interface, func()", func() {
		result, err := parser.Parse("test_data/interface.go")
		Ω(err).Should(BeNil())
		Ω(result).ShouldNot(BeNil())
	})
})
