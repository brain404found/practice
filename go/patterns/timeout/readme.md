# Timeout Pattern in Go

This pattern demonstrates how to implement timeouts in Go using context and select statements.

## Overview

The timeout pattern is crucial for:
- Preventing infinite operations
- Managing resource usage
- Improving application responsiveness
- Handling external service calls

## Implementation

```go
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

ticker := time.NewTicker(time.Second)
defer ticker.Stop()

for {
    select {
    case <-ctx.Done():
        return
    case <-ticker.C:
        fmt.Println("tick")
    }
}
```

## Key Components

1. **Context with Timeout**
   - `context.WithTimeout` creates a context that automatically cancels after specified duration
   - `defer cancel()` ensures proper cleanup

2. **Select Statement**
   - Monitors multiple channels simultaneously
   - Handles both timeout and normal operation cases

3. **Ticker**
   - Provides regular intervals for operations
   - Should be stopped with `defer ticker.Stop()`

## Common Use Cases

1. **API Calls**
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

resp, err := http.GetWithContext(ctx, "https://api.example.com")
```

2. **Database Operations**
```go
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM users")
```

3. **External Service Calls**
```go
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()

result, err := externalService.ProcessWithContext(ctx, data)
```

## Best Practices

1. **Always Clean Up**
   - Use `defer cancel()` to prevent context leaks
   - Stop tickers and timers when done

2. **Reasonable Timeouts**
   - Set timeouts based on operation type
   - Consider retry strategies for failed operations

3. **Error Handling**
   - Check for `context.DeadlineExceeded`
   - Implement graceful fallbacks

4. **Resource Management**
   - Close connections and resources
   - Clean up goroutines

## Example with Error Handling

```go
func processWithTimeout(ctx context.Context, duration time.Duration) error {
    ctx, cancel := context.WithTimeout(ctx, duration)
    defer cancel()

    done := make(chan error, 1)
    go func() {
        // Simulate work
        time.Sleep(2 * time.Second)
        done <- nil
    }()

    select {
    case err := <-done:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

## Testing Timeouts

```go
func TestTimeout(t *testing.T) {
    ctx := context.Background()
    err := processWithTimeout(ctx, 1*time.Second)
    
    if err != context.DeadlineExceeded {
        t.Errorf("expected timeout error, got %v", err)
    }
}
```

## Common Pitfalls

1. **Forgetting to Cancel**
   - Always call cancel() to prevent resource leaks
   - Use defer for guaranteed cleanup

2. **Too Short Timeouts**
   - Can cause unnecessary failures
   - Consider operation complexity

3. **Too Long Timeouts**
   - May block resources unnecessarily
   - Can impact system responsiveness

4. **Ignoring Context**
   - Always pass context to child operations
   - Check for context cancellation

## Related Patterns

- Circuit Breaker
- Retry Pattern
- Rate Limiting
- Graceful Shutdown
