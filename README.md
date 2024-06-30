# Описание GoSysMess v0.10.7
Этот репозиторий содержит описание библиотеки GoSysMess.

## Статус библиотеки
Библиотека находится в стадии разработки.

## Описание библиотеки
Библиотека решает две основные задачи:
- формирует пользовательские сообщения на различных языках, в том числе и сообщения об ошибках;
- даёт инструменты для более удобной обработки ошибок как пользовательских, так и программных
  (более подробно см. ниже);

## Подключение библиотеки
`go get -u github.com/mondegor/go-sysmess@v0.10.7`

## Установка библиотеки для её локальной разработки

- Выбрать рабочую директорию, где должна быть расположена библиотека
- `mkdir go-sysmess && cd go-sysmess` // создать и перейти в директорию проекта
- `git clone git@github.com:mondegor/go-sysmess.git .`
- `cp .env.dist .env`

### Консольные команды используемые при разработке библиотеки

> Перед запуском консольных скриптов сервиса необходимо скачать и установить утилиту Mrcmd.\
> Инструкция по её установке находится [здесь](https://github.com/mondegor/mrcmd#readme)

- `mrcmd go-dev fmt` // исправляет форматирование кода (gofumpt -l -w -extra ./)
- `mrcmd go-dev goimports-fix` // исправление imports, если это требуется (goimports -d -local ${GO_DEV_LOCAL_PACKAGE} ./)
- `mrcmd go-dev check` // статический анализ кода библиотеки
- `mrcmd go-dev test` // запуск тестов библиотеки
- `mrcmd go-dev help` // выводит список всех доступных команд

## Обработка ошибок. Общие сведения
### Выделяются три вида ошибок:
- Пользовательские ошибки (user) - все те ошибки, которые должен исправить сам пользователь
  (примеры: недопустимый ввод данных; обращение к несуществующему ресурсу);
- Системные (system) - происходящие в системе и независящие от программы (примеры: отсутствует
  соединение с БД или с каким либо внешнем API; отсутствие прав на запись файлов и т.д.);
- Программные (internal) - любые другие ошибки, допущенные разработчиками, или ошибки, которые
  не были классифицированы предыдущими видами (примеры: обращение к нулевому указателю,
  создание файла в несуществующей папке; выход из диапазона значений массива);

### Действия при обработке пользовательской ошибки:
- понятно объяснить причину возникновения;
- предложить варианты действий;
- возможно, превратить ошибку в пользовательский сценарий (например, если логин уже существует,
  то сгенерить похожие варианты и предложить пользователю выбрать один из них);

### Действия при обработке системной ошибки:
- понятно объяснить причину возникновения;
- предложить варианты действий, если возможно;
- сообщить, с чем обратиться в поддержку;
- отобразить уникальный код ошибки, под которым система предварительно записала эту ошибку в логи;
- узнавать об ошибках автоматически, а не от пользователя (т.к. не все пользователи обращаются
  в поддержку, да и оперативное исправление ошибок увеличивает лояльность пользователей);

### Действия при обработке программной ошибки:
- в заголовке всегда пишется фраза типа "Internal error";
- сообщить, с чем обратиться в поддержку;
- показать подробности, если это безопасно;
- отобразить уникальный код ошибки, под которым система предварительно записала эту ошибку в логи;
- узнавать об ошибках автоматически, а не от пользователя;

### Важные замечания:
- К какому виду отнести ту или иную ошибку могут только разработчики, помочь им в этом могут QA специалисты;
- Обработка ошибок - это тоже бизнес логика, поэтому её также нужно продумывать;
- Ошибка является интерфейсом вызова функции, поэтому её необходимо обрабатывать;

### Сценарии работы с ошибками:
- user - понимает что случилось и как это исправить;
- developer - исследует возникшую проблему и выявляет её причины;
- team lead - отслеживает всё ли работает нормально; выявляет новые ошибки,
  принесённые релизом; оценивает влияние ошибок на пользователей;

## Архитектура обработки ошибок
Внутренний архитектурный слой не должен влиять на поведение верхнего, поэтому ошибки должны быть
частью интерфейса каждого слоя. Определение типов ошибок должны быть в том же слое где возникает
данная ошибка. Причём для каждой ошибки должен быть указан её вид: user, system, internal.

Основная обработка (перехват) ошибок должна происходить в слое UseCase, только он знает как
именно обрабатывать ошибки поступивших из других слоёв, а также правильно определить нужный
вид ошибки. Конкретный слой, обрабатывая ошибку, должен создать объект с ошибкой, определённую
в этом слое, а если этот слой обрабатывает уже перехваченную ошибку, то он должен вложить
её в созданную им ошибку, для того, чтобы внешний перехватчик смог обработать всю цепочку
ошибок и выполнить необходимые действия для каждого вида ошибок.

Чтобы избавить код от текстов пользовательских ошибок, а также одновременно решить проблему
локализации, все тексты пользовательских ошибок необходимо хранить отдельно и загружать
их по заранее выданным им идентификаторам.

## Пример архитектуры обработки ошибок

### Подсистема формирования ошибок
- [ProtoAppError](https://github.com/mondegor/go-sysmess/blob/master/mrerr/error_proto.go)
- [AppError](https://github.com/mondegor/go-sysmess/blob/master/mrerr/error.go)
- [CustomError](https://github.com/mondegor/go-sysmess/blob/master/mrerr/custom_error.go)
- [IDGenerator](https://github.com/mondegor/go-sysmess/blob/master/mrerr/features/generate_instance_id.go)
- [StackTraceCaller](https://github.com/mondegor/go-sysmess/blob/master/mrcaller/caller.go)

![image](docs/resources/packages/c4/sysmess.svg)

### Подсистема обработки ошибок
- [ErrorManager](https://github.com/mondegor/go-webcore/blob/master/mrcore/mrinit/error_manager.go)
- [Базовый ErrorHandler](https://github.com/mondegor/go-webcore/blob/master/mrcore/mrcoreerr/error_handler.go)
- [Создание системных типов ошибок](https://github.com/mondegor/go-webcore/blob/master/mrcore/usecase_errors.go)
- [Хелпер обёртки ошибок](https://github.com/mondegor/go-webcore/blob/master/mrcore/mrcoreerr/error_wrapper.go)
- [Пример первичной обработки ошибки Http сервером](https://github.com/mondegor/go-webcore/blob/18fb17d935f4f6b1640fe68ed687b9248457a42f/mrserver/mrresp/sender_error.go#L51)

![image](docs/resources/packages/c4/errcore.svg)

### Система использующая обработку ошибок
- [Создание и настройка ErrorManager](https://github.com/mondegor/go-sample/blob/master/app/cmd/factory/error_manager.go)
- [Создание типов ошибок модуля](https://github.com/mondegor/go-sample/blob/master/app/internal/catalog/product/module/errors.go)
- [Регистрация типов ошибок модулей](https://github.com/mondegor/go-sample/blob/34b638018314dcc99f1f8e93c8172fc49a9c8c9d/app/cmd/factory/app_environment.go#L122)
- [Создание ошибок, прямая обёртка ошибок и с использованием хелпера](https://github.com/mondegor/go-sample/blob/34b638018314dcc99f1f8e93c8172fc49a9c8c9d/app/internal/catalog/product/section/adm/usecase/product.go#L115)

![image](docs/resources/packages/c4/app.svg)

### Верхнеуровневая архитектура системы обработки ошибок
![image](docs/resources/hld/c4/diagram.svg)