package mrstatus

type (
	// FlowMap - интерфейс управления переходами между статусами.
	// Определяет, какие статусы зарегистрированы и какие переходы между ними допустимы.
	FlowMap[Status ~uint8] interface {
		Registered() []Status
		Exists(status Status) bool
		IsPossible(from, to Status) bool
		PossibleToStatuses(from Status) []Status
		PossibleFromStatuses(to Status) []Status
	}

	// FlowNode - описывает допустимые переходы из одного статуса в другие.
	// From - исходный статус.
	// To - список статусов, в которые разрешён переход из From.
	FlowNode[Status ~uint8] struct {
		From Status
		To   []Status
	}

	statusFlow[Status ~uint8] struct {
		fromToMap     map[Status][]Status
		toFromMap     map[Status][]Status
		registeredMap map[Status]bool
		registered    []Status
	}
)

// NewFlowMap - создаёт карту допустимых переходов между статусами.
// Параметр list - список узлов переходов (FlowNode), определяющих граф состояний.
// Автоматически строит двунаправленную карту: from→to и to→from.
func NewFlowMap[Status ~uint8](list []FlowNode[Status]) FlowMap[Status] {
	fromToMap := make(map[Status][]Status, len(list))
	toFromMap := make(map[Status][]Status, len(list))
	registeredMap := make(map[Status]bool, len(list))

	for _, item := range list {
		fromToMap[item.From] = item.To
		registeredMap[item.From] = true

		for _, to := range item.To {
			toFromMap[to] = append(toFromMap[to], item.From)
			registeredMap[to] = true
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

// Exists - сообщает, имеется ли данный статус в карте статусов.
func (f *statusFlow[Status]) Exists(status Status) bool {
	return f.registeredMap[status]
}

// IsPossible - сообщает, возможно ли переключить данный статус в указанный статус.
func (f *statusFlow[Status]) IsPossible(from, to Status) bool {
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
