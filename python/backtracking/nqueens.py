#!/usr/bin/python3

import typing

# Returns all valid ways to place n queens on an n X n chess board such that no queen can attack another
class NQueens:
    def __init__(self, n: int):
        self.solutions = []
        self.state = []
        """
        A 1D array of O(n) size can be used to check if a space is safe for a queen in O(1)
        I keep track of n columns, and 2+(n-2)*4 diagonals. 2+(n-2)*4 is the number of diagonals in an nXn chessboard that go through 2 or more cells
        A 0 means that column or diagonal is safe for a queen and a 1 means unsafe.
        My diagonals are in order from bottom-left to top-right, then top-left to bottom-right
        e.g. for n=2 safe_spaces=[0, 1, 1, 0] safe_spaces[1] says that column 1 is unsafe and safe_spaces[2] says that diagonal 1 unsafe
             | |Q|
             | | |      
        """
        num_spaces = n + 2 + (n - 2) * 4
        self.safe_spaces = [0] * num_spaces
        self.n = n
    
    def is_solution(self) -> bool:
        return len(self.state) == self.n
    
    # Returns two indices into self.safe_spaces that represent the 2 diagonals going through a particular cell in an nXn chessboard
    # A return value of -1 just means there's no relevant diagonal e.g. at (0,0) the diagonal from bottom left to top right doesn't matter since it only goes through one cell
    def get_diagonal_indices(self, choice) -> tuple[int, int]:
        row, col = len(self.state), choice
        # the diagonal going from bottom left to top right
        steps = min(self.n - row, col)
        d_row, d_col = row+steps, col-steps
        diagonal_index_1 = d_row - 1 + d_col
        offset = self.n
        if diagonal_index_1 >= 0:
            diagonal_index_1 += offset
        # the diagonal going from top left to bottom right
        steps = min(row, col)
        d_row, d_col = (row-steps, col-steps)
        offset = self.n + 1 + (2 * (self.n-2))
        diagonal_index_2 = self.n - 2 - d_row + d_col
        if diagonal_index_2 >= 0:
            diagonal_index_2 += offset
        return diagonal_index_1, diagonal_index_2
    
    def is_valid(self, choice: int) -> bool:
        if choice < 0 or choice >= self.n:
            return False
        diagonal_index_1, diagonal_index_2 = self.get_diagonal_indices(choice)
        return self.safe_spaces[choice] == 0 and (diagonal_index_1 == -1 or self.safe_spaces[diagonal_index_1] == 0) and (diagonal_index_2 == -1 or self.safe_spaces[diagonal_index_2] == 0)
    
    def add_queen(self, choice: int):
        diagonal_index_1, diagonal_index_2 = self.get_diagonal_indices(choice)
        self.safe_spaces[choice] = 1
        if diagonal_index_1 != -1:
            self.safe_spaces[diagonal_index_1] = 1
        if diagonal_index_2 != -1:
            self.safe_spaces[diagonal_index_2] = 1
        self.state.append(choice)
        
    def remove_queen(self, choice: int):
        self.state.pop() # I pop first because get_diagonal_indices depends on self.state
        diagonal_index_1, diagonal_index_2 = self.get_diagonal_indices(choice)
        self.safe_spaces[choice] = 0
        self.safe_spaces[diagonal_index_1] = 0
        self.safe_spaces[diagonal_index_2] = 0
        
    def solve_nqueens(self) -> list[list[int]]:
        if (self.n == 1):
            return [[0]]
        if self.is_solution():
            self.solutions.append(self.state[:]) # add a copy of the state to final result list
            return
        for choice in range(0, self.n):
            if self.is_valid(choice):
                self.add_queen(choice) # make move
                self.solve_nqueens()
                self.remove_queen(choice) # backtrack
        return self.solutions

if __name__ == "__main__":
    solver = NQueens(4)
    print(solver.solve_nqueens())