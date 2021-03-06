# HackerRank Easy Challenges

## 1. Sock Merchant 
John works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of 
integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
For example, there are n = 7 socks with colors ar = [1,2,1,2,1,3,2]. There is one pair of color 1 and one of color 2. 
There are three odd socks left, one of each color. The number of pairs is 2.

Complete the sockMerchant function in the editor below. It must return an integer representing the number of matching pairs of socks that are available.

## 2. Valley Counting

Gary is an avid hiker. He tracks his hikes meticulously, paying close attention to small details like topography. During
his last hike he took exactly n steps. For every step he took, he noted if it was an uphill, U, or a downhill, D step. 
Gary's hikes start and end at sea level and each step up or down represents 1 unit change in altitude. We define the 
following terms:

A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a 
step down to sea level.

A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a 
step up to sea level.

Given Gary's sequence of up and down steps during his last hike, find and print the number of valleys he walked through.

For example, if Gary's path is s = [DDUUUUDD], he first enters a valley 2 units deep. Then he climbs out an up onto a 
mountain 2 units high. Finally, he returns to sea level and ends his hike.

Complete the countingValleys function in the editor below. It must return an integer that denotes the number of valleys Gary traversed.

## 3. Jumping On Clouds
Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads 
and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current 
cloud plus 1 or 2. She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from
her starting position to the last cloud. It is always possible to win the game.

For each game, Emma will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided. 
For example, c = [ 0,1,0,0,0,1,0] indexed from 0...6. The number on each cloud is its index in the list so she must 
avoid the clouds at indexes 1 and 5. She could follow the following two paths:  0->2->4->6 or 0->2->3->4->6. The first 
path takes 3 jumps while the second takes 4.

Complete the jumpingOnClouds function in the editor below. It should return the minimum number of jumps required, as an 
integer.

## 4. Repeated String
Lilah has a string, s, of lowercase English letters that she repeated infinitely many times.

Given an integer, n, find and print the number of letter a's in the first  letters of Lilah's infinite string.

For example, if the string s='abcac' and n = 10, the substring we consider is abcacabcac, the first 10 characters of her
infinite string. There are 4 occurrences of a in the substring.

Complete the repeatedString function in the editor below. It should return an integer representing the number of 
occurrences of a in the prefix of length n in the infinitely repeating string.

## 5. Hourglass Sum
Given a 6x6 2D Array, arr:
```

    1 1 1 0 0 0
    0 1 0 0 0 0
    1 1 1 0 0 0
    0 0 0 0 0 0
    0 0 0 0 0 0
    0 0 0 0 0 0
```
We define an hourglass in A to be a subset of values with indices falling in this pattern in arr's graphical representation:
```
    a b c
      d
    e f g
```
There are 16 hourglasses in arr, and an hourglass sum is the sum of an hourglass' values. Calculate the hourglass sum 
for every hourglass in arr, then print the maximum hourglass sum.

For example, given the 2D array:

```
    -9 -9 -9  1 1 1 
     0 -9  0  4 3 2
    -9 -9 -9  1 2 3
     0  0  8  6 6 0
     0  0  0 -2 0 0
     0  0  1  2 4 0
```
We calculate the following  hourglass values:

```
    -63, -34, -9, 12, 
    -10, 0, 28, 23, 
    -27, -11, -2, 10, 
    9, 17, 25, 18
```
Our highest hourglass value is 28 from the hourglass:

```
    0 4 3
      1
    8 6 6
```

## 6. Left Rotation
A left rotation operation on an array shifts each of the array's elements 1 unit to the left. For example, if 2 left 
rotations are performed on array [1,2,3,4,5] , then the array would become [3,4,5,1,2].

Given an array a of n integers and a number, d, perform d left rotations on the array. Return the updated array to be 
printed as a single line of space-separated integers.

Complete the function rotLeft in the editor below. It should return the resulting array of integers.
