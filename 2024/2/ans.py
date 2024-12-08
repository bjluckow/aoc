from pathlib import Path

examples = [[7, 6, 4, 2, 1],
            [1, 2, 7, 8, 9],
            [9, 7, 6, 2, 1],
            [1, 3, 2, 4, 5],
            [8, 6, 4, 4, 1],
            [1, 3, 6, 7, 9]]

input: list[list[int]] = []
with open(Path(__file__).with_name('input.txt'), 'r') as file:
    for line in file:
        input.append([int(x) for x in line.split()])

# part 1


def is_safe(levels: list[int]) -> bool:
    # len(levels)-1 pairs of (levels[i], levels[i+1]) for i..len(levels)-1
    consec = list(zip(levels, levels[1:]))
    all_asc = all(map(lambda x: x[0] < x[1], consec))
    all_desc = all(map(lambda x: x[0] > x[1], consec))
    all_in_range = all(map(lambda x: 1 <= abs(x[1]-x[0]) <= 3, consec))
    return all_in_range and (all_asc or all_desc)


ans1 = sum(map(lambda x: is_safe(x), input))
print(ans1)  # correct

# part 2


def is_safe_with_tolerance(levels: list[int]) -> bool:
    if is_safe(levels):
        return True

    for idx in range(len(levels)):
        if is_safe(levels[:idx]+levels[idx+1:]):
            return True
    return False


ans2 = sum(map(lambda x: is_safe_with_tolerance(x), input))
print(ans2)  # Correct
