# Count Commas in Range — 3683

https://leetcode.com/problems/count-commas-in-range/

## Problem

Given an integer `n`, return the **total** number of commas used when writing
all integers from `[1, n]` (inclusive) in **standard** number formatting:

- a comma is inserted after every **three** digits from the right
- numbers with fewer than 4 digits contain no commas

**Example 1:** `n = 1002` → `3`
("1,000", "1,001", "1,002" — each has one comma)

**Example 2:** `n = 998` → `0`
(all numbers have fewer than four digits)

**Constraints:** `1 <= n <= 10^5`

## Approach

A number with `d` digits is written with `floor((d-1)/3)` commas.

For each digit length `d`, the `d`-digit numbers not exceeding `n` form the
interval `[10^(d-1), min(n, 10^d - 1)]`, and each contributes `(d-1)/3` commas:

```
total = Σ over digit lengths d of  max(0, min(n, 10^d - 1) - 10^(d-1) + 1) * floor((d-1)/3)
```

`n <= 10^5` spans at most 6 digit lengths → **O(log n)** time, **O(1)** space.