package internal

import (
	"os"
	"testing"
)

func TestGetContentTypeAndExtension(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantCt   string
		wantExt  string
	}{
		{
			name:     "image/jpeg",
			filename: "./helper_test.jpg",
			wantCt:   "image/jpeg",
			wantExt:  ".jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.filename)
			if err != nil {
				t.Errorf("open file (%s), err: %+v", tt.filename, err)
			}
			defer f.Close()

			data, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Errorf("read file (%s), err: %+v", tt.filename, err)
			}

			gotCt, gotExt := GetContentTypeAndExtension(data)
			if gotCt != tt.wantCt {
				t.Errorf("GetContentTypeAndExtension() gotCt = %v, want %v", gotCt, tt.wantCt)
			}
			if gotExt != tt.wantExt {
				t.Errorf("GetContentTypeAndExtension() gotExt = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}
