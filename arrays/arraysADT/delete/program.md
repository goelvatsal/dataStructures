**Input:**
1. Array structure with
   2. Array with [20]int type
   3. Length of array
   4. Number of elements in array
5. Index of element to delete in array

**Output:**
1. Modified array structure with element with index removed from it

**Process:**
1. Checks if the index given is greater than or equal to 0
2. Runs a for loop that starts with the index and ends at the last element
3. Changes each element to the following element
4. Decreases the array's length
5. Sets the last element to 0
6. Efficiency of O(1) min, O(n) max