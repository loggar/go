# basic types

```
bool
string
Numeric types:
uint either 32 or 64 bits
int same size as uint
uintptr an unsigned integer large enough to store the uninterpreted bits of
a pointer value
uint8 the set of all unsigned 8-bit integers (0 to 255)
uint16 the set of all unsigned 16-bit integers (0 to 65535)
uint32 the set of all unsigned 32-bit integers (0 to 4294967295)
uint64 the set of all unsigned 64-bit integers (0 to 18446744073709551615)
int8 the set of all signed 8-bit integers (-128 to 127)
int16 the set of all signed 16-bit integers (-32768 to 32767)
int32 the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64 the set of all signed 64-bit integers
(-9223372036854775808 to 9223372036854775807)
float32 the set of all IEEE-754 32-bit floating-point numbers
float64 the set of all IEEE-754 64-bit floating-point numbers
complex64 the set of all complex numbers with float32 real and imaginary parts
complex128 the set of all complex numbers with float64 real and imaginary parts
byte alias for uint8
rune alias for int32 (represents a Unicode code point)
```

## Size of integers

```go
var i int
var i32 int32

fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(i))
fmt.Printf("Size of int32: %d bytes\n", unsafe.Sizeof(i32))
```

The problem with converting from int to int32 in Go arises due to the potential loss of data. The int type can be either 32 bits or 64 bits depending on the system architecture (32-bit or 64-bit). When converting a larger int value to int32, there is a risk that the value may exceed the range that int32 can represent, which is from -2,147,483,648 to 2,147,483,647. This can lead to overflow and incorrect values.

```go
var largeInt int = 3000000000 // This is within the range of int on a 64-bit system
var smallInt32 int32 = int32(largeInt) // This will cause overflow and result in an incorrect value
```
