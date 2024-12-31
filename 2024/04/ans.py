from pathlib import Path

# 140x141
input = [list(line[:-1]) for line in list(
    open(Path(__file__).with_name('input.txt'), 'r'))]


def get_char(line: int, char: int) -> str:
    if 0 <= line < len(input) and 0 <= char < len(input[line]):
        return input[line][char]
    return ''


def get_word(line: int, char: int, x_offset: int, y_offset: int, length: int) -> str:
    return ''.join([get_char(line+(idx*y_offset), char+(idx*x_offset)) for idx in range(length)])


"""
(-1, -1) ( 0, -1) (+1, -1)
(-1,  0) ( 0,  0) ( 1,  0)
(-1, +1) ( 0, +1) (+1, +1)
"""


def north(line, char):
    return get_word(line, char, 0, -1, 4)


def east(line, char):
    return get_word(line, char, 1, 0, 4)


def south(line, char):
    return get_word(line, char, 0, 1, 4)


def west(line, char):
    return get_word(line, char, -1, 0, 4)


def northeast(line, char):
    return get_word(line, char, -1, 1, 4)


def northwest(line, char):
    return get_word(line, char, -1, -1, 4)


def southeast(line, char):
    return get_word(line, char, 1, 1, 4)


def southwest(line, char):
    return get_word(line, char, 1, -1, 4)


def count_xmas(line_idx, char_idx):
    return sum(map(lambda x: x == "XMAS", [north(line_idx, char_idx),
                                           east(line_idx, char_idx),
                                           west(line_idx, char_idx),
                                           south(line_idx, char_idx),
                                           northeast(line_idx, char_idx),
                                           northwest(line_idx, char_idx),
                                           southeast(line_idx, char_idx),
                                           southwest(line_idx, char_idx)]))


total = sum([count_xmas(line_idx, char_idx)
             for line_idx in range(len(input))
             for char_idx in range(len(input[line_idx]))
             ])


print(total)


# part 2


def is_x_mas(line, char):
    left = ''.join([get_char(line-1, char-1),
                   get_char(line, char),
                   get_char(line+1, char+1)])

    right = ''.join([get_char(line-1, char+1),
                    get_char(line, char),
                    get_char(line+1, char-1)])

    return (left == "MAS" or left == "SAM") and (right == "MAS" or right == "SAM")


total = sum([is_x_mas(line_idx, char_idx)
             for line_idx in range(len(input))
             for char_idx in range(len(input[line_idx]))
             ])
print(total)
