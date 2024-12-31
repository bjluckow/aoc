from pathlib import Path

# example = """
# ....#.....
# .........#
# ..........
# ..#.......
# .......#..
# ..........
# .#..^.....
# ........#.
# #.........
# ......#...
# """
# input = [list(line) for line in example.split()]

input = []
with open(Path(__file__).with_name('input.txt'), 'r') as file:
    for line in file:
        input.append(list(line[:-1]))

EMPTY = '.'
BLOCKED = '#'
START = '^'

# up, right, down, left
DIRS = ((0, -1), (1, 0), (0, 1), (-1, 0))

empty_set = set()
blocked_set = set()
start_pos = (None, None)


def get_spot(pos):
    x, y = pos
    return input[y][x]


for y in range(len(input)):
    for x in range(len(input[y])):
        pos = (x, y)
        spot = get_spot(pos)
        if spot == EMPTY:
            empty_set.add(pos)
        elif spot == BLOCKED:
            blocked_set.add(pos)
        elif spot == START:
            start_pos = pos
        else:
            raise Exception('unknown char')


reachable_set = set()
current_dir_idx = 0
current_pos = start_pos


def is_in_bounds(pos):
    x, y = pos
    return (0 <= x < len(input)) and (0 <= y < len(input[0]))


def move(pos, dir_idx):
    dir = DIRS[dir_idx]
    # print(f"moved {dir}")
    return (pos[0] + dir[0], pos[1] + dir[1])


def rotate(dir_idx):
    return (dir_idx + 1) % 4


while is_in_bounds(current_pos):
    reachable_set.add(current_pos)
    next_pos = move(current_pos, current_dir_idx)
    while is_in_bounds(next_pos) and get_spot(next_pos) == BLOCKED:
        current_dir_idx = rotate(current_dir_idx)
        next_pos = move(current_pos, current_dir_idx)
    current_pos = next_pos

print(len(reachable_set))

# part 2


def detect_cycle(extra_block: tuple[int, int]) -> bool:
    seen = set()
    current_pos = start_pos
    current_dir = 0

    def step(pos, dir_idx):
        next = move(pos, dir_idx)
        while is_in_bounds(next) and (get_spot(next) == BLOCKED or next == extra_block):
            dir_idx = rotate(dir_idx)
            next = move(pos, dir_idx)
        return next, dir_idx

    while is_in_bounds(current_pos):
        if (current_pos, current_dir) in seen:
            return True
        seen.add((current_pos, current_dir))
        current_pos, current_dir = step(current_pos, current_dir)

    return False


cycle_positions = set()
trial_num = 1  # just for logging
for pos in empty_set:
    print(f"trying extra block position {trial_num}/{len(empty_set)}")
    trial_num += 1
    if detect_cycle(pos):
        cycle_positions.add(pos)

print(len(cycle_positions))
