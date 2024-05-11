package mrcaller

import "testing"

func TestCallStack_shortenFilePath(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		file   string
		want   string
	}{
		{
			name:   "prefix and file are empty",
			prefix: "",
			file:   "",
			want:   "",
		},
		{
			name:   "prefix is empty",
			prefix: "",
			file:   "bbb",
			want:   "bbb",
		},
		{
			name:   "prefix is not in file",
			prefix: "aaa",
			file:   "bbb",
			want:   "bbb",
		},
		{
			name:   "length of prefix is less than or equal to 3",
			prefix: "aaa",
			file:   "aaabbb",
			want:   "aaabbb",
		},
		{
			name:   "length of prefix is more than 3",
			prefix: "aaa/",
			file:   "aaa/bbb",
			want:   ".../bbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CallStack{
				prefix: tt.prefix,
			}
			if got := c.shortenFilePath(tt.file); got != tt.want {
				t.Errorf("shortenFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
