package client

const (
	PipelineConditionType                    = "pipelineCondition"
	PipelineConditionFieldLastTransitionTime = "lastTransitionTime"
	PipelineConditionFieldLastUpdateTime     = "lastUpdateTime"
	PipelineConditionFieldMessage            = "message"
	PipelineConditionFieldReason             = "reason"
	PipelineConditionFieldStatus             = "status"
	PipelineConditionFieldType               = "type"
)

type PipelineCondition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}
