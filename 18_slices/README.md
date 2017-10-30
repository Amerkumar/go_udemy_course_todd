# Slices

* A slice is a descriptor for a contiguous segment of an underlying array and
* provides access to a numbered sequence of elements from that array.
* The value of uninitialized slice is nil.
* **It is a reference type.**
* Like arrays , slices are indexable and have a length.
    * Unlike arrays, slices are dynamic.
    * Their length may change during execution.
* A slice once intialized is always associated with an underlying array that holds 
* its elements.
* The arrays underlying a slice extend past the end of a slice.
    - Capcity is the measure of that extent.
        - It is the sum of the length of the slice and length of array beyond the slice.
    - The capcity of a slice can be discovered using cap(a)
    - When we exceed a capcity, the compiler creates a new underlying of greater size.
        - It carries its cost.
        - It effectively doubles the array size each time they are created(2,4,8,16) upto 
        - a certain point and after that they scale in smaller proportion.
        
 ** we can slice a slice. **
 
 ## creating a slice
    - shorthand
    - var 
    - make


$$ make
    - Two ways to create a slice
        - ```go
            mySlice := []int{1,3,5,7,9,11,} // Cannot specify capaciry.
            mySlice := make([]int, 0, 5)
            ```
    - A new intialized slice value for given element type T is made using the 
    - built-in function make, which makes a slice type and parameters specifying
    - the length and optionally the capacity.
    
    - A slice created with make always allocates a new hidden array to which the
    - returned slice value refers.
        - make([]T, length, capacity)
        - make([]int, 50, 100)
            - same as this : new([100]int)[0:50]
    
    - Like arrays, slices are one dimensional but may be composed to construct 
    - to higher dimensional objects. (multi-dimensional slices)
    - basic slices.
    
    
$$ append 
    - append an element to slice
    - if beyound capacity, append will double the size of slice.
    - We can also append a slice with another slice.
    - We can also delete an element using append.