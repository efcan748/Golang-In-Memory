# Store Package API Documentation

## Overview
The store package provides an in-memory key-value store with support for strings and lists, including TTL (Time-To-Live) functionality.

## Initialization
```go
New() *Store
```
Creates a new Store instance.

```go
store := store.New()
StartCleanup(interval time.Duration) *Store
```
Starts a background goroutine to clean up expired entries.

```go
store := store.New().StartCleanup(1 * time.Minute) // Cleanup every minute
```

## String Operations

### - `func Set(key, value string, ttl time.Duration) (string, bool)`

Sets a string value with optional TTL.

#### Parameters:

`key`: The key to set

`value`: The value to store

`ttl`: Time-To-Live duration (0 for no expiration)

#### Returns:

`(errorMessage, success)` - If success is false, errorMessage contains the reason

#### Example:

```go
if msg, ok := store.Set("user:1", "John Doe", 10*time.Minute); !ok {
    fmt.Println("Error:", msg)
}
```

### - `func Get(key string) (string, bool)`
Retrieves a string value.

#### Returns:

`(value, exists)` - If exists is false, the key wasn't found or was expired

#### Example:

```go
if value, exists := store.Get("user:1"); exists {
    fmt.Println("Value:", value)
}
```

### - `func Update(key, value string, ttl time.Duration) (string, bool)`

Updates an existing string value.

#### Parameters: 
Same as Set

#### Returns: 
Same as Set

#### Example:

```go
if msg, ok := store.Update("user:1", "Jane Doe", 15*time.Minute); !ok {
    fmt.Println("Error:", msg)
}
```

### - `func Delete(key string) bool`
Deletes a string value.

### Returns: 
`true` if the key existed and was deleted

### Example:

```go
if deleted := store.Delete("user:1"); deleted {
    fmt.Println("Key deleted")
}
```

## List Operations 
### - `func LPush(key string, ttl time.Duration, values ...string) (string, bool)`

Pushes values to the beginning of a list.

#### Parameters:

`key`: The list key

`ttl`: Time-To-Live for the list

`values`: One or more values to push

#### Returns:

`(errorMessage, success)`

#### Example:

```go
if _, ok := store.LPush("tasks", 1*time.Hour, "task1", "task2"); !ok {
    fmt.Println("Error pushing to list")
}
```

### - `func Pop(key string) (string, bool)`
Pops a value from the end of a list.

#### Returns:

`(value, success)` - If success is false, value contains the error reason

#### Example:

```go
if value, ok := store.Pop("tasks"); ok {
    fmt.Println("Popped:", value)
} else {
    fmt.Println("Error:", value)
}
```

# Data Structures
## Store
The main store structure containing:

- Strings `map[string]stringEntry` - String storage

- Lists `map[string]listEntry` - List storage

- mu `sync.RWMutex` - Synchronization mutex

### `type stringEntry`
Contains:

- Value `string` - The stored string

- expiration `time.Time` - Expiration timestamp

### `type listEntry`
Contains:

- Values []string - The list items

- expiration time.Time - Expiration timestamp

## Error Handling
Common error messages returned:

`"Key %s is already exist"` - When setting an existing key

`"key %s not found"` - When updating a non-existent key

`"entity is expired"` - When operating on expired entries

`"key is not exist"` - When popping from non-existent list

`"list is empty"` - When popping from empty list

`"list entity is expired"` - When operating on expired list

#### Example Usage
```go
package main

import (
	"fmt"
	"time"
	"yourmodule/store"
)

func main() {
	s := store.New().StartCleanup(1 * time.Minute)
	
	// String operations
	s.Set("user:1", "Alice", 10*time.Minute)
	val, exists := s.Get("user:1")
	fmt.Println(val, exists)
	
	// List operations
	s.LPush("tasks", 1*time.Hour, "clean", "code")
	task, ok := s.Pop("tasks")
	fmt.Println(task, ok)
}
```
