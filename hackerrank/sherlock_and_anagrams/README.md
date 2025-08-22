# Sherlock and Anagrams

## Definitions

Two strings are _anagrams_ of each other if the letters of one string can be rearranged to form the other string.

## Problem

Given a string $s$, find the number of pairs of substrings of the string that are anagrams of each other.

### Input format

The first line contains an integer $q$, the number of queries.

Each of the next $q$ lines contains a string $s$ to analyze.

### Constraints

- $1 \leq q \leq 10$

- $2 \leq length of s \leq 100$

- $s$ containts only lowercase letters in the range ascii (a - z)

## Example 0

```bash
# Input
2
abba
abcd

# Output
4
0
```

## Example 1

```bash
# Input
2
ifailuhkqq
kkkk

# Output
3
10
```

## Example 2

```bash
# Input
1
cdcd

# Output
5
```
