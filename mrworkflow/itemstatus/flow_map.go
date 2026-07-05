package itemstatus

import (
	"github.com/mondegor/go-core/mrworkflow"
)

// NewFlowMap - создаёт карту допустимых переходов для статусов элементов (ItemStatus).
// Правила переходов:
//   - Draft → Enabled, Disabled
//   - Enabled → Disabled
//   - Disabled → Enabled
func NewFlowMap() mrworkflow.FlowMap[Enum] {
	return mrworkflow.NewFlowMap(
		[]mrworkflow.FlowNode[Enum]{
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
