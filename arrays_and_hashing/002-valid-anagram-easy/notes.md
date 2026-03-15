PROBLEM
------------------------------
Title: Is Anagram
Difficulty: 2/10
Date Solved: 3/6/2026


PATTERN
------------------------------
Pattern / Category: Hashmap frequency counting

Key Signals That Indicate This Pattern:
- Realizing duplicates could be found with a single pass


CORE INSIGHT
------------------------------
The strings don't have to be sorted to be compared.

SOLUTION SUMMARY
------------------------------
Place both strings in two different hashmaps. Compare whether the hashmaps are equal.

COMPLEXITY
------------------------------
Time Complexity: O(n+m), where n is the length of string s and m is the length of string t

Space Complexity: O(1), a map with at most 26 elements doesn't increase as n or m increase.

Explanation: Two hashmaps are created and every character in both string s and t don't have to be looked at more than 2 times.


MISTAKE / LESSON
------------------------------
What initially confused me: 
O(1) space complexity confused me. I thought more memory couldn't allocated. This is not true. A simple hashmap with size of the alphabet is O(1) space complexity.


What should I recognize faster next time: If a solution feels right don't immediately discard it. 


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems:
Frequency counting and hash maps appear frequently in data processing pipelines and text analysis.

ChatGPT Examples:
-Detecting duplicate or equivalent records in data pipelines.

-Text analytics (word frequency counting).

-Log analysis where event frequencies are tracked.

-Cache invalidation or request deduplication using hashed keys.

-Security systems comparing normalized inputs (e.g., password or token validation patterns).