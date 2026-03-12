PROBLEM
------------------------------
Title: Top K Frequent Elements
Difficulty: 4/10
Date Solved: 3/9/2026


PATTERN
------------------------------
Pattern / Category: Frequency Counting / Hashmap with Bucket Sort

Key Signals That Indicate This Pattern:
- Most times occurred indicates hashmap
- Top frequencies and not value indicates bucket sort 


CORE INSIGHT
------------------------------
Bucket sort can quickly find top frequencies

SOLUTION SUMMARY
------------------------------
Count frequencies of integers in array with hashmap.

Use bucketsort to find top k frequencies. Buckets go from 0 -> n+1 and are for frequencies. Each bucket contains an array with every integer that had that bucket's frequency.

COMPLEXITY
------------------------------
Variables: 
- n = number of elements in nums

Time Complexity: O(n)

Space Complexity: O(n)

Explanation: 
Creating and filling a hashmap frequency map is only O(n) in time and space. 
Creating and filling frequencyBuckets is at most O(n) in time and space.
Finding top k most frequent is at most O(n) in Time.


MISTAKE / LESSON
------------------------------
What initially confused me:
Did not know about bucket sort's ability to sort in O(n) time under certain conditions.


What should I recognize faster next time:
Bucket sort is able to sort in O(n) time under special conditions.


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems:

ChatGPT Examples:
- Token frequency analysis when identifying common tokens in prompts or datasets.
- Trending keyword detection in search engines or social media feeds.
- Log aggregation systems identifying the most frequent error messages or events.