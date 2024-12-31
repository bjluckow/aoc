from pathlib import Path
from functools import cmp_to_key


rules = []
updates = []

with open(Path(__file__).with_name('input.txt'), 'r') as file:
    # read rule pairs
    for line in file:
        if line == "\n":
            break
        n1, n2 = line.split('|')
        rules.append((int(n1), int(n2)))

    # read update lists
    for line in file:
        updates.append(list(map(lambda x: int(x), line.split(','))))

rules_set = set(rules)  # tuples are hashable


def compare(x: int, y: int) -> int:
    if (x, y) in rules_set:
        return -1
    if (y, x) in rules_set:
        return 1
    return 0


correct_updates = []
corrected_incorrect_updates = []
for update_list in updates:
    correct_list = sorted(update_list, key=cmp_to_key(compare))
    if update_list == correct_list:
        correct_updates.append(update_list)
    else:
        corrected_incorrect_updates.append(correct_list)

correct_middle_nums = list(map(lambda x: x[len(x)//2], correct_updates))
print(sum(correct_middle_nums))

# part 2
corrected_incorrect_middle_nums = list(
    map(lambda x: x[len(x)//2], corrected_incorrect_updates))
print(sum(corrected_incorrect_middle_nums))
