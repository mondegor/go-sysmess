# GoSysMess Changelog
Все изменения библиотеки GoSysMess будут документироваться на этой странице.

## 2023-11-12
### Changed
- Изменён `callerSkip` с 3 на 4, для того чтобы в логах выводить путь к родительской функции, где произошла ошибка;
- Переименованы некоторые переменные и функции (типа Id -> ID) в соответствии с code style языка go;
- Все файлы библиотеки были пропущены через `gofmt`;

## 2023-11-01
### Changed
- Оптимизирована работа с некоторыми структурами данных;
- Обновлены зависимости библиотеки;

## 2023-10-08
### Changed
- В методе `mrlang.newLocale` обработка ошибок приведена к более компактному виду;

## 2023-09-16
### Changed
- Заменено `*string` на `string` при формировании `traceId` и пути файла `file`;
- Сообщение об ошибке теперь формируется с помощью `strings.Builder`;  

## 2023-09-13
### Changed
- Заменено `Message -> string`, чтобы избежать лишних преобразований;

### Fixed
- `mrlang.Locale -> *mrlang.Locale`;

## 2023-09-12
### Added
- Добавлено описание о принципах обработки ошибок;

### Changed
- Изменена логика определения языка по умолчанию;
- `TranslatorOptions.LangByDefault` -> `DefaultLang`;

## 2023-09-11
### Fixed
- Формат глобальных `const`, `type`, `var` приведён к общему виду;

### Removed
- Удалены из пакета переменные предопределёнными ошибками;

## 2023-09-10
### Changed
- Обновлены зависимости библиотеки;
- `FactoryDataContainer` -> `ErrFactoryInternalNoticeDataContainer`;
- `FactoryInternal*` -> `ErrFactoryInternal*`;

### Fixed
- Исправлен баг в `examples/field/main.go`;

## 2023-09-03
### Changed
- `ErrorId` -> `string`;
- `FieldErrorList` -> `*FieldErrorList`;