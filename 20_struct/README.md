# Structs
Documentation - GOLANG SPEC

- It is composite type
- Whenever a user defined type it must have an underlying type.
- Go is statically typed
- A struct is a sequence of named elements , called fields, each of which has a name and type. Field names may be specified explicitly or implicitly. Within a struct , non-blank field names must be unique.
- A field declared with a type but no explicit field name is an anonymous field also called embedded field or an embedding ofthe type in the struct. An embedded type must be specified as a type name T or as a pointer to a non-interface type name *T and T itself may not be pointer type. The unqualified type name acts as a field name.
- Promoted 

# Object Oriented In Go

Go is object oriented

- Encapsulation
    - state
    - behaviour
    - exported / unexported
    
- Reusability 
    - inheritance("embedded types")

- Polymorphosim
    - interfaces
    
- Overriding
    - "promotion"
    
* Traditional OOP
    - Classes are blueprints
    - You instantiate objects from them.
    
* OOP in GO
    - you don't create classes, you create a type.
    - you create values of that type.

* User Defined Types
- It is bad practice to alias unless you want to attach methods to it.

* Commpostion 
    - Review -> Methods ,Embedded Types, struct pointers
    - Receiver attach function to type and this is called **methods** to type.
    - Type embedded in another type is called **embedded types.** Inner type is promoted to outer type.
    
    
* JSON
    - Definition
    - Encoding
        - Marshal/ UnMarshal   
            - string
        - encode / decode
            - stream
            
    - Marshal
        - string to json conversion
        - In case of stream, it is called encode,
    - UnMarshal
        - json to string conversion
        - In case of stream, it is called decode.
   