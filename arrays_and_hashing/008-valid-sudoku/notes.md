PROBLEM
------------------------------
Title: Valid Sudoku

Difficulty: 2/10

Date Solved: 3/13/2026


PATTERN
------------------------------
Pattern / Category: HashMap

Key Signals That Indicate This Pattern:
- Need to check whether duplicates exist in each row, column, square


CORE INSIGHT
------------------------------
Nothing much problem was trivial.


SOLUTION SUMMARY
------------------------------
Create a hashmap for every row, column, square and check if duplicates ever happen.


COMPLEXITY
------------------------------
Variables:
- n: number of elements in the row

Time Complexity: O(n^2) 

Space Complexity: O(n^2)

Explanation: The entire grid is visited no more than 3 times and n^2 items are stored stored 3 times.


MISTAKE / LESSON
------------------------------
What initially confused me:
Nothing. 


What should I recognize faster next time:
Nothing.


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems ChatGPT Examples:
- Database constraint validation: Ensuring unique values in rows or columns (similar to SQL UNIQUE constraints).

- Distributed system request deduplication: Hash sets used to detect repeated message IDs.

- Input validation pipelines: Ensuring unique identifiers within a batch of records before processing.