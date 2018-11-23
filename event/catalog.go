package event

const (
	VmExecutionRequestEvent     = "VmExecutionRequest"
	HowlVmExecutionCommandEvent = "HowlVmExecutionCommand"
)

type VmExecutionRequest struct {
	ComputationId string `json:"computationId"`
	Type          string `json:"type"`
	WantedNodes   int    `json:"wantedNodes"`
	Code          string `json:"code"`
}

type HowlVmExecutionCommand struct {
	ComputationId string   `json:"computationId"`
	Type          string   `json:"type"`
	Nodes         []string `json:"nodes"`
	Code          string   `json:"code"`
}
