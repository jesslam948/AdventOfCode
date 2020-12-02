nums = []
with open('input.txt') as file_in:
    for line in file_in:
        num = int(line)
        nums.append(num)
nums.sort()

i = 0
j = len(nums) - 1

while i < j:
    my_sum = nums[i] + nums[j]
    if my_sum == 2020:
        print(str(nums[i]) + " " + str(nums[j]))
        print(str(nums[i] * nums[j]))
        break
    elif my_sum < 2020:
        i += 1
    else:
        j -= 1

