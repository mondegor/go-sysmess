package mrstatus

type (
	// FlowMap - интерфейс для управления статусами их переключениями.
	FlowMap[Status ~uint8] interface {
		Registered() []Status
		Exists(status Status) bool
		Has(from, to Status) bool
		PossibleToStatuses(from Status) []Status
		PossibleFromStatuses(to Status) []Status
	}

	// FlowItem - comment struct.
	FlowItem[Status ~uint8] struct {
		From Status
		To   []Status
	}

	// statusFlow - comment struct.
	statusFlow[Status ~uint8] struct {
		fromToMap     map[Status][]Status
		toFromMap     map[Status][]Status
		registeredMap map[Status]struct{}
		registered    []Status
	}
)

// NewFlowMap - создаёт объект statusFlow.
func NewFlowMap[Status ~uint8](list []FlowItem[Status]) FlowMap[Status] {
	fromToMap := make(map[Status][]Status, len(list))
	toFromMap := make(map[Status][]Status, len(list))
	registeredMap := make(map[Status]struct{}, len(list))

	for _, item := range list {
		fromToMap[item.From] = item.To
		registeredMap[item.From] = struct{}{}

		for _, to := range item.To {
			toFromMap[to] = append(toFromMap[to], item.From)
			registeredMap[to] = struct{}{}
		}
	}

	registered := make([]Status, 0, len(registeredMap))

	for status := range registeredMap {
		registered = append(registered, status)
	}

	return &statusFlow[Status]{
		fromToMap:     fromToMap,
		toFromMap:     toFromMap,
		registeredMap: registeredMap,
		registered:    registered,
	}
}

// Registered - возвращает список зарегистрированных статусов в карте.
func (f *statusFlow[Status]) Registered() []Status {
	return f.registered
}

// Exists - сообщает, зарегистрирован ли данный статус в карте статусов.
func (f *statusFlow[Status]) Exists(status Status) bool {
	_, ok := f.registeredMap[status]

	return ok
}

// Has - сообщает, возможно ли переключить данный статус в указанный статус.
func (f *statusFlow[Status]) Has(from, to Status) bool {
	toStatuses, ok := f.fromToMap[from]
	if !ok {
		return false
	}

	for i := range toStatuses {
		if toStatuses[i] == to {
			return true
		}
	}

	return false
}

// PossibleToStatuses - возвращает список статусов в которые можно переключить указанный статус.
func (f *statusFlow[Status]) PossibleToStatuses(from Status) []Status {
	if toStatuses, ok := f.fromToMap[from]; ok {
		return toStatuses
	}

	return nil
}

// PossibleFromStatuses - возвращает список статусов из которых можно переключиться в указанный статус.
func (f *statusFlow[Status]) PossibleFromStatuses(to Status) []Status {
	if fromStatuses, ok := f.toFromMap[to]; !ok {
		return fromStatuses
	}

	return nil
}
