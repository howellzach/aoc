with open("input.txt") as f:
    lines = f.read().splitlines()

nums = []
for i in lines:
    nums.append(int(i))

def count_increases(input):
    count = 0
    for i, v in enumerate(input[1:]):
        value = input[i]
        nextvalue = input[i+1]
        if nextvalue > value:
            count = count + 1
    return count

def make_groups(input):
    c = 0
    groups = []
    try:
        for i in range(len(input)):
            group = input[c] + input[c+1] + input[c+2]
            c += 1
            groups.append(group)
    except IndexError:
        return groups
    
    return groups

groups = make_groups(nums)
count = count_increases(groups)

print(count)