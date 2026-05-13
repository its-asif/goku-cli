/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	gokuInputPath  string
	gokuOutputType string
	gokuOutputPath string
)

var converterCmd = &cobra.Command{
	Use:   "converter",
	Short: "Use to convert json to yaml and viceversa",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGokuConversion()
	},
}

func init() {
	converterCmd.Flags().StringVarP(&gokuInputPath, "input", "i", "", "input file path")
	converterCmd.Flags().StringVarP(&gokuOutputType, "output", "o", "", "input file type (Default: json)")
	converterCmd.Flags().StringVarP(&gokuOutputPath, "outputPath", "p", "", "(optional)output file path (Default: */output)")

	_ = converterCmd.MarkFlagRequired("input")
	_ = converterCmd.MarkFlagRequired("output")
}

func runGokuConversion() error {

	// read file
	inputbyte, err := os.ReadFile(gokuInputPath)
	if err != nil {
		return fmt.Errorf("read input file %w", err)
	}

	// take out input and outpur ext

	inputExt := strings.ToLower(strings.TrimPrefix(filepath.Ext(gokuInputPath), "."))
	outputExt := strings.ToLower(gokuOutputType)

	if outputExt != "json" && outputExt != "yaml" {
		return fmt.Errorf("unsupported output format :%q... use json or yml", outputExt)
	}

	if inputExt == outputExt {
		return fmt.Errorf("Input extension is similar to required extension")
	}

	// unmarshal according to inputtype
	var data any
	switch inputExt {
	case "json":
		if err := json.Unmarshal(inputbyte, &data); err != nil {
			return fmt.Errorf("parse json input: %w", err)
		}
	case "yaml":
		if err := yaml.Unmarshal(inputbyte, &data); err != nil {
			return fmt.Errorf("parse yml input: %w", err)
		}
	default:
		return fmt.Errorf("unsupported input extension %q", inputExt)
	}

	// convert to output (Marshall em)
	var outputBytes []byte
	switch outputExt {
	case "json":
		outputBytes, err = json.MarshalIndent(data, "", " ")
		if err != nil {
			return fmt.Errorf("encode json output error: %w", err)
		}
	case "yaml":
		outputBytes, err = yaml.Marshal(data)
		if err != nil {
			return fmt.Errorf("encode yml output error: %w", err)
		}
	default:
		return fmt.Errorf("unsupported output extension: %q", outputExt)
	}

	// find output directory... create if not present
	outputDir := filepath.Join(filepath.Dir(filepath.Dir(gokuInputPath)), "output")
	if gokuOutputPath != "" {
		outputDir = gokuOutputPath
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("creating folder error: %w", err)
	}

	// file name & write output file in output path
	outputPath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(gokuInputPath), inputExt)+outputExt)
	if err := os.WriteFile(outputPath, outputBytes, 0o644); err != nil {
		return fmt.Errorf("write output file %w", err)
	}

	fmt.Printf("Wrote %s\n", outputPath)

	return nil
}
