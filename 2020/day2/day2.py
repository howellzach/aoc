# Day 2 Challenge

with open("day2.txt", "r") as f:
    data = [line.rstrip('\n') for line in f]

pwdb = [ i.split(" ") for i in data ]
policy = []

for i in range(len(pwdb)):
    policy.append(pwdb[i][0].split("-"))

policy = [[int(x) for x in lst] for lst in policy]

count1 = 0

for i in range(len(pwdb)):
    letter = pwdb[i][1][0]
    pwd = pwdb[i][2].count(letter)
    if pwd in range(policy[i][0], policy[i][1]+1):
        count1 += 1
    else:
        count1 + 0

valid = 0
invalid = 0

for i in range(len(pwdb)):
    letter = pwdb[i][1][0]
    pwd = pwdb[i][2]
    pos1 = policy[i][0] - 1
    pos2 = policy[i][1] - 1

    if letter == pwd[pos1] and letter != pwd[pos2]:
        valid += 1
    elif letter != pwd[pos1] and letter == pwd[pos2]:
        valid += 1
    elif letter == pwd[pos1] and letter == pwd[pos2]:
        invalid += 1
    elif letter != pwd[pos1] and letter != pwd[pos2]:
        invalid += 1

combined = valid + invalid

print("The answer to part 1: " + str(count1))
print("The answer to part 2: " + str(valid))
print("Part 2 invalid count: " + str(invalid))
print("Part 2 combined counts: " + str(combined))