# GoSysMess Changelog
Все изменения библиотеки GoSysMess будут документироваться на этой странице.

## 2023-12-09
### Changed
- Перенесён `mrerr.Arg` -> `mrmsg.Data`;
- Доработана логика копирования объекта в `AppErrorFactory.Caller`;
- `CallStack` заменён на `Caller`, который теперь сам формирует `CallStack` и может использоваться независимо;

### Removed
- Удалён mrmsg.NewArg;

## 2023-12-06
### Added
- Добавлен метод `AppErrorFactory.ErrorID()`, который возвращает идентификатор типа ошибки;
- Добавлено управление `CallStack` для `Internal` и `System` ошибок, с помощью метода `SetCallStackOptions` можно указывать глубину отображаемого стека;

## 2023-12-04
### Added
- Добавлен метод `Locale.CheckErrorID()` для проверки есть ли перевод для указанного идентификатора типа ошибки;

### Changed
- Теперь описания всех Internal и System ошибок можно переопределять в yaml файле для вывода их пользователю;

### Removed
- Удалены методы:
    - `FieldErrorList.Add()`;
    - `FieldErrorList.AddAppError()`;

## 2023-12-03
### Changed
- Генерация ID ошибки реализована на стандартных библиотеках и вынесена в отдельный метод `generateTraceID`;
- Метод `NewFieldMessage` переименован в `NewFieldErrorMessage` и в нём изменилась логика формирования id ошибки; 
- Переименован метод `FieldErrorList.AddAppErr` -> `AddAppError`;
- ErrorKind вынесен в отдельный файл и добавлен метод String() к нему;

### Removed
- Удалён метод `FieldError.Kind()`;

## 2023-11-19
### Changed
- Переработан механизм работы с пользовательскими ошибками, которые привязываются к конкретным полям объектов:
    - Для `FieldError` добавлены методы: `NewFieldError`, `NewFieldErrorAppErr`, `NewFieldMessage`;
    - У `FieldErrorList` удалены методы `NewList` и `NewListWith` (теперь необходимо пользоваться методами у `FieldError`), добавлен метод `AddAppErr`;
    - Обновлён пример работы с такими ошибками;
- В некоторых местах оптимизирована конкатенация строк (`Sprintf` заменён на нативный "+");
- Обновлён `.editorconfig`;

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