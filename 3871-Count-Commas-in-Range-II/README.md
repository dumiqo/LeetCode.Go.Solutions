# Count Commas in Range II — 3871

https://leetcode.com/problems/count-commas-in-range-ii/

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

**Constraints:** `1 <= n <= 10^15`

## Approach

A number `x` is written with **at least k commas** iff `x >= 10^(3k)` —
every additional comma requires three more leading digits (10^3 = 1,000;
10^6 = 1,000,000; ...).

So the answer is the sum, over each threshold `10^(3k) <= n`, of how many
numbers in `[1, n]` reach that threshold:

```
total = Σ over k >= 1 of  max(0, n - 10^(3k) + 1)
```

`n <= 10^15` has at most five thresholds (10^3, 10^6, 10^9, 10^12, 10^15),
giving **O(log n)** time and **O(1)** space. The maximum answer (~4 * 10^15)
fits comfortably in `int64`, so no modulo is needed; an overflow guard
(`pow > n/1000` before `pow *= 1000`) keeps the multiplication in range.

## Tests

- All problem examples plus boundary clauses, including crossing the second,
  third and fifth comma thresholds (`10^6 -> 999002`, `10^9 -> 1998999003`,
  `10^15 -> 3998998998999005`).
- Brute-force cross-check for `n = 1..10^6` against an independent
  digit-length oracle that counts a number's commas as `(digits-1)/3`
  (no shared threshold logic), accumulated incrementally so the test stays
  linear in `n`.