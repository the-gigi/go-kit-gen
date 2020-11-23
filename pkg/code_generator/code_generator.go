package code_generator

import (
	"os"

	"github.com/the-gigi/go-kit-gen/pkg/code_analyzer"
)

type CodeGenerator struct {
	templatesDir string
}

func (g *CodeGenerator) GenerateGRPCTransport(in code_analyzer.Interface) {
	// Generate header

	// Generate request and response structs

	// Generate decoders

	// Generate response encoder

	// Generate endpoints

}

func (g *CodeGenerator) GenerateService(in code_analyzer.Interface) {

}

func (g *CodeGenerator) GenerateClient(in code_analyzer.Interface) {

}

func (g *CodeGenerator) GenerateSerialization() {

}

func NewCodeGenerator(templatesDir string) (g *CodeGenerator, err error) {
	_, err = os.Stat(templatesDir)
	if os.IsNotExist(err) {
		return
	}

	g = &CodeGenerator{
		templatesDir: templatesDir,
	}
	return
}
