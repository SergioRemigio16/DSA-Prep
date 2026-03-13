PROBLEM
------------------------------
Title: Products of Array Except Self
Difficulty: 4/10
Date Solved: 3/1/2026


PATTERN
------------------------------
Pattern / Category: Prefix-Sufix cumulative products

Key Signals That Indicate This Pattern:
- Calculating products in an array many times.


CORE INSIGHT
------------------------------
Calculating the Prefix[n] = Prefix[n-1] * nums[n-1]. 
The Previous prefixes' calculation is able to carry over to the next prefix calculation.
This lets the algorithm stay O(n)


SOLUTION SUMMARY
------------------------------
Create a Prefixes and Suffixes array. 
Prefixes[i] contains the products of all the elements to the left of the array.
Suffixes[i] contains the products of all the elements to the right of the array.
Multiply Prefixes[i] * Suffixes[i] to get solution[i]
Return solution



COMPLEXITY
------------------------------
Variables:
- n: number of elements in nums

Time Complexity: O(n)

Space Complexity: O(n)

Explanation: Number multiplications never grow with n. There were only 3 for loops and no nested for loops.
The Prefixes, Suffixes, and solution arrays are all n elements long.


MISTAKE / LESSON
------------------------------
What initially confused me:
I didn't see that the prefixes[i] could be calculated using prefixes[i-1]


What should I recognize faster next time:
Whether a previous calculation can be used to calculate the current step.


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems ChatGPT Examples:
- Parallel processing pipelines where partial aggregates are computed independently.
- Distributed computing (MapReduce-style) prefix/suffix scans.
- Database analytics where cumulative aggregates (running totals, window functions) are computed efficiently.