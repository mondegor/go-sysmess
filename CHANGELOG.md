# GoSysMess Changelog
Все изменения библиотеки GoSysMess будут документироваться на этой странице.

## 2023-09-11
### Fixed
- Формат глобальных const, type, var приведён к общему виду;

### Removed
- Удалены из пакета переменные предопределёнными ошибками;

## 2023-09-10
### Changed
- Обновлены зависимости библиотеки;
- FactoryDataContainer -> ErrFactoryInternalNoticeDataContainer;
- FactoryInternal* -> ErrFactoryInternal*;

### Fixed
- Исправлен баг в examples/field/main.go;

## 2023-09-03
### Changed
- ErrorId -> string;
- FieldErrorList -> *FieldErrorList;