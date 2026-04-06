package itemstatus

import (
	"github.com/mondegor/go-sysmess/mrstatus"
)

// NewFlowMap - создаёт карту допустимых переходов для статусов элементов (ItemStatus).
// Правила переходов:
//   - Draft → Enabled, Disabled
//   - Enabled → Disabled
//   - Disabled → Enabled
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
