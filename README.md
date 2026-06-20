# Pozorishe

Ну тут короче задачки по основам голанга до хттп 
Код лучше даже не читайте

# Этап 1. Типы, указатели, методы

Цель: перестать путаться в value/reference semantics.

## Задача: банковский счет

Создай структуру:

```go
type Account struct {
    ID      int
    Owner   string
    Balance float64
}
```

Реализуй методы:

```go
Deposit(amount float64)
Withdraw(amount float64) error
```

После этого ответь себе на вопросы:

* почему методы должны иметь pointer receiver?
* что будет если receiver сделать value?
* почему изменение поля внутри метода иногда не сохраняется?

### Специальные ошибки

Попробуй намеренно сделать:

```go
func (a Account) Deposit(amount float64)
```

и объяснить результат.

---

# Этап 2. Слайсы и память

Это одна из самых частых тем на собеседованиях.

## Задача: менеджер задач

```go
type Task struct {
    ID   int
    Name string
}
```

Сделай функции:

```go
AddTask(...)
DeleteTask(...)
FindTask(...)
```

Обязательно удаление через слайсы:

```go
tasks = append(tasks[:i], tasks[i+1:]...)
```

После этого проверь:

```go
fmt.Println(len(tasks))
fmt.Println(cap(tasks))
```

и разберись:

* чем len отличается от cap
* когда append создает новый массив
* когда использует старый

### Эксперимент

Напечатай адреса:

```go
fmt.Printf("%p\n", &tasks[0])
```

до и после append.

---

# Этап 3. Карты

## Задача: подсчет слов

На вход:

```go
[]string{
    "go",
    "java",
    "go",
    "python",
    "go",
}
```

На выход:

```go
map[string]int
```

Потом усложни:

Сделай рейтинг слов по популярности.

Тут потренируешь:

* map
* сортировку
* структуры

---

# Этап 4. Интерфейсы

Очень важный этап.

## Задача: фигуры

```go
type Shape interface {
    Area() float64
}
```

Реализуй:

```go
Circle
Rectangle
Triangle
```

Сделай функцию:

```go
func PrintArea(s Shape)
```

Потом ответь:

* что такое interface value?
* чем отличается nil interface от interface с nil внутри?

Это классическая ловушка Go.

---

# Этап 5. Ошибки

## Задача: файловый менеджер

Функции:

```go
CreateFile(...)
ReadFile(...)
DeleteFile(...)
```

Нельзя:

```go
panic(...)
```

Нужно:

```go
return err
```

Отдельно потренируй:

```go
errors.New()
fmt.Errorf()
errors.Is()
errors.As()
```

---

# Этап 6. Структуры и композиция

В Go нет наследования.

## Задача: сотрудники компании

```go
type Person struct {}
type Developer struct {}
type Manager struct {}
```

Через embedding:

```go
type Developer struct {
    Person
}
```

Пойми:

* promotion methods
* composition over inheritance

---

# Этап 7. Каналы

Начинается самое интересное.

## Задача: генератор чисел

Горутина:

```go
go generate(...)
```

Канал:

```go
chan int
```

Поток:

```text
generator -> channel -> consumer
```

Потренируй:

* unbuffered channel
* buffered channel
* закрытие канала

---

# Этап 8. Worker Pool

Обязательная задача.

## Реализовать

```text
jobs -> workers -> results
```

Например:

```go
square(2) -> 4
square(3) -> 9
```

3 воркера.

Тут всплывут:

* deadlock
* close(channel)
* range по каналу

---

# Этап 9. Fan-Out Fan-In

Следующий уровень.

```text
input
  |
fan-out
 / | \
w1 w2 w3
 \ | /
fan-in
  |
output
```

Например обработка строк.

---

# Этап 10. Sync primitives

## Задача: безопасный счетчик

Сначала сделай неправильно:

```go
counter++
```

из 1000 горутин.

Увидишь гонку.

Потом исправь через:

```go
sync.Mutex
```

После:

```go
sync.RWMutex
```

После:

```go
atomic
```

---

# Этап 11. Context

Хотя HTTP еще нет, context уже нужен.

## Задача

Есть горутина:

```go
for {
    ...
}
```

Через:

```go
context.WithCancel(...)
```

остановить ее.

Потом:

```go
context.WithTimeout(...)
```

---

# Финальный проект

Когда все этапы пройдены.

## Concurrent Task Manager

Структуры:

```go
Task
Worker
TaskManager
```

Возможности:

* добавить задачу
* удалить задачу
* список задач
* обработка задач воркерами
* отмена через context
* логирование ошибок
* безопасный доступ через mutex

В одном проекте ты используешь:

* struct
* methods
* interfaces
* errors
* slices
* maps
* goroutines
* channels
* mutex
* atomic
* context



