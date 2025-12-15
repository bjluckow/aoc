def lines():
    with open("input.txt", 'r') as f:
        for line in f:
            yield line

SOURCE, SPLIT = 'S', '^'

beams = set([next(lines()).index(SOURCE)])
count = 0

for line in lines():
    to_add = set()
    to_del = set()
    for beam in beams:
        if line[beam] == SPLIT:
            count += 1
            to_del.add(beam)
            if beam-1 >= 0:
                to_add.add(beam-1)
            if beam+1 < len(line):
                to_add.add(beam+1)
    beams -= to_del
    beams |= to_add

    # just for aesthetics
    line = list(line)
    for b in to_add:
        line[b] = '|'
    print("".join(line))

print(count) 
