import re

count = 0
with open('input.txt') as file_in:
    for line in file_in:
        split_line = line.rstrip().split(': ')
        
        rules = re.split('-| ', split_line[0])
        first = int(rules[0])
        second = int(rules[1])
        character = rules[2]
        
        password = split_line[1]
        ch_first = password[first - 1]
        ch_second = password[second - 1]

        # wow this if statement is kind of gross -- fix this later
        if ((ch_first == character and ch_second != character)
            or (ch_first != character and ch_second == character)):
            count += 1

print(count)
