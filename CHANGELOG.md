# GoSysMess Changelog
Все изменения библиотеки GoSysMess будут документироваться на этой странице.

## 2024-11-16
### Added
- Добавилась возможность добавления следующих опций в конструкторе `NewProto()`:
  - `WithProtoCaller()`;
  - `WithProtoOnCreated()`;
- Добавлена функция `InitDefaultOptions()` для первичной инициализации всех `Proto` ошибок,
  которые были созданы в глобальных переменных и установки обработчика `ProtoOptionsHandler`
  формирования Proto опций по умолчанию.

### Removed
- Удалены следующие функции:
  - `mrerr.WithExtra()`;
  - `mrerrfactory.NewProtoAppErrorByDefault()`;
  - `mrerrfactory.NewProtoAppError()`;

## 2024-10-27
### Added
- Добавлены значения по умолчанию для опций `mrcaller.Caller`;
- Добавлена вспомогательная функция `mrmsg.ToString()` для перевода значения переменной в строку;
- Добавлен метод `mrmsg.Data.ValueString()`;
- Добавлена другая функция `mrmsg.Render` с параметром для переменных `map[string]string`;
- Добавлена функция `mrmsg.RenderWithData` с параметром для переменных `mrmsg.Data`;

### Changed
- Переименованы:
    - `mrmsg.Render` -> `mrmsg.RenderWithNamedArgs`;
    - `mrmsg.MustRender` -> `mrmsg.MustRenderWithNamedArgs`;
    - `mrmsg.CheckParse` -> `mrmsg.CheckRender`;
- Метод `mrmsg.NamedArg.ValueString()` теперь работает на базе `mrmsg.ToString()`;

## 2024-10-11
### Added
- Подключён и настроен линтер `gci`;

### Changed
- Исправлены ссылки в документации;

## 2024-10-06
### Changed
- Переименованы:
    - `ErrorCodeUnexpectedInternal` -> `ErrorCodeInternal`;
    - `ErrorCodeUnexpectedSystem` -> `ErrorCodeSystem`;
- Включена проверка наличия глобальных переменных в тестах; 

## 2024-09-29
### Changed
- Добавлена директория `.cache` в `.gitignore`, скорректирована документация;

## 2024-09-27
### Changed
- Поправлен `.editorconfig`, добавлены `*.proto`, `*.mk`;
- Удалена нулевая ёмкость у map;

## 2024-09-15
### Added
- Добавлены короткие команды в Makefile, обновлена инструкция команд.

## 2024-09-14
### Changed
- Переименованы CallerOption:
    - `DepthOption` -> `WithDepth`;
    - `ShowFuncNameOption` -> `WithShowFuncName`;
    - `FilterStackTraceOption` -> `WithFilterStackTrace`;

## 2024-09-08
### Changed
- Поправлены `.env` переменные под новую версию `mrcmd`;

## 2024-08-03
### Changed
- Добавлен более информативный вывод в `mrmsg.NamedArg.ValueString()`;

## 2024-07-06
### Changed
- Доработана документация по архитектурному решению обработки ошибок;
- Обновлены `github.workflows` и версия `go` до `1.22`;
- В тестах `defer` заменён на `t.Cleanup`;

## 2024-06-30
### Added
- Добавлена документация по архитектуре обработки ошибок и ссылки на пример реализации;

## 2024-06-22
### Added
- Добавлена опция `ShowFuncNameOption` для управления выводом имени функции в `CallStack`;
- Добавлена функция `WithoutStackTrace` для возможности вывода ошибки без стека вызовов если он был сформирован;
- Добавлена структура `ProtoExtra` для дополнительных настроек при создании ошибок;

### Changed
- Функция `FilterStackTraceTrimTop` переименована в `FilterStackTraceTrimUpper`
  и добавлено отсечение нижней границы;

## 2024-06-16
### Changed
- Настроен линтер `reviev` (`.golangci.yaml`);

## 2024-06-15
### Added
- Добавлено несколько новых линтеров;
- Добавлена `NewProtoAppErrorByDefault` и обновлена `NewProtoAppError`;
- Подключены линтеры с их настройками (`.golangci.yaml`);

### Changed
- Доработан `CustomError`, добавлен к нему метод `IsValid()`, покрыт тестами;
- Обновлены примеры использования `ProtoAppError` и `CustomError`;

## 2024-06-09
### Changed
- Переработан пакет ошибок:
    - объект `AppErrorFactory` заменён на `ProtoAppError` который теперь сам является ошибкой;
    - добавление генераторов ID и стека к прототипу ошибок делается через `mrerr.WithExtra`,
      это позволило отказаться от глобальных переменных;
- Написаны тесты для пакета `mrerr`;
- Изменился интерфейс `CustomError`;
- Добавлены примеры работы с пакетом `mrerr`;
- Добавлена константа `GO_DEV_LOCAL_PACKAGE` и поправлены `imports` при помощи `goimports`;

## 2024-06-02
### Changed
- Формирование `CallStack` переработано и вынесено в отдельный пакет,
  с помощью `mrerr.GlobalCallerFunc` этот пакет состыковывается с ошибкой;
- Генерация ID ошибки вынесена в отдельный объект с возможностью его
  переопределения (см. `mrerr.GlobalIDGenerator`);
- Изменён формат создания новых ошибок;
- К проекту подключены линтеры с их настройками (`.golangci.yaml`);
- Добавлены комментарии для публичных объектов и методов;

## 2024-03-19
### Changed
- Поправлено форматирование документации;

## 2024-03-18
### Changed
- Переработан механизм формирования `CallStack`:
    - теперь он не зависит от типа ошибки, а включается с помощью конструктора `mrerr.NewFactoryWithCaller`;
    - переименован метод `AppErrorFactory.Caller -> WithCaller`, который теперь принудительно включает
      формирование `CallStack`, а также добавлен метод `DisableCaller()` для принудительного его отключения;
    - удалена константа `ErrorKindInternalNotice`, вместо неё достаточно использовать `ErrorKindInternal` вместе
      с `mrerr.NewFactory`, который не формирует по умолчанию `CallStack`;
    - для `AppError` добавлен метод `HasCallStack()`, который возвращает, был сформирован для самой ошибки или
      для одной из её вложенных ошибок `CallStack`;

## 2024-03-14
### Changed
- В функции `mrlang.NewTranslator` переименован параметр `opt -> opts`;

## 2024-02-05
### Changed
- Добавлены в ошибку `error parsing dictionary file` подробности о причинах её возникновения;

## 2024-01-30
### Added
- Добавлены: `mrlang.WithContext()`, `mrlang.Ctx()`, `mrlang.Locale.WithContext()`;

### Changed
- Переименован `mrmsg.ErrorTranslator.CheckError() -> HasErrorCode()`;

### Removed
- Удалён `mrlang.DefaultLocale()` (необходимо использовать `mrlang.Translator.DefaultLocale()`);

## 2024-01-22
### Added
- Добавлен метод `AppErrorFactory.WithAttr` для прикрепления доп. информации к ошибке;

### Changed
- Переименовано:
    - `FieldError -> CustomError`;
    - `FieldErrorList -> CustomErrorList`;
    - `AppError.ID() -> Code()`;
    - `AppErrorFactory.ErrorID() -> Code()`;
    - `Locale.CheckErrorID -> CheckError`;
    - `NamedArg.valueString -> ValueString`;
- Добавлен интерфейс `mrmsg.ErrorTranslator` чтобы избавиться от зависимости пакета `mrlang`;

## 2024-01-19
### Changed
- Заменён тип ID языка с `int32` на `uint16`;

## 2024-01-16
### Added
- Добавлена поддержка регионов для языка;
- Добавлен целочисленный идентификатор для языка `langID`;

### Changed
- Доработана ParseAcceptLanguage для полноценной поддержки регионов;
- Удалена зависимость от `cleanenv.ReadConfig`;
- Добавлена поддержка словарей для локализации наборов данных (перечисления, таблицы БД);

## 2023-12-09
### Changed
- Перенесён `mrerr.Arg -> mrmsg.Data`;
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
- Переименован метод `FieldErrorList.AddAppErr -> AddAppError`;
- ErrorKind вынесен в отдельный файл и добавлен метод String() к нему;

### Removed
- Удалён метод `FieldError.Kind()`;

## 2023-11-19
### Changed
- Переработан механизм работы с пользовательскими ошибками, которые привязываются к конкретным полям объектов:
    - для `FieldError` добавлены методы: `NewFieldError`, `NewFieldErrorAppErr`, `NewFieldMessage`;
    - у `FieldErrorList` удалены методы `NewList` и `NewListWith` (теперь необходимо пользоваться методами у `FieldError`), добавлен метод `AddAppErr`;
    - обновлён пример работы с такими ошибками;
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
- `TranslatorOptions.LangByDefault -> DefaultLang`;

## 2023-09-11
### Fixed
- Формат глобальных `const`, `type`, `var` приведён к общему виду;

### Removed
- Удалены из пакета переменные предопределёнными ошибками;

## 2023-09-10
### Changed
- Обновлены зависимости библиотеки;
- `FactoryDataContainer -> ErrFactoryInternalNoticeDataContainer`;
- `FactoryInternal* -> ErrFactoryInternal*`;

### Fixed
- Исправлен баг в `examples/field/main.go`;

## 2023-09-03
### Changed
- `ErrorId -> string`;
- `FieldErrorList -> *FieldErrorList`;