# goku-cli

`goku-cli` is a small Go command-line tool to convert files between JSON and YAML.

## Features

- Convert `.json` to `.yaml`
- Convert `.yaml` to `.json`
- Write output to a default output directory or a custom path

## Prerequisites

- Go `1.26.1` or newer

## Install dependencies

```bash
go mod tidy
```

## Run the CLI

```bash
go run . --help
```

## Command usage

```bash
goku converter -i <input-file> -o <output-format> [-p <output-path>]
```

### Flags

- `-i, --input` (required): input file path
- `-o, --output` (required): output format (`json` or `yaml`)
- `-p, --outputPath` (optional): output directory path

## Examples

Convert JSON to YAML:

```bash
go run . converter -i assets/input/temp.json -o yaml
```

Output:

```text
Wrote assets/output/temp.yaml
```

Convert YAML to JSON:

```bash
go run . converter -i assets/input/temp.yaml -o json
```

Write converted file to a custom directory:

```bash
go run . converter -i assets/input/temp.json -o yaml -p ./my-output
```

## Notes

- Supported input file extensions: `json`, `yaml`
- Supported output formats: `json`, `yaml`
- If `-p` is not provided, output is written to an `output` folder next to your `input` folder.
