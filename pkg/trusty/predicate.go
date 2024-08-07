package trusty

import "time"

// TODO(puerco): Protobuf this

type Predicate struct {
	Metadata Metadata       `json:"metadata"`
	Packages []PackageScore `json:"packages"`
}

type Metadata struct {
	Date        *time.Time `json:"date,omitempty"`
	PackageInfo `json:"package,omitempty"`
}

type PackageInfo struct {
	Package     string            `json:"package"`
	Version     string            `json:"version"`
	Identifiers map[string]string `json:"identifiers,omitempty"`
	Ecosystem   string            `json:"ecosystem"`
}

type PackageScore struct {
	PackageInfo
	Score           float64        `json:"score"`
	ActivityScore   float64        `json:"activity"`
	ProvenanceScore float64        `json:"provenance"`
	Details         map[string]any `json:"details"`
	Malicious       bool           `json:"malicious"`
	Deprecated      bool           `json:"deprecated"`
}
