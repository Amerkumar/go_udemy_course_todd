# Intro to Functions
    * main is the entry point of program
    * only one main in entire program
    * you can't have packages without main
    * func [receiver] identifier(parameters) returns
    * Difference between parameter and arguments
    * Difference between variable and functions * greet and greet()


# Func Returns:
    * Func return
    * multiple return
    * named return

# Variadic Functions:
    * The final parameter in a function signature may have type prefixed with ...A function with such a parameter is called variadic
    * and may be invoked with zero or more arguments
    * Difference between variadic parameter and variadic arguments

# Func expression:
    * When you assign a variable to a function is called func expression
    * Only way to add a function inside a function

# Closure:
    * Closure helps us the limit the scope of variables used by multiple functions
    * Without closure, for two or more functions to have access to the same variable that variable need to have package scope
    * package level scope
    * block level scope
    * lexical scope refers to its local lexical environment.
    * name binding is the association of data or entries with identifiers.
    * Closures are techniques for implementing lexically scoped name binding in languages with first class function.
    ```
    function startAt(x)
       function incrementBy(y)
           return x + y
       return incrementBy

    variable closure1 = startAt(1)
    variable closure2 = startAt(5)
    ```
    * closure is a record storing a function together with an environment.
    * A mapping associating each free variable of the function (variables that are defined locally, in an enclosing scope) with the value
    * or reference to which the name was bound when the closure was created.


# Callbacks
        * Function is also a type
        * passing a function like an argument
        * Calling back that has called it.
        * Closure is return a function while callback is taking a function as an argument
        ```
        callback =  func(n int) {
        		fmt.Println(n)
        	}
        func visit(numbers []int, callback func(int))  {
        	for _, n := range numbers{
        		callback(n)
        	}
        }
        ```
        * functional programming is not main domain of go,but you can use it.
        * Keep on the side of clarity


# Recursion
    * Just like infinite mirror.
    * Defined in terms of itself
    * Functions that calls itself

# Defer
    * It is a keyword