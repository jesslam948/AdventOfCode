nums = []
with open('input.txt') as file_in:
    for line in file_in:
        num = int(line)
        nums.append(num)
nums.sort()

for i in range(len(nums)):
    low = i + 1
    high = len(nums) - 1

    while low < high:
        two_sum = nums[low] + nums[high]
        three_sum = two_sum + nums[i]
        if three_sum == 2020:
            print(str(nums[i]) + " " + str(nums[low]) + " " +
                str(nums[high]))
            print(str(nums[i] * nums[low] * nums[high]))
            break
        elif three_sum < 2020:
            low += 1
        else:
            high -= 1

