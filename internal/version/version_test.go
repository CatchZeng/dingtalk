package version

import (
	"testing"
)

func TestGetVersionWithOps(t *testing.T) {
	tests := []struct {
		name string
		ops  Options
		want string
	}{
		{name: "test1",
			ops: Options{
				Version:   "1.0.0",
				BuildTime: "20190909",
				GoVersion: "1.0",
				Os:        "Win",
				Arch:      "X86",
			},
			want: `Version:      1.0.0
Go version:   1.0
Built:        20190909
OS/Arch:      Win/X86`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVersionWithOps(tt.ops); got != tt.want {
				t.Errorf("\n%v\n%v", got, tt.want)
			}
		})
	}
}

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get version with default option",
			want: GetVersionWithOps(DefaultOps),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
