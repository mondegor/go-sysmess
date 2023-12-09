# Описание GoSysMess v0.5.3
Этот репозиторий содержит описание библиотеки GoSysMess.

## Статус библиотеки
Библиотека находится в стадии разработки.

## Описание библиотеки
Библиотека решает две основные задачи:
- формирует пользовательские сообщения на различных языках, в том числе и сообщения об ошибках;
- даёт инструменты для более удобной обработки ошибок как пользовательских, так и программных (более подробно см. ниже);

## Подключение библиотеки
go get github.com/mondegor/go-sysmess

## Обработка ошибок. Общие сведения
### Выделяются три вида ошибок:
- Пользовательские ошибки (user) - все те ошибки, которые должен исправить сам пользователь (примеры: недопустимый ввод данных; обращение к несуществующему ресурсу);
- Системные (system) - происходящие в системе и независящие от программы (примеры: отсутствует соединение с БД или с каким либо внешнем API; отсутствие прав на запись файлов и т.д.);
- Программные (internal) - любые другие ошибки, допущенные разработчиками, или ошибки, которые не были классифицированы предыдущими видами (примеры: обращение к нулевому указателю, создание файла в несуществующей папке; выход из диапазона значений массива);

### Действия при обработке пользовательской ошибки:
- понятно объяснить причину возникновения;
- предложить варианты действий;
- возможно, превратить ошибку в пользовательский сценарий (например, если логин уже существует, то сгенерить похожие варианты и предложить пользователю выбрать один из них);

### Действия при обработке системной ошибки:
- понятно объяснить причину возникновения;
- предложить варианты действий, если возможно;
- сообщить, с чем обратиться в поддержку;
- отобразить уникальный код ошибки, под которым система предварительно записала эту ошибку в логи;
- узнавать об ошибках автоматически, а не от пользователя (т.к. не все пользователи обращаются в поддержку, да и оперативное исправление ошибок увеличивает лояльность пользователей);

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
- team lead - отслеживает всё ли работает нормально; выявляет новые ошибки, принесённые релизом; оценивает влияние ошибок на пользователей;

## Архитектура обработки ошибок
Внутренний архитектурный слой не должен влиять на поведение верхнего, поэтому ошибки должны быть частью интерфейса каждого слоя. Определение типов ошибок должны быть в том же слое где возникает данная ошибка. Причём для каждой ошибки должен быть указан её вид: user, system, internal.

Основная обработка (перехват) ошибок должна происходить в слое UseCase, только он знает как именно обрабатывать ошибки поступивших из других слоёв, а также правильно определить нужный вид ошибки. Конкретный слой, обрабатывая ошибку, должен создать объект с ошибкой, определённую в этом слое, а если этот слой обрабатывает уже перехваченную ошибку, то он должен вложить её в созданную им ошибку, для того, чтобы внешний перехватчик смог обработать всю цепочку ошибок и выполнить необходимые действия для каждого вида ошибок.

Чтобы избавить код от текстов пользовательских ошибок, а также одновременно решить проблему локализации, все тексты пользовательских ошибок необходимо хранить отдельно и загружать их по заранее выданным им идентификаторам.