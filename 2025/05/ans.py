ranges: list[tuple[int, int]] = list()
items: list[int] = list()

with open("input.txt", 'r') as f:
    is_reading_ranges = True
    for line in f:
        line = line.strip()
        if line == "":
            is_reading_ranges = False
        elif is_reading_ranges:
            low, high, *rest = line.split('-')
            ranges.append((int(low), int(high)))
        else:
            items.append(int(line))


num_ok = 0
for item in items:
    num_ok += any(map(lambda r: r[0]<=item<=r[1],ranges))

print(num_ok)

# part 2
# would love to use set |= but not today...
sorted_ranges = sorted(ranges, key=lambda r: r[0])
merged_ranges: list[tuple[int, int]] = list()

for sr in sorted_ranges:
    if not merged_ranges or merged_ranges[-1][1] < sr[0]:
        merged_ranges.append(sr)
    else:
        merged_ranges[-1] = (merged_ranges[-1][0], max(merged_ranges[-1][1], sr[1]))

total = sum(map(lambda r: r[1]-r[0]+1, merged_ranges))
print(total)
