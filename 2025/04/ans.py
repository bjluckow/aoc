from typing import Optional

with open("input.txt", 'r') as f:
    grid = [list(line.strip()) for line in f]
max_row, max_col = len(grid), len(grid[0])

adj_vecs = [
        (-1, -1), (0, -1), ( 1, -1),
        (-1,  0),          ( 1,  0),
        (-1,  1), (0,  1), ( 1,  1),
        ]


def add_vecs(a: tuple[int, int], b: tuple[int, int]) -> tuple[int, int]:
    return (a[0]+b[0], a[1]+b[1])

def is_in_bounds(pos: tuple[int, int]) -> bool:
    if not pos:
        return False
    row, col = pos
    return 0 <= row < max_row and 0 <= col < max_col 

def get_pos(pos: tuple[int, int]) -> Optional[str]:
    if is_in_bounds(pos):
        row, col = pos 
        return grid[row][col]
    else:
        return None

def get_adj(pos: tuple[int, int]) -> list[str]:
    return list(filter(lambda x: x, map(lambda x: get_pos(add_vecs(pos, x)), adj_vecs)))

def is_removable(pos: tuple[int, int]) -> bool:
    return get_pos((r,c)) == '@' and sum(map(lambda x: x == '@', get_adj((r,c)))) < 4

count = 0
for r in range(max_row):
    for c in range(max_col):
        if is_removable((r, c)): 
            count += 1

print(count)            

# part 2
count = 0
removed = list()
while True:
    for r in range(max_row):
        for c in range(max_col):
            if is_removable((r, c)):
                count += 1
                removed.append((r,c))
    if not removed:
        break
    
    for (r, c) in removed:
        grid[r][c] = '.'

    removed = list()
    

print(count)

