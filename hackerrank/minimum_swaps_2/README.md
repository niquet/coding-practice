# Minimum Swaps 2

## Problem

You are given an unordered array consisting of consecutive integers $[1, 2, 3, ..., n]$ without any duplicates. You are allowed to swap any two elements. Find the minimum number of swaps required to sort the array in ascending order.

## Constraints

- $1 \leq n \leq 10^5$
- $1 \leq arr[i] \leq n$

## Idea #1

- Unordered array consisting of **consecutive integers** _without any duplicates_
- Loop over the unordered array
    - If the integer at index $i$ is not at the correct position, swap integer at $i$ with integer at the correct position
    - Increment swap counter by one
    - If integer at index $i$ is at the correct position, move pointer to the next bigger index
