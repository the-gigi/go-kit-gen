package code_generator

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Code generator tests", func() {
	var (
		err          error
		templatesDir string
		g            *CodeGenerator
	)

	BeforeEach(func() {
		templatesDir, err = filepath.Abs("templates")
		Ω(err).Should(BeNil())
		g, err = NewCodeGenerator(templatesDir)
		Ω(err).Should(BeNil())
		Ω(g).ShouldNot(BeNil())
	})

	Context("Creation tests", func() {
		It("Should create a code generator with a valid templates dir", func() {
			g, err = NewCodeGenerator(templatesDir)
			Ω(err).Should(BeNil())
			Ω(g).ShouldNot(BeNil())
		})

		It("Should fail to create a code generator with a non-existing templates dir", func() {
			g, err = NewCodeGenerator("no-such-dir")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("GRPC transport generation tests", func() {
		It("Should generate a GRPC transport with valid input", func() {
		})

		It("Should fail to generate a GRPC transport with invalid input", func() {
		})
	})

	Context("Service generation tests", func() {
		It("Should generate a service with valid input", func() {
		})

		It("Should fail to generate a service invalid input", func() {
		})
	})

	Context("Client generation tests", func() {
		It("Should generate a client with valid input", func() {
		})

		It("Should fail to generate a client invalid input", func() {
		})
	})

	Context("Serialization generation tests", func() {
		It("Should generate a serialization with valid input", func() {
		})

		It("Should fail to generate a serialization invalid input", func() {
		})
	})
})
