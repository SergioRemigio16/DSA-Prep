PROBLEM
------------------------------
Title: Group Anagrams
Difficulty: 5/10
Date Solved: 3/9/2026


PATTERN
------------------------------
Pattern / Category: Array as Key / Hashmap

Key Signals That Indicate This Pattern:
- Having to compare anagrams 
- Efficiency of Hashmaps


CORE INSIGHT
------------------------------
Take advantage of lowercase words only having up to 26 character.

SOLUTION SUMMARY
------------------------------
Turn words into character frequency arrays and put them in a hashmap as the key. 
Make the value an array of all the original words that form the same character frequency array.
Iterate through hashmap and return all the character frequency arrays. 

COMPLEXITY
------------------------------
n = number of words
s = word length

Time Complexity: O(ns)

Space Complexity: (s)

Explanation: 
Turning words into character frequency array is O(s).
Going through every word O(n)


MISTAKE / LESSON
------------------------------
What initially confused me: I did not realize I could take advantage of the limited number of letters in the alphabet to make a character frequency array for words.


What should I recognize faster next time: Character frequency arrays exists.


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems, ChatGPT Examples:
-Grouping prompts with identical token distributions for caching.

-Detecting prompt variations that are permutations of the same keywords.

-Clustering similar queries in semantic preprocessing pipelines.