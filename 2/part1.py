import re

count = 0
with open('input.txt') as file_in:
    for line in file_in:
        split_line = line.rstrip().split(': ')
        
        rules = re.split('-| ', split_line[0])
        occ_min = int(rules[0])
        occ_max = int(rules[1])
        character = rules[2]
        
        password = split_line[1]
        char_count = password.count(character)

        if char_count >= occ_min and char_count <= occ_max:
            count += 1

print(count)
