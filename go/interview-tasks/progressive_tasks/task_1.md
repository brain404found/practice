# Task-1 
## Пример задачи из реального интервью где в процессе интервьюер постепенно усложняет задачу

### Задача 1: Последовательный вывод чисел
**Проблема**: Вывести числа от 0 до 1000 с задержкой в 1 секунду между каждым числом.
```go
func printNumber(n int) {
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
   
}
```

**Решение**:
```go
func printNumber(n int) {
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    for i := 0; i <= 1000; i++ {
        printNumber(i)
    }
}
```

### Задача 2: Нужно распаралеллить вывод чисел
**Проблема**: Вывести числа параллельно для улучшения производительности.

**Решение**:
```go
func printNumber(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    wg := &sync.WaitGroup{}
    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go printNumber(i, wg)
    }
    wg.Wait()
}
```
**Ключевые моменты**:
- Использование `sync.WaitGroup` для ожидания завершения всех горутин
- Каждая горутина работает независимо
- Основная функция ждет завершения всех горутин

### Задача 3: Ограниченное параллельное выполнение
**Проблема**: Ограничить количество одновременных горутин до 10 для предотвращения перегрузки сервера.

**Решение**:
```go
func printNumber(n int, wg *sync.WaitGroup, sem chan struct{}) {
    defer wg.Done()
    
    // Получение семафора
    sem <- struct{}{}
    defer func() {
        <-sem // Освобождение семафора
    }()

    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    wg := &sync.WaitGroup{}
    sem := make(chan struct{}, 10) // Максимум 10 одновременных горутин

    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go printNumber(i, wg, sem)
    }
    wg.Wait()
}
```
**Ключевые моменты**:
- Использование канала как семафора для ограничения конкурентности
- Максимум 10 горутин работают одновременно
- Семафор освобождается после завершения каждой горутины

### Задача 4: Обработка ошибок в горутинах
**Проблема**: Корректно обработать возможные паники в горутинах.

**Решение**:
```go
func printNumber(n int, wg *sync.WaitGroup, sem chan struct{}) {
    defer wg.Done()
    
    sem <- struct{}{}
    defer func() {
        <-sem
    }()

    time.Sleep(time.Second)
    fmt.Println(n)
    panic("err") // Имитация ошибки
}

func main() {
    wg := &sync.WaitGroup{}
    sem := make(chan struct{}, 10)

    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Восстановление после паники:", r)
                }
            }()
            printNumber(i, wg, sem)
        }()
    }
    wg.Wait()
}
```
**Ключевые моменты**:
- Использование `recover()` в отложенной функции для перехвата паник
- Обработка паники происходит в той же горутине, где она возникла
- Программа продолжает работу даже при панике в некоторых горутинах
- Семафор корректно освобождается даже в случае паники


## Изучаемые концепции

Эти задачи демонстрируют важные концепции Go:
1. Базовое использование горутин
2. Синхронизация с помощью WaitGroup
3. Контроль конкурентности с помощью семафоров
4. Обработка ошибок в конкурентном коде
5. Очистка ресурсов в горутинах
6. Паттерны восстановления после паник

------------------------------------------------------------------------------------

#Progressive Tasks and Solutions

### Task 1: Sequential Number Printing
**Problem**: Print numbers from 0 to 1000 with a 1-second delay between each number.
```go
func printNumber(n int) {
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
   
}
```

**Solution**:
```go
func printNumber(n int) {
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    for i := 0; i <= 1000; i++ {
        printNumber(i)
    }
}
```

### Task 2: Parallel Number Printing
**Problem**: Print numbers in parallel to improve performance.

**Solution**:
```go
func printNumber(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    wg := &sync.WaitGroup{}
    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go printNumber(i, wg)
    }
    wg.Wait()
}
```
**Key Points**:
- Using `sync.WaitGroup` to wait for all goroutines to complete
- Each goroutine runs independently
- Main function waits for all goroutines to finish

### Task 3: Limited Parallel Execution
**Problem**: Limit the number of concurrent goroutines to 10 to prevent server overload.

**Solution**:
```go
func printNumber(n int, wg *sync.WaitGroup, sem chan struct{}) {
    defer wg.Done()
    
    // Acquire semaphore
    sem <- struct{}{}
    defer func() {
        <-sem // Release semaphore
    }()

    time.Sleep(time.Second)
    fmt.Println(n)
}

func main() {
    wg := &sync.WaitGroup{}
    sem := make(chan struct{}, 10) // Maximum 10 concurrent goroutines

    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go printNumber(i, wg, sem)
    }
    wg.Wait()
}
```
**Key Points**:
- Using channel as a semaphore to limit concurrency
- Maximum 10 goroutines running at any time
- Semaphore is released after each goroutine completes

### Task 4: Error Handling in Goroutines
**Problem**: Handle potential panics in goroutines gracefully.

**Solution**:
```go
func printNumber(n int, wg *sync.WaitGroup, sem chan struct{}) {
    defer wg.Done()
    
    sem <- struct{}{}
    defer func() {
        <-sem
    }()

    time.Sleep(time.Second)
    fmt.Println(n)
    panic("err") // Simulated error
}

func main() {
    wg := &sync.WaitGroup{}
    sem := make(chan struct{}, 10)

    for i := 0; i <= 1000; i++ {
        wg.Add(1)
        go func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Recovered from panic:", r)
                }
            }()
            printNumber(i, wg, sem)
        }()
    }
    wg.Wait()
}
```
**Key Points**:
- Using `recover()` in a deferred function to catch panics
- Panic handling is done in the same goroutine where it occurs
- Program continues running even if some goroutines panic
- Semaphore is properly released even in case of panic

## Learning Points

These tasks demonstrate important Go concepts:
1. Basic goroutine usage
2. Synchronization with WaitGroup
3. Concurrency control with semaphores
4. Error handling in concurrent code
5. Resource cleanup in goroutines
6. Panic recovery patterns 
