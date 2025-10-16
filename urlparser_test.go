package urlparser_test

import (
	"testing"

	"github.com/AlvinNathanael/urlparser"
)

func TestRemoveURLQueries(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		inputURL  string
		queryKeys []string
		want      string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := urlparser.RemoveURLQueries(tt.inputURL, tt.queryKeys...)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("RemoveURLQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
