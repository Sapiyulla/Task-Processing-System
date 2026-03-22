# Task

## Описание
Задача - это `aggregate root`, который является центральной сущностью.

## Атрибуты:
 - `id` - уникальный идентификатор
 - `type` - тип задачи (_подробнее см. [типы задачи](#типы-задач)_)
 - `payload` - дополнительные параметры **`[json]`**
    > Нужен для определения контекстных аргументов, дополнительных данных и тд.
 - `status` - текущее состояние (_подробнее см. [статусы задач](#статусы-задач)_)

### Типы задач
type | description | abbreviature
-|-|-
*AI Picture Prepare* | Обработка изображения с ИИ | `ai_picture_prepare`
*Data Processing* | Обработка данных | `data_processing`
*Report Generate* | Составление заявления | `report_generate`

### Статусы задач

status | description
-|-
`CREATED` | Задача создана
`WAIT` | Задача ожидает выполнения
`QUEUED` | Задача в очереди на выполнение
`IN_PROGRESS`| Задача выполняется
`COMPLETED` | Задача завершена
`FAILED` | Задача завершена с ошибкой

## Команды

command | description
-|-
`QueueTask` | Отправка в очередь
`MarkTaskReadyForDelivery` | Обозначить задачу как готовую к исполнению
`MarkTaskCompleted` | Обозначить задачу как выполненную
`MarkTaskFailed` | Обозначить задачу как проваленную

## State Machine

FROM    |    → COMMAND         |            → TO  
-|-|-
 \- | → CreateTask | → CREATED
CREATED |    → MarkTaskReadyForDelivery   | → WAIT  
WAIT     |   → QueueTask            |       → QUEUED  
QUEUED   |   → StartExecution       |       → IN_PROGRESS  
IN_PROGRESS | → CompleteTask            |    → COMPLETED

## Инварианты

- Нельзя завершить задачу не из `IN_PROGRESS`
- Нельзя перевести `COMPLETED` в любое другое состояние
- Нельзя ретраить `FAILED` (если NotRetryableError)
- Нельзя утанавливать `payload` более чем **10МБ**