PROBLEM
------------------------------
Title: Two Sum
Difficulty: Easy
Date Solved: 3/7/2026


PATTERN
------------------------------
Pattern / Category: Hashset lookup

Key Signals That Indicate This Pattern:
- O(1) lookup time needed


CORE INSIGHT
------------------------------
Looking up whether a number exists in O(1) time removes the need of bloating the run time to O(n^2) Can covert two numbers into one by subtracting one from the other.


SOLUTION SUMMARY
------------------------------
Given nums[i] + nums[j] = target
Store all numbers in hashmap. set target - nums[j] and do a O(1) search for nums[i]. 
Check nums[i] + nums[i] = target. If it doesn't return i and j. 

COMPLEXITY
------------------------------
Time Complexity: O(n)

Space Complexity: O(n) 

Explanation: One hashmap can be created with at most length of n. No more than 3n searches are made. 


MISTAKE / LESSON
------------------------------
What initially confused me:
Didn't realize I could take advantage of hashmap's O(1) search time. 

What should I recognize faster next time:
When a problem needs to look through the array multiple times, check whether a hashmap's O(1) search time can be used.

SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems:

ChatGPT Examples:
-Token deduplication / caching: using a hash table to detect whether a token or embedding has already been processed.
-Memory or object lookup tables: fast retrieval of objects by key rather than scanning arrays.
-Rate limiting systems: quickly checking whether a request identifier already exists in a tracking structure.