package code_analyzer

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodeAnalyzer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeAnalyzer Suite")
}
