# Generic type for struct methods(error): counterintuitive  && cannot use (untyped string constant) as string value in assignment.

There is an exmaple: https://go.dev/play/p/guzOWRKi-yp

```
func (c *cachedFn[string, V]) Get0() (V, error) {
	// var s any
	var s string
	s = "abc" // error: cannot use "abc" (untyped string constant) as string value in assignment
	fmt.Printf("cache key: %#v, %T\n", s, s) // cache key: 0, uint8
	return c.Get(s)
}
```

I find the generic type of the struct method a bit confusing.

1. `(c *cachedFn[string, V])` does not really constrain the type to string.  It's actual type is uint8.
2. This error(`s = "abc" // error: ...` ) is a bit counterintuitive.
