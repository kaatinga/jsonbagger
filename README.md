# jsonbagger

The `jsonbagger` package provides a simple yet effective solution for extracting JSON objects from strings in Go. It is ideal for dealing with raw data streams or logs where JSON data needs to be isolated from non-JSON content. This package offers two primary functions: `ExtractJSON` and `ExtractJSONIndexes`, both aimed at identifying and extracting JSON objects.

The `jsonbagger` package is especially useful when handling responses from Large Language Models (LLMs) that occasionally include additional comments or text alongside JSON objects. This "noisy" response can cause standard functions like `json.Unmarshal()` to fail due to the presence of non-JSON text. `jsonbagger` efficiently extracts the clean JSON object from such responses, ensuring compatibility with JSON processing functions.

## Installation

To use the `jsonbagger` package in your Go project, simply install it using `go get`:

```bash
go get github.com/kaatinga/jsonbagger
```

## Usage

### ExtractJSON

The `ExtractJSON` function is designed to find and return the first JSON object in a given input string. If no JSON object is found, it returns an error.

#### Syntax

```go
func ExtractJSON(input string) (string, error)
```

#### Parameters

- `input`: The string from which you want to extract the JSON object.

#### Return Value

- Returns the first JSON object as a string.
- If no JSON object is found, an error is returned.

### ExtractJSONIndexes

The `ExtractJSONIndexes` function returns the start and end indexes of the first JSON object found in the input string. This is useful when you need to know the exact position of the JSON object within the larger string.

#### Syntax

```go
func ExtractJSONIndexes(input string) (begin, end int, err error)
```

#### Parameters

- `input`: The string from which you want to extract the JSON indexes.

#### Return Value

- `begin`: The start index of the JSON object.
- `end`: The end index of the JSON object.
- `err`: Returns an error if no JSON object is found.

## Error Handling

The package defines `ErrNotFound` which is returned when no JSON object is found in the provided input string.

## Example

Here is a simple example of how to use the `jsonbagger` package:

```go
package main

import (
	"fmt"
	"github.com/kaatinga/jsonbagger"
)

func main() {
	input := `{"name": "John", "age": 30, "city": "New York"} and some non-JSON text`
	json, err := jsonbagger.ExtractJSON(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Extracted JSON:", json)
}
```

## Contribution

Contributions to the `jsonbagger` package are welcome. Please feel free to submit issues and pull requests through the GitHub repository at `github.com/kaatinga/jsonbagger`.

## License

This package is released under MIT License, which can be found in the [license file](LICENSE.md) in the repository.
