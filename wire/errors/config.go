package errors

type (
	// ErrorConfig - конфигурация для инициализации runtime-ошибок.
	// Определяет, использовать ли стек вызовов и с какой глубиной.
	ErrorConfig struct {
		// HasCaller - включает сбор информации (файл/строка/функция).
		HasCaller bool

		// CallerDepth - глубина стека вызовов (количество кадров).
		CallerDepth uint8

		// CallerShowFunc - включает отображение имени функции.
		CallerShowFunc bool

		// CallerUpperBounds - список имён пакетов/функций для ограничения стека вызова.
		CallerUpperBounds []string
	}
)
