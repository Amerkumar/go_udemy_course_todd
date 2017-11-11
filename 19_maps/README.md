# Maps
- Maps are key value associations just like dictionary.
- Structs are composite types - Bunch of different variables together.
- an unordered group of elements of one type, called the element type.
    * indexed by the set of unique keys of another type called the key type.
    * The value of unintialized type is nil.
    -  **It is reference type** 
    - **Channels are also reference type.**
    - Maps store pointer to the underlying data.
    - The underlying data structure of maps is hashtable.
- Underlying a slice is an array whereas underlying is a hashtable.    
    
# Making Slices
```go
      m := make(map[string]int)
      n := map[string]int{"foo" : 1, "bar" : 2} // Composite Literal
      o := map[string]int{}
```  

# Resources
- Caleb Doxsey Book 
- Language Spec 
- Effective Go

# Best Books
1. Go Programming Language By Alan
2. Go in action
3. The way to Go

# Github Pull
- Two computers and a cloud server for code repository.    

# Maps Resources
- Ardan Labs - Macro View of Maps Internals in GO
- Go Blog - go maps in action


# Hash Tables
- Unsorted to Sorted to Buckets Analogy
- Buffer Analogy Printer
