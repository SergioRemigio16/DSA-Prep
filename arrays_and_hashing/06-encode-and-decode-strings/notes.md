PROBLEM
------------------------------
Title: Encode and Decode Strings
Difficulty: 1/10
Date Solved: 3/10/2026


PATTERN
------------------------------
Pattern / Category: Delimiter and Length Prefix Parsing

Key Signals That Indicate This Pattern:
- Messages need to be compressed into one string then decoded


CORE INSIGHT
------------------------------
Delimiter may appear in message

SOLUTION SUMMARY
------------------------------
Convert messages into this string format:

(int: message length)#(string: message)(int: message length)#(string: message)!

Where # is a delimiter and ! marks the end of the encoded string


COMPLEXITY
------------------------------
Variables:
- m : sum of length of all strings
- n : number of all strings

Time Complexity: O(m)

Space Complexity: O(m + n)

Explanation: 

Decoding and encoding requires algorithm to look at every character in every string at most twice.

It can take O(m) space to store every character of every string but if there are n  empty messages then it would still take O(n) to store all that information. 


MISTAKE / LESSON
------------------------------
What initially confused me: That space complexity is O(m + n) because the strings might be empty.


What should I recognize faster next time: That empty strings still take up space if there are n of them.


SYSTEMS CONNECTION
------------------------------
Where this concept appears in real systems ChatGPT Examples:
- Network protocols: Many protocols use length-prefixed messages to frame packets.
- Database storage formats: Variable-length fields are stored using length prefixes.
- Serialization frameworks: Formats such as Protocol Buffers or Thrift use similar mechanisms to define message boundaries.