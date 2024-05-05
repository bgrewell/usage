package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetExecutableName(t *testing.T) {
	expected := filepath.Base(os.Args[0])
	expected = expected[:len(expected)-len(filepath.Ext(expected))]
	got := GetExecutableName()
	if got != expected {
		t.Errorf("GetExecutableName() = %s; want %s", got, expected)
	}
}

func TestWrapText(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		width     int
		prefixLen int
		want      []string
	}{
		{
			name:      "Short case",
			text:      "This is a short case",
			width:     60,
			prefixLen: 0,
			want:      []string{"This is a short case"},
		},
		{
			name:      "Empty text",
			text:      "",
			width:     10,
			prefixLen: 0,
			want:      []string{""},
		},
		{
			name:      "Long text without prefix",
			text:      "Length longer than the column width. In this case, the column width is 20",
			width:     20,
			prefixLen: 0,
			want:      []string{"Length longer than", "the column width. In", "this case, the", "column width is 20"},
		},
		{
			name:      "Long text with a prefix",
			text:      "It's almost a browser but not quite. Instead it's just a example of how to use the 'usage' package. It is designed to show an example of how to use the package and a sample of what the output would look like.",
			width:     60,
			prefixLen: 13,
			want: []string{
				"It's almost a browser but not quite. Instead",
				"it's just a example of how to use the 'usage' package. It is",
				"designed to show an example of how to use the package and a",
				"sample of what the output would look like.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wrapText(tt.text, tt.width, tt.prefixLen)
			if len(got) != len(tt.want) {
				t.Errorf("wrapText() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("wrapText() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
