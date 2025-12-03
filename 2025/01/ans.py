input_path = "./input.txt"

def read_lines(path=input_path):
    with open(path, 'r') as f:
        for line in f:
            yield line.strip()

pos = 50
max_pos = 100
counter = 0

for line in read_lines():
    dir, *turns = line
    sgn = 1 if dir == 'R' else -1
    turns = sgn * int("".join(turns))
    pos = (pos + turns) % max_pos
    counter += int(pos == 0)

print(counter)
