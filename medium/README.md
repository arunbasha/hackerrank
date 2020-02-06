# HackerRank Medium Challenges

## 1. Caesar Cipher
Perform a Caesar Cipher shift on it using the num parameter as the shifting number. A Caesar Cipher works by shifting 
each letter in the string N places down in the alphabet (in this case N will be num). Punctuation, spaces, and 
capitalization should remain intact.

For example:

    if the string is "Caesar Cipher" and num is 2 the output should be "Ecguct Ekrjgt"
    if the string is "abc!@# z" and num is 1 the output should be "bcd!@# a"
    if the string is "X^& yZ" and num is 1 the output should be "Y^& zA"

## 2. Camel Case Worlds
Alice wrote a sequence of words in CamelCase as a string of letters, s, having the following properties:

It is a concatenation of one or more words consisting of English letters.
All letters in the first word are lowercase.
For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.
Given s, print the number of words in s on a new line.

For example, s = oneTowThree. There are 3 words in the string.

Complete the camelcase function in the editor below. It must return the integer number of words in the input string.

## 3. New Year Chaos
It's New Year's Day and everyone's in line for the Wonderland rollercoaster ride! There are a number of people queued 
up, and each person wears a sticker indicating their initial position in the queue. Initial positions increment by 1 
from 1 at the front of the line to n at the back.

Any person in the queue can bribe the person directly in front of them to swap positions. If two people swap positions, 
they still wear the same sticker denoting their original places in line. One person can bribe at most two others. 
For example, if n = 8 and Person 5 bribes Person 4, the queue will look like this: 1,2,3,4,5,6,7,8.

Fascinated by this chaotic queue, you decide you must know the minimum number of bribes that took place to get the queue
into its current state!

## 4. Minimum Swaps
You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates. You are allowed to swap any two elements. You need to find the minimum number of swaps required to sort the array in ascending order.

For example, given the array arr == [7,1,3,2,4,5,6] we perform the following steps:

```
    i   arr                         swap (indices)
    0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
    1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
    2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
    3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
    4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
    5   [1, 2, 3, 4, 5, 6, 7]
```
It took 5 swaps to sort the array.

Function Description

Complete the function minimumSwaps in the editor below. It must return an integer representing the minimum number of 
swaps to sort the array.

