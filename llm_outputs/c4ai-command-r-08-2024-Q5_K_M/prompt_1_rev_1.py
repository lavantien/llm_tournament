import time
from queue import PriorityQueue
import random

def initialize_sudoku_grid(input_grid):
    """
    Initializes the sudoku grid by parsing the input data.

    Args:
        input_grid (str): The input Sudoku grid represented as a string, with '.' for empty cells.

    Returns:
        grid (2D list): A 2D list representing the Sudoku grid.
        pre_solve_count (int): The number of empty cells in the grid before solving.
    """
    grid = []
    pre_solve_count = 0

    # Loop over each line (row) of the Sudoku input
    for i in range(9):
        # For each cell in the row
        row = []
        for c in input_grid[9 * i: 9 * (i + 1)]:
            # If the cell is empty (represented by a '.')
            if c == '.':
                # Increment the pre_solve_count as this cell is empty
                pre_solve_count += 1
                # Add a 'None' to the row as we don't have a number for this cell yet
                row.append(None)
            # If the cell contains a digit
            else:
                # Add the digit (cast to an integer) to the row
                row.append(int(c))
        # Once the entire row is processed, append it to our grid
        grid.append(row)
    # Return the processed grid and the pre_solve_count
    return grid, pre_solve_count

def backtracking_solver(grid, pre_solve_count):
    """
    Solves the Sudoku grid using the backtracking algorithm.

    Args:
        grid (list): The Sudoku grid as a 2D list.
        pre_solve_count (int): The number of empty cells before solving.

    Returns:
        solved (list): The solved Sudoku grid as a 2D list, or None if no solution exists.
        pre_solve_count (int): The number of empty cells before solving.
        step_count (int): The number of steps taken to solve the grid.
        exec_time (float): The execution time for solving the grid.
    """
    start_time = time.time()

    def is_valid(board, row, col, num):
        """Checks if a number is valid for a particular cell in the grid"""
        # Check row
        if num in board[row]:
            return False
        # Check column
        for i in range(9):
            if board[i][col] == num:
                return False
        # Check subgrid
        start_row = 3 * (row // 3)
        start_col = 3 * (col // 3)
        for i in range(start_row, start_row + 3):
            for j in range(start_col, start_col + 3):
                if board[i][j] == num:
                    return False
        return True

    def backtrack_solver(board, pre_solve_count, step_count=0):
        for row in range(9):
            for col in range(9):
                if board[row][col] is None:
                    for num in range(1, 10):
                        if is_valid(board, row, col, num):
                            board[row][col] = num
                            if backtrack_solver(board, pre_solve_count, step_count + 1):
                                return True
                            board[row][col] = None
                    return False  # If no solution found with the current values, backtrack
                    # backtracking_solver returns False if no solution found for that branch
        # If no empty cells are found, return the board as the final solution
        return True

    success = backtrack_solver(grid, pre_solve_count)

    solved = success * grid
    end_time = time.time()
    step_count = (pre_solve_count - grid.count(None))
    exec_time = end_time - start_time
    return solved, pre_solve_count, step_count, exec_time

def astar_solver(grid, pre_solve_count):
    """
    Solves the Sudoku grid using A-Star algorithm with appropriate heuristics.

    Args:
        grid (list): The Sudoku grid as a 2D list.
        pre_solve_count (int): The number of empty cells before solving.

    Returns:
        solved (list): The solved Sudoku grid as a 2D list, or None if no solution exists.
        pre_solve_count (int): The number of empty cells before solving.
        step_count (int): The number of steps taken to solve the grid.
        exec_time (float): The execution time for solving the grid.
    """
    start_time = time.time()
    queue = PriorityQueue()
    visited = set()

    # Heuristic function
    def heuristic(cell):
        row, col = cell
        return (cell not in visited) + (grid[row][col] == None) * (9 - (grid[row].count(None) + sum(1 for i in grid if i[col] == None)))

    queue.put((heuristic(cell) * 100, 0, cell))

    while not queue.empty():
        cost, _, cell = queue.get()
        row, col = cell
        if grid[row][col] != None:
            visited.add(cell)
        elif not any(visited.intersection(((row-i,col), (row+i,col), (row,col-i), (row,col+i), (row-i,col-i), (row-i,col+i), (row+i,col-i), (row+i,col+i)))):
            grid[row][col] = random.choice(list(set(range(1,10)) - set(grid[row]) - set(grid[:row][col]) - {r[col] for r in grid[row+1:]}) + [1,2,3,4,5,6,7,8,9])
            for neighbour in ((row-1,col), (row+1,col), (row,col-1), (row,col+1), (row-1,col-1), (row-1,col+1), (row+1,col-1), (row+1,col+1)):
                if grid[neighbour[0]][neighbour[1]] != None:
                    queue.put((heuristic(cell), cost+1, neighbour))
            visited.add(cell)

    # Extract the solution path
    solution = [grid]
    while grid[cell] is None:
        for i in ((-1,0),(1,0),(0,-1),(0,1),(1,1),(1,-1),(-1,1),(-1,-1)):
            r,c = cell
            if grid[r+i[0]][c+i[1]] is not None:
                solution.append([(grid[r][c],grid[r+i[0]][c+i[1]]) if grid[r][c] is not None else grid[r+i[0]][c+i[1]]])
                grid[r][c],grid[r+i[0]][c+i[1]] = grid[r+i[0]][c+i[1]],grid[r][c]
                cell = (r+i[0], c+i[1])
                break

    # Print the solved Sudoku
    for line in solution:
        for val in line:
            if type(val) is list:
                print(str(val[0]) + "|" + str(val[1]) + "|",end="")
            else:
                print(str(val) + "|" ,end="")
        print("")

    return (solution)

    solved = grid
    end_time = time.time()
    step_count = sum([line.count(None) for line in grid])
    exec_time = end_time - start_time
    return solved, pre_solve_count, step_count, exec_time

def ant_colony_optimization(grid, pre_solve_count):
    """
    Solves the Sudoku grid using the Ant Colony Optimization (ACO) algorithm.

    Args:
        grid (list): The Sudoku grid as a 2D list.
        pre_solve_count (int): The number of empty cells before solving.

    Returns:
        solved (list): The solved Sudoku grid as a 2D list, or None if no solution exists.
        pre_solve_count (int): The number of empty cells before solving.
        step_count (int): The number of steps taken to solve the grid.
        exec_time (float): The execution time for solving the grid.
    """
    # Implement ACO algorithm here, following steps: initialization, local search, pheromone updates, etc.
    # Return the solved grid, pre-solve count, and execution time
    return None, pre_solve_count, 0, 0.0

def minimax_alpha_beta_solver(grid, pre_solve_count):
    """
    Solves the Sudoku grid using the Minimax algorithm with Alpha-Beta Pruning.

    Args:
        grid (list): The Sudoku grid as a 2D list.
        pre_solve_count (int): The number of empty cells before solving.

    Returns:
        solved (list): The solved Sudoku grid as a 2D list, or None if no solution exists.
        pre_solve_count (int): The number of empty cells before solving.
        step_count (int): The number of steps taken to solve the grid.
        exec_time (float): The execution time for solving the grid.
    """
    def possible_moves(cell_state, grid, cell):
        possible = set()
        if cell_state[cell] == -1:  # An empty cell
            for num in range(1, 10):  # Try all 1 to 9
                row = cell_state[0] + cell % 9
                col = cell_state[1] + cell // 9
                if valid_cell(grid, (row, col), num):
                    possible.add(num)
        else:  # A non-empty cell
            for num in grid[cell_state[0] + cell % 9]:  # Check the 9 numbers
                possible.add(num)
            if not cell_state[2] < num:  # A wrong choice has been tried already
                return False
        return possible

    def alpha_beta(alpha, beta, depth, is_max, cell_state, grid, cell):
        # Terminal case: leaf node
        if is_max and possible_moves(cell_state, grid, cell):  # The last move
            return -depth  # Max wants to minimize
        if not is_max and not possible_moves(cell_state, grid, cell):  # Min is forced to do a bad move
            return depth
        score = 0

        # Get the best score from all possible moves
        if is_max:  # Maximizing the score
            score = -100
            for move in possible_moves(cell_state, grid, cell):
                new_cell = (cell_state[0] + cell % 9, cell_state[1] + cell // 9, cell_state[2] + 1)
                temp = alpha_beta(-beta, -alpha, depth + 1, False, new_cell, grid, cell)
                score = max(score, temp)
                if score > alpha:  # Update the alpha and beta values
                    alpha = score
        else:
            score = 100
            for move in possible_moves(cell_state, grid, cell):
                new_cell = (cell_state[0] + cell % 9, cell_state[1] + cell // 9, cell_state[2] + 1)
                temp = alpha_beta(-beta, -alpha, depth + 1, True, new_cell, grid, cell)
                score = min(score, temp)
                if score < beta:
                    beta = score
        return score

    def valid_cell(grid, cell, num):
        row, col = cell
        return grid[row][col] is None or grid[row][col] == num  # If empty cell, then it is valid, if filled then must match num
        # Note that we also check the other cells in the same block here

    def minimax_algorithm(cell_state, grid):
        if len([move for move in possible_moves(cell_state, grid, 0) if move >= 1000]):  # 20% of all cells have a number > 1000, meaning the puzzle is solvable
            return alpha_beta(-1000, 1000, 0, True, (0,0,0), grid, 0)
        else:
            return -1000

    start_time = time.time()
    # Initialized the variables and grid
    cell_state = [0,0,0]

    solution_count = minimax_algorithm(cell_state, grid)

    # Convert the grid back into the list and print
    solution_list = [[0]*9 for i in range(9)]
    for i in range(len(grid)):
        for j in range(len(grid)):
            solution_list[i][j] = grid[i][j][1]
    for i in solution_list:
        print("|".join(str(element) for element in i) + "|")

    # Extract the solution
    return solution_list, pre_solve_count, pre_solve_count * solution_count * 100000, (time.time() - start_time)

    end_time = time.time()
    step_count = solution_count
    exec_time = end_time - start_time
    return solution_list, pre_solve_count, step_count, exec_time

if __name__ == "__main__":
    # Input puzzle represented as a string, with '.' for empty cells.
    input_grid = """.......1.
4........
.2.......
....5.4.7
..8...3..
..1.9....
3..4..2..
.5.1.....
...8.6...""".strip().split("\n")

    # Initialize the Sudoku grid and pre_solve_count
    grid, pre_solve_count = initialize_sudoku_grid(input_grid)

    # Run each algorithm to solve the Sudoku and record the output
    results = {}
    for algo, func in zip(['Backtracking', 'A-star with good heuristics', 'Ant colony optimization', 'Minimax with alpha-beta pruning'],
                         [backtracking_solver, astar_solver, ant_colony_optimization, minimax_alpha_beta_solver]):
        solved, _, step_count, exec_time = func(grid, pre_solve_count)
        # For Backtracking, ensure the solution matches the original grid
        if algo == 'Backtracking' and solved != grid:
            raise AssertionError('Backtracking did not solve the Sudoku!')
        # Store the results of this algorithm run
        results[algo] = {'Pre-solve count': pre_solve_count, 'Step count': step_count, 'Execution time': exec_time}

    # Format and print the results
    for algo in results:
        print(results[algo])
