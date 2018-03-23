package output

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/MovieStoreGuy/go-gitkeys/types"
	"github.com/MovieStoreGuy/resources/marshal"
	yaml "gopkg.in/yaml.v2"
)

const (
	// Using underscores to avoid package name clashing
	_yaml int = iota + 1
	_json
	_raw
)

var (
	formats = map[string]int{
		"yaml": _yaml,
		"json": _json,
		"raw":  _raw,
	}
)

// Printer is a container for io.Writer with formatted output.
type Printer struct {
	output io.Writer
}

// CreatePrinter will create a new printer object with
// the set writer
func CreatePrinter(w io.Writer) (*Printer, error) {
	if w == nil {
		return nil, errors.New("Trying to create printer with a nil writer")
	}
	return &Printer{output: w}, nil
}

// Print takes all the users and print them out desired format
func (p *Printer) Print(format string, users []types.Users) error {
	switch formats[strings.ToLower(format)] {
	case _yaml:
		return p.yamlOutput(users)
	case _json:
		return p.jsonOutput(users)
	case _raw:
		for _, user := range users {
			fmt.Fprintln(p.output, user.Name)
			for _, key := range user.Keys {
				fmt.Fprintln(p.output, key)
			}
			// Adding an extra blank line to make it easier to pass via the command line
			fmt.Fprint(p.output, "\n")
		}
	default:
		return errors.New("Unknown format being defined")
	}
	return nil
}

func (p *Printer) yamlOutput(users []types.Users) error {
	buff, err := yaml.Marshal(&users)
	if err != nil {
		return err
	}
	fmt.Fprint(p.output, "---")
	fmt.Fprintln(p.output, string(buff))
	return nil
}

func (p *Printer) jsonOutput(users []types.Users) error {
	buff, err := marshal.PureMarshal(&users)
	if err != nil {
		return err
	}
	fmt.Fprintln(p.output, string(buff))
	return nil
}
