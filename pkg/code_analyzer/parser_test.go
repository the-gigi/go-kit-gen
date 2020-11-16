package code_analyzer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser tests", func() {
	var parser Parser

	BeforeEach(func() {
		parser = Parser{}
	})
	It("should parse function", func() {
		code, err := parser.Parse("../test_data/function.go")
		Ω(err).Should(BeNil())
		Ω(code.Interfaces).Should(BeNil())
		Ω(code.Functions).ShouldNot(BeNil())
		Ω(code.Functions).Should(HaveLen(1))
		x := code.Functions[0]
		Ω(x.Name).Should(Equal("Increment"))
		Ω(x.Arguments).Should(HaveLen(1))
		Ω(x.Arguments[0].Name).Should(Equal("x"))
		Ω(x.Arguments[0].Type).Should(Equal("int"))
		Ω(x.Result).Should(HaveLen(1))
		Ω(x.Result[0].Name).Should(BeEmpty())
		Ω(x.Result[0].Type).Should(Equal("int"))
	})
	It("should parse interface", func() {
		code, err := parser.Parse("test_data/interface.go")
		Ω(err).Should(BeNil())
		Ω(code.Interfaces).ShouldNot(BeNil())
		Ω(code.Interfaces).Should(HaveLen(1))
		x := code.Interfaces[0]
		Ω(x.Name).Should(Equal("Foo"))
		Ω(x.Methods).Should(HaveLen(3))
		m1 := x.Methods[0]
		Ω(m1.Name).Should(Equal("Op1"))
		m2 := x.Methods[1]
		Ω(m2.Name).Should(Equal("Op2"))
		m3 := x.Methods[2]
		Ω(m3.Name).Should(Equal("Op3"))
		Ω(m3.Arguments).Should(HaveLen(0))
		Ω(m3.Result).Should(HaveLen(0))
		Ω(code.Functions).Should(BeNil())
	})
})
