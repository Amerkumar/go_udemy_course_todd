# Interfaces
- Interfaces is bunch of function signatures written together.
- A type implementating a function signature associated with interfaces implicitly implements a interface.
- Now interface type can take those type that implements that interface and call their methods accordingly.
- Interfaces provide us polymorphosim or some people call it code substitutability.
- Inteface is a type that defines only behaviour. Any type that implements that behaviour implements that interface implicitly and then we can pass any type that implements interfaces to the interface type and it will judge accordingly which method to call based on the concrete type. We call the interface type as implementing type because it actually implements the function defined by the interface.
- Concrete Type vs Abstract Type 

# Go Commands
- ```
    golint ./...
    go fmt ./...
    ``

#Empty Interface
-  It is an interface with no method. Every type implements empty interface.

#Method Sets
- Set of methods attached to a type.
- You can value or pointer type to value receiver.
- You can only pass pointer type to pointer receiver.
- You cannot pass value type to pointer receiver, because its address might not exist yet.

# Conversion vs Assertion
- Coversion is converting one type to another like converting int to float.
- Assertion is for interfaces
