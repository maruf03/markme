package version

import "runtime"

var (
	AppName   string = "dev"
	Version   string = "unknown"
	Commit    string = "unknown"
	BuildTime string = "unknown"
	GoVersion string = runtime.Version()
)

func Get() VersionInfo {
	return VersionInfo{
		AppName:   AppName,
		Version:   Version,
		Commit:    Commit,
		BuildTime: BuildTime,
		GoVersion: GoVersion,
	}
}
