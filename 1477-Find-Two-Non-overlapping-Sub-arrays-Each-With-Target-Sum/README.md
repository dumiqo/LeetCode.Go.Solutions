# 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum

**Difficulty:** Medium  
**Topics:** Array, Hash Table, Dynamic Programming, Sliding Window, Binary Search  

You are given an array of integers `arr` and an integer `target`.

You have to find **two non-overlapping sub-arrays** of `arr` each with a sum equal `target`. There can be multiple answers so you have to find an answer where the sum of the lengths of the two sub-arrays is **minimum**.

Return _the minimum sum of the lengths_ of the two required sub-arrays, or return `-1` if you cannot find such two sub-arrays.

## Examples

**Example 1:**

```
Input: arr = [3,2,2,4,3], target = 3
Output: 2
Explanation: Only two sub-arrays have sum = 3 ([3] and [3]). The sum of their lengths is 2.
```

**Example 2:**

```
Input: arr = [7,3,4,7], target = 7
Output: 2
Explanation: Although we have three non-overlapping sub-arrays of sum = 7 ([7], [3,4] and [7]), but we will choose the first and third sub-arrays as the sum of their lengths is 2.
```

**Example 3:**

```
Input: arr = [4,3,2,6,2,3,4], target = 6
Output: -1
Explanation: We have only one sub-array of sum = 6.
```

## Constraints

- `1 <= arr.length <= 10^5`
- `1 <= arr[i] <= 1000`
- `1 <= target <= 10^8`

## Solution

**Algorithm: Hash Table + Prefix Sum + Dynamic Programming**

A hash table maps each prefix sum to its most recent 1-based position.  
`f[i]` stores the minimum length of a target-sum sub-array found within the first `i` elements.

When a target-sum sub-array ends at position `i` (1-based), its start was `j = d[s-target]`, so its length is `i-j`.  
Combining it with the best sub-array fully to its left (`f[j]`) gives a candidate answer `f[j] + (i - j)`.  
Since `j < i`, the two sub-arrays can never overlap.

**Time complexity:** O(n)  
**Space complexity:** O(n)