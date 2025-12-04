def read_lines():
    with open("input.txt", 'r') as f:
        for line in f:
            yield line.strip()

input = list(read_lines())

def maximize_joltage(s: str) -> int:
    nums = list(map(lambda x: int(x), s))
    max_i, max_total = nums[0], 0

    for idx, i in enumerate(nums[:-1]):
        if i < max_i:
            continue

        max_i = i

        for jdx, j in enumerate(nums[idx+1:]):
            max_total = max(max_total, 10*i + j)

    return max_total

total = sum(map(maximize_joltage, input))
print(total)

# part 2 

def maximize_joltage_2(batteries: list[int], order_of_magnitude: int) -> int:
    if order_of_magnitude == 0:
        return max(batteries)

    max_num = max(batteries[:-order_of_magnitude])
    num_idx = batteries.index(max_num) 
    
    return max_num * 10**order_of_magnitude + maximize_joltage_2(batteries[num_idx+1:], order_of_magnitude-1)

total = 0

for line in input:
    chars = list(line)
    nums = list(map(lambda x: int(x), chars))
    total += maximize_joltage_2(nums, 12-1)
    
print(total)
