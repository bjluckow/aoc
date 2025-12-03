import re

def read_lines():
    with open("input.txt", 'r') as f:
        for line in f:
            yield line.strip()

input = next(read_lines())
# example input
#input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

ranges = input.split(',')
ranges = map(lambda x: x.split('-'), ranges)
ranges = map(lambda x: (int(x[0]), int(x[1])), ranges)
ranges = map(lambda x: range(x[0], 1+x[1]), ranges) # 1+x[1] for inclusive ranges
ranges = list(ranges)

pattern = re.compile(r"^(\d+)\1{1}$") # use {1} instead of {2} since backref \1 already counts first instance 

total = sum([i if pattern.match(str(i)) else 0 for r in ranges for i in r])
print(total)

# part 2

pattern = re.compile(r"^(\d+)\1+$")
total = sum([i if pattern.match(str(i)) else 0 for r in ranges for i in r])
print(total)

