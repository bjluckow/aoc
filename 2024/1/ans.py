from pathlib import Path

list1 = []
list2 = []
with open(Path(__file__).with_name('input.txt'), 'r') as file:
    for line in file:
        entry1, entry2 = line.split()
        list1.append(entry1)
        list2.append(entry2)

if len(list1) != len(list2):
    print('different list lengths')

list1 = list(map(lambda x: int(x), list1))
list2 = list(map(lambda x: int(x), list2))

# Part 1

ans1 = sum(map(lambda x,
               : abs(x[0]-x[1]), zip(sorted(list1), sorted(list2))), 0)
print(ans1)  # Correct

# Part 2
list1_set = set(list1)
count_map = dict()
for entry in list2:
    if entry in list1_set:
        count_map[entry] = count_map.get(entry, 0) + 1

ans2 = sum(list(map(lambda x: x * count_map.get(x, 0), list1_set)))
print(ans2)  # Correct
