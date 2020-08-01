package version

import (
	"bytes"
	"runtime"
	"text/template"
)

var (
	// Version for dingtalk
	Version = "2.0.0"
	// BuildTime for dingtalk
	BuildTime = "2020/08/01"
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
	tpl, _ := template.New("version").Parse(versionTemplate)
	_ = tpl.Execute(&doc, options)
	return doc.String()
}
