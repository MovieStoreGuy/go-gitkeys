package output

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/MovieStoreGuy/keyobtainer/types"
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
	Formats = map[string]int{
		"yaml": _yaml,
		"json": _json,
		"raw":  _raw,
	}
)

type Printer struct {
	output io.Writer
}

func CreatePrinter(w io.Writer) (*Printer, error) {
	if w == nil {
		return nil, errors.New("Trying to create printer with a nil writer")
	}
	return &Printer{output: w}, nil
}

func (p *Printer) Print(format string, users []types.Users) error {
	switch Formats[strings.ToLower(format)] {
	case _yaml:
		buff, err := yaml.Marshal(&users)
		if err != nil {
			return err
		}
		fmt.Fprintln(p.output, string(buff))
	case _json:
		buff, err := marshal.PureMarshal(&users)
		if err != nil {
			return err
		}
		fmt.Fprint(p.output, "---")
		fmt.Fprintln(p.output, string(buff))
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
