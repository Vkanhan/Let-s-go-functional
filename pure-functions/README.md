# Pure functions

A pure function is a function where the function return values are identical for identical argument and the function has no side effects.

Although Go functions are not natively pure functions. For instance, functions that handle I/O operations or interact with external systems (such as databases) cannot be pure by their nature. 

But we can write pure function in Go as it helps in predictability, testing and helps in concurrency as they are inherently thread-safe because they do not modify shared state, making them ideal for concurrent execution.
