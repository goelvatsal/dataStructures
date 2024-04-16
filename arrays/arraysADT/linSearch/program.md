**Input:**
1. Array structure with
    1. Array with [20]int type
    2. Length of array
    3. Number of elements in array
2. Number to lookup in array (key)

**Output:**
1. Index of requested array
2. Or -1 if number was not found

**Process:**
1. Creates a for loop that ends at the length
2. For every iteration, it compares the key to the current element
3. If match is found, then it returns the current index
4. If match is not found, then it returns -1