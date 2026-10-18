package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

func main() {
	inputPath := flag.String("input", "docs/swagger.json", "path to the Swagger 2.0 document")
	outputPath := flag.String("output", "docs/openapi.json", "path for the OpenAPI 3 document")
	flag.Parse()

	if err := convert(*inputPath, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "convert Swagger to OpenAPI: %v\n", err)
		os.Exit(1)
	}
}

func convert(inputPath string, outputPath string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", inputPath, err)
	}

	var swagger openapi2.T
	if err := json.Unmarshal(input, &swagger); err != nil {
		return fmt.Errorf("decode %s: %w", inputPath, err)
	}

	openapi, err := openapi2conv.ToV3(&swagger)
	if err != nil {
		return fmt.Errorf("convert document: %w", err)
	}
	if err := openapi.Validate(context.Background()); err != nil {
		return fmt.Errorf("validate converted document: %w", err)
	}

	output, err := json.MarshalIndent(openapi, "", "  ")
	if err != nil {
		return fmt.Errorf("encode OpenAPI document: %w", err)
	}
	output = append(output, '\n')

	if err := os.WriteFile(outputPath, output, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", outputPath, err)
	}

	return nil
}
