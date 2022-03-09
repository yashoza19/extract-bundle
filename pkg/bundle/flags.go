package bundle

// ReportFlags define the flags used to generate the bundle report
type Flags struct {
	IndexImage      string `json:"image"`
	Filter          string `json:"filter"`
	ContainerEngine string `json:"containerEngine"`
}
