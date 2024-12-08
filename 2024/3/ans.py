from pathlib import Path
import re

input = open(Path(__file__).with_name('input.txt'), 'r').read()

# part 1
pattern = re.compile('mul\(\d+,\d+\)')
matches = re.findall(pattern, input)

ans1 = sum(map(lambda x: int(x[0])*int(x[1]), map(lambda x: x[4:-1].split(','),
                                                  matches)))
print(ans1)

# part 2
total = 0
pattern = re.compile('mul\(\d+,\d+\)|do\(\)|don\'t\(\)')
matches = re.findall(pattern, input)

do = True
for match in matches:
    if match == 'do()':
        do = True
    elif match == 'don\'t()':
        do = False
    else:
        if do:
            x, y = match[4:-1].split(',')
            total += int(x)*int(y)
print(total)
