package version

import (
	"bytes"
	"runtime"
	"text/template"
)

var (
	// Version for dingtalk
	Version = "1.1.1"
	// BuildTime for dingtalk
	BuildTime = "2020/05/26"
)

// Options for dingtalk
type Options struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = `Version:      {{.Version}}
Go version:   {{.GoVersion}}
Built:        {{.BuildTime}}
OS/Arch:      {{.Os}}/{{.Arch}}`

// DefaultOps default options
var DefaultOps = Options{
	Version:   Version,
	BuildTime: BuildTime,
	GoVersion: runtime.Version(),
	Os:        runtime.GOOS,
	Arch:      runtime.GOARCH,
}

// GetVersion get version string
func GetVersion() string {
	return GetVersionWithOps(DefaultOps)
}

// GetVersionWithOps get version string with versionOptions
func GetVersionWithOps(options Options) string {
	var doc bytes.Buffer
	template, _ := template.New("version").Parse(versionTemplate)
	template.Execute(&doc, options)
	return doc.String()
}
