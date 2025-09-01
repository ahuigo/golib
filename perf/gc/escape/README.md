# Go 逃逸分析详解

本文档展示了各种导致变量逃逸到堆上的情况，特别是 unsafe.Pointer 相关的逃逸行为。

## 逃逸分析结果总结

### 1. 基础类型逃逸

| 函数 | 逃逸情况 | 原因 |
|------|----------|------|
| `returnInt()` | ❌ 不逃逸 | 直接返回值类型 |
| `returnInterface()` | ✅ 逃逸 | 接口转换需要在堆上创建类型信息 |
| `returnPointer()` | ✅ 逃逸 | 返回局部变量指针 |
| `localInterface()` | ✅ 逃逸 | 即使本地使用，接口转换仍逃逸 |

### 2. Unsafe Pointer 逃逸情况

| 函数 | 逃逸情况 | 逃逸原因 |
|------|----------|----------|
| `returnUnsafePointer()` | ✅ 逃逸 | 返回指向局部变量的 unsafe.Pointer |
| `unsafeTypeConversion()` | ✅ 逃逸 | unsafe.Pointer 类型转换后返回指针 |
| `unsafePointerToUintptr()` | ✅ 逃逸 | 通过 unsafe.Pointer 取地址 |
| `localUnsafePointer()` | ✅ 逃逸 | fmt.Printf 需要接口转换 |
| `unsafeMemoryManipulation()` | ✅ 部分逃逸 | 循环变量和解引用值用于打印 |

## 关键发现

### 1. unsafe.Pointer 的逃逸特点

```go
func returnUnsafePointer() unsafe.Pointer {
    p := 42                    // 局部变量
    return unsafe.Pointer(&p)  // ✅ p 逃逸: moved to heap
}
```

**原因**: 一旦使用 `&p` 取地址并通过函数返回，局部变量必须移到堆上。

### 2. 类型转换链式逃逸

```go
func unsafeTypeConversion() *int {
    p := 100                       // 局部变量
    ptr := unsafe.Pointer(&p)      // 取地址
    return (*int)(ptr)             // ✅ p 逃逸: moved to heap
}
```

**原因**: 即使经过 `unsafe.Pointer` 中间转换，最终返回的仍是指向局部变量的指针。

### 3. uintptr 转换也会逃逸

```go
func unsafePointerToUintptr() uintptr {
    p := 200                       // 局部变量  
    ptr := unsafe.Pointer(&p)      // 取地址
    return uintptr(ptr)            // ✅ p 逃逸: moved to heap
}
```

**原因**: 虽然返回的是 `uintptr` 值，但编译器在取地址时就决定了逃逸。

### 4. 本地使用 unsafe.Pointer 仍可能逃逸

```go
func localUnsafePointer() {
    p := 300                                    // 局部变量
    ptr := unsafe.Pointer(&p)                   // 取地址
    fmt.Printf("...", ptr, *(*int)(ptr))        // ✅ p 逃逸: moved to heap
}
```

**原因**: 
- `ptr` 作为参数传递给 `fmt.Printf`，需要接口转换
- `*(*int)(ptr)` 解引用的值也需要传递给 `fmt.Printf`

## 逃逸分析的核心规则

1. **取地址操作**: 任何 `&variable` 操作都可能导致逃逸
2. **跨函数边界**: 指针、接口类型跨函数传递必然逃逸  
3. **接口转换**: 任何到 `interface{}` 的转换都会逃逸
4. **函数参数**: 传递给复杂函数（如 fmt.Printf）的参数容易逃逸

## 性能优化建议

1. **避免不必要的接口转换**
2. **减少指针的跨函数传递**
3. **使用值类型而非指针类型（当可能时）**
4. **谨慎使用 unsafe.Pointer，它几乎总是会导致逃逸**
