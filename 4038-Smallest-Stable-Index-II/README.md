# Smallest Stable Index II — 4038

https://leetcode.com/problems/smallest-stable-index-ii/

## Problem

You are given an integer array `nums` of length `n` and an integer `k`.

For each index `i`, define its **instability score** as

```
max(nums[0..i]) - min(nums[i..n-1])
```

where:

- `max(nums[0..i])` is the largest value among elements from index 0 to
  index `i`.
- `min(nums[i..n-1])` is the smallest value among elements from index `i` to
  index `n-1`.

An index `i` is called **STABLE** if its instability score is less than or
equal to `k`.

Return the **smallest** stable index. If no such index exists, return `-1`.

**Example 1:**

```
Input:  nums = [5,0,1,4], k = 3
Output: 3
```

| i | max(nums[0..i]) | min(nums[i..n-1]) | score |
|---|------------------|--------------------|-------|
| 0 | 5                | 0                  | 5     |
| 1 | 5                | 0                  | 5     |
| 2 | 5                | 1                  | 4     |
| 3 | 5                | 4                  | 1     |

Index `3` has score `1 <= 3`, so the answer is `3`.

**Example 2:**

```
Input:  nums = [3,2,1], k = 1
Output: -1
```

Every index has score `3 - 1 = 2 > 1`, so no index is stable.

**Example 3:**

```
Input:  nums = [0], k = 0
Output: 0
```

The single index has score `0 - 0 = 0 <= 0`.

**Constraints:**

- `1 <= nums.length <= 10^5`
- `0 <= nums[i] <= 10^9`
- `0 <= k <= 10^9`

## Approach

Compute two auxiliary arrays in a single pass each:

- `prefMax[i]` = the maximum of `nums[0..i]`.
- `sufMin[i]` = the minimum of `nums[i..n-1]`.

Then scan `i` from `0` to `n-1`, returning the first index where

```
prefMax[i] - sufMin[i] <= k
```

If no index satisfies the condition, return `-1`.

This runs in **O(n)** time and **O(n)** space.

## Tests

- All three problem examples.
- Edge cases: single-element arrays and `k = 0` with strictly
  increasing/decreasing input (where the answer is `0` or `-1` depending on
  whether the last element equals the first/largest prefix value).
- A brute-force cross-check over random arrays, where the oracle recomputes
  `max(nums[0..i])` and `min(nums[i..n-1])` independently for each `i`.