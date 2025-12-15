from functools import reduce

with open('input.txt', 'r') as f:
    grid = [line.strip().split(" ") for line in f]

print(grid[0])
