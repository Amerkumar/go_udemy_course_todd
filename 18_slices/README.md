#Slices

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

* make
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