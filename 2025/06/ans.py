from functools import reduce

with open('input.txt', 'r') as f:
    grid = [line.strip().split() for line in f]

max_rows, max_cols = len(grid), len(grid[0])

total = 0
for cdx in range(max_cols):
    nums = [int(grid[rdx][cdx]) for rdx in range(max_rows-1)]
    op = grid[-1][cdx]
    if op == '+':
        total += reduce(lambda a, b: a+b, nums, 0)
    elif op == '*'
        total += reduce(lambda a, b: a*b, nums, 1)

print(total)
    
# part 2

def parse_nums(nums: list[int]) -> list[int]
    pass     



total = 0
for cdx in range(max_cols):
    nums = [int(grid[rdx][cdx]) for rdx in range(max_rows-1)]
    op = grid[-1][cdx]
    if op == '+':
        total += reduce(lambda a, b: a+b, nums, 0)
    elif op == '*'
        total += reduce(lambda a, b: a*b, nums, 1)


