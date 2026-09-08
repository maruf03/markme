package api

type StatusEnum string

const (
	Pass StatusEnum = "pass"
	Fail StatusEnum = "fail"
	Warn StatusEnum = "warn"
)

type LivezResponse struct {
	Status StatusEnum `json:"status"`
}

type ReadyzResponse struct {
	Status StatusEnum               `json:"status"`
	Checks map[string]LivezResponse `json:"checks"`
}

type HealthzResponse struct {
	Status  StatusEnum               `json:"status"`
	Version string                   `json:"version"`
	Checks  map[string]LivezResponse `json:"checks"`
}

type VersionResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
}
