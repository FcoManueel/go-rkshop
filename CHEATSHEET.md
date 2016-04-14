# Go in a taco shell
Cheatsheet by Manuel Valle, 2016  
## Generalities
- **Paradigm:** Imperative, structured, concurrent.
- **Typing:** Static (a string will always be a string), Strong (no implicit conversion).
- **Syntax:** C-like (with curly braces, but a lot less parenthesis and semicolons).
- **Execution:** Fast! Compiles to native code (no bytecode or VM).
- **Memory management:** Garbage collected. Pointers exist, but not pointer arithmetic.
- **Functions:** Are first class citizens. Can return multiple values. Can be closures.
- **Classes:** No classes. And no inheritance! Just good ol' composition.
- **Concurrency primitives:** Goroutines and Channels.
- **Godfather:** Google.

## Types

```go
type Hobby string

type Dog struct {
    Name      string  // Public scope (first letter capitalized)
    age       int     // Package-only (first letter underscore)
    Hobbies   []Hobby // User-defined types can be used here as well
}
```

## Variables

```go
const maxWaitingTime = 1500    // Scope is defined by first-letter case
var iDontKnowTheValueYet Hobby // Declaration formula is `var NAME TYPE`
iAlreadyKnowTheValue := 1500   // Declare+initialize at once using `:=`
lassie := Dog{"Lassie", 100, []Hobby{"Greek art", "Skydiving"}}
```
  
<br><br><br>
  
## Arrays and Slices

```go
a := [2]string{"uno", "deux"} // Arrays have a defined size, and it's fixed
s := []int{0, 1, 2, 3, 4, 5}  // Slices are what we'll need 99% of the time

s[1:4] // == [1 2 3]  // A slice can be 're-sliced' using `[BEGIN:END]`
s[:3]  // == [0 1 2]  // END is *not* included (i.e. 3 is not in the outcome)
s[3:]  // == [3 4 5]  // BEGIN is included (i.e. 3 is in the outcome)

var foo []int32 // Declared, not initialized. It doesn't have memory assigned. It's a 'nil slice',
// foo[0]       // Can't do. Trying to access a nil slice causes a panic error
foo = make([]int32, 5) // Here we allocate memory for 5 elements of type int32
foo[3] = 3 // foo == [0 0 0 3 0]

```

## Maps
```go
m[key] = elem     // Assign
delete(m, key)    // Delete
elem, ok = m[key] // Reading: `ok` is a bool that indicates if `elem` was found in the map


// Similarly as with slices, we can't assign to nil maps. Be sure to initialize
employeeByID := make(map[string]Employee)
employeeByID["A123"] = Employee{"A123", "Malandro Gutierrez", 35}
e, ok := employeeByID["A123"] // ok == true, e == Malandro's Employee struct
delete(employeeByID, "A123")
e, ok = employeeByID["A123"]  // ok == false, e == Empty employee struct

```

## Ranges

```go
array := []string{"a", "b", "c"}
for i := range myArr {    // We can range through slice/array indexes
    fmt.Println(i) // 0 1 2
}

for i, v := range array { // Or through both indexes and values
    fmt.Printf("%d:%s  ", i, v) // 0:a  1:b  2:c
}

for _, value := range array { //Or only values! The `_` is a 'blank identifier'
    
```
<br>

## Functions

```go
// Without parameters, without return value
func helloWorld() {
	fmt.Println("Hola world!")
}

// With parameters, with return value
func Greet(name string) string {
	return "Hola " + name + "!"
}

// Function defined on existing type (aka 'method')
func (d Dog) Greet(name string) string {
	return fmt.Sprintf("%s the dog says 'Hola!' to %s", d.Name, name)
}

// Multiple return values (and error handling example)
func getHonorTitle(gender string) (string, error) {
	switch gender {
	case "male":
	   return "Monsieur ", nil
	case "female":
       return "Madame ", nil
    default:
    	return "", error.New("No appropriate honor title was found")
	}
}

// Tricky one, a function returning a function
func fibonacci() func() int {
    x, y := 0, 1
	return func() int {
		x,y = y,x+y
		return x
	}
}
```

<br><br><br>
<br><br><br>
<br>

### Interfaces

```go
// An interface type is defined by a set of methods.
type Greeter interface {
    Greet(name string) string
}
var g Greteer          // g can hold any type satisfying the Greeter interface
g = Dog{Name: "Jake"}  // we previously defined `(d Dog) Greet(string) string`
// Interfaces are satisfied implicitly. No need to say "Dog implements Greeter"

// Let's asume we have a Person struct defined
func (p Person) Greet(name string) string {
	return fmt.Sprintf("%s the Humabn says 'Woof!' to %s", p.Name, name)
}

// Person is also a Greeter, so we can assign it to `g`
g = Person{Name: "Finn"}

```
