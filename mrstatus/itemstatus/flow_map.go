package itemstatus

import (
	"github.com/mondegor/go-sysmess/mrstatus"
)

// NewFlowMap - возвращает карту возможных переходов ItemStatus.
func NewFlowMap() mrstatus.FlowMap[Enum] {
	return mrstatus.NewFlowMap(
		[]mrstatus.FlowNode[Enum]{
			{
				From: Draft,
				To: []Enum{
					Enabled,
					Disabled,
				},
			},
			{
				From: Enabled,
				To: []Enum{
					Disabled,
				},
			},
			{
				From: Disabled,
				To: []Enum{
					Enabled,
				},
			},
		},
	)
}
