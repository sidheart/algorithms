#!/usr/bin/python3
import typing

# O(N) safety check, can be reduced to O(1) using O(N) space by storing n columns and O(N) diagonals
def _can_place_queen(col: int, state: list[int]):
  row = len(state)
  for j in range(0, len(state)):
    queen_row = j
    queen_col = state[j]
    slope = (queen_col - col) / (queen_row - row)
    if slope == 0 or slope == 1 or slope == -1:
      return False
  return True

# O(N!) backtracking 
def _solve_n_queens(n: int, state: list[int], answer: list[int]):
  if n <= 0:
    return
  if len(state) == n:
    answer.append(state)
    return
  for col in range(0, n):
    if _can_place_queen(col, state, n):
      state.append(col)
      _solve_n_queens(n, state, answer)
      state.pop()

# returns all valid ways to place n queens on an nXn chess board such that no queen can attack another
# the output is a list of column-wise placements 
# e.g. [2, 0, 3, 1] says queens may be safely placed at (0,2), (1,0), (2,3), and (3,1)
def solve_n_queens(n: int) -> list[list[int]]:
  # Write your code here
  answer = []
  state = []
  _solve_n_queens(n, state, answer)
  return answer