# Count Commas in Range — 3870

https://leetcode.com/problems/count-commas-in-range/

## Problem

You are given an integer `n`.

Return the **total** number of commas used when writing all integers from
`[1, n]` (inclusive) in **standard** number formatting.

In **standard** formatting:

- A comma is inserted after **every three** digits from the right.
- Numbers with **fewer** than 4 digits contain no commas.

**Example 1:**

```
Input:  n = 1002
Output: 3
```

The numbers `"1,000"`, `"1,001"`, and `"1,002"` each contain one comma,
giving a total of 3.

**Example 2:**

```
Input:  n = 998
Output: 0
```

All numbers from 1 to 998 have fewer than four digits, therefore no commas
are used.

**Constraints:** `1 <= n <= 10^5`

## Approach

A number `x` is written with **at least k commas** iff `x >= 10^(3k)` —
every additional comma requires three more leading digits (10^3 = 1,000;
10^6 = 1,000,000; ...).

So the answer is the sum, over each threshold `10^(3k) <= n`, of how many
numbers in `[1, n]` reach that threshold:

```
total = Σ over k >= 1 of  max(0, n - 10^(3k) + 1)
```

`n <= 10^5` has at most two thresholds, giving **O(log n)** time and
**O(1)** space.

## Tests

- All problem examples plus boundary checks (`1000 -> 1`, `100000 -> 99001`).
- Brute-force cross-check for `n = 1..10000` that literally formats each
  number with thousands separators and counts the commas.