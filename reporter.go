package reporter

import (
	"fmt"
	"github.com/Hranoprovod/shared"
	"io"
)

// Reporter is the main report structure
type Reporter struct {
	output io.Writer
}

// NewReporter creates new reporter
func NewReporter(writer io.Writer) *Reporter {
	return &Reporter{
		writer,
	}
}

// PrintAPISearchResult prints a list of search resilts
func (r *Reporter) PrintAPISearchResult(nl shared.APINodeList) error {
	for _, n := range nl {
		err := r.PrintAPINode(n)
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintAPINode prints single API result
func (r *Reporter) PrintAPINode(n shared.APINode) error {
	_, err := fmt.Fprintln(r.output, n.Name)
	return err
}
