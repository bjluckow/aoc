from pathlib import Path
import math

input = []
with open(Path(__file__).with_name('input.txt'), 'r') as file:
    for line in file:
        total, operands = line.split(':')
        input.append(
            (int(total), list(map(lambda x: int(x), operands.strip().split(' ')))))


def can_produce(total: int, operands: list[int]) -> bool:
    if len(operands) == 0:
        return False
    if len(operands) == 1:
        return total == operands[0]
    if len(operands) == 2:
        return ((operands[0] + operands[1] == total) or (operands[0] * operands[1] == total))

    return can_produce(total - operands[-1], operands[:-1]) or \
        can_produce(total / operands[-1], operands[:-1])


result = sum(list(map(lambda x: x[0], filter(
    lambda x: can_produce(x[0], x[1]), input))))
print(result)

# part 2


def concat(a, b) -> int:
    # return int(''.join([str(a), str(b)]))
    return a*(10**int(1+(math.log(b, 10)))) + b


def unconcat(a, b) -> int:
    return (a-b)/(10**int(1+(math.log(b, 10))))


def can_produce_2(total: int, operands: list[int]) -> bool:
    if len(operands) == 0:
        return False
    if len(operands) == 1:
        return total == operands[0]
    if len(operands) == 2:
        return ((operands[0] + operands[1] == total) or
                (operands[0] * operands[1] == total) or
                (concat(operands[0], operands[1]) == total))

    return can_produce_2(total - operands[-1], operands[:-1]) or \
        can_produce_2(total / operands[-1], operands[:-1]) or \
        can_produce_2(unconcat(total, operands[-1]), operands[:-1])


result_2 = sum(list(map(lambda x: x[0], filter(
    lambda x: can_produce_2(x[0], x[1]), input))))
print(result_2)
