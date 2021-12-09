# Day 1 Challenge
import sys

with open("day1.txt", "r") as f:
    list = [line.rstrip('\n') for line in f]
expenses = [ int(x) for x in list ]

def puzzle1():
    for i in range(len(expenses)):
        x = expenses[i]
        for i in range(len(expenses)):
            z = expenses[i]
            y = x + z
            answer = x * z
            if y == 2020:
                print(x)
                print(z)
                return ("THE ANSWER FOR PART 1 IS... " + str(answer))

def puzzle2():
    for i in range(len(expenses)):
        a = expenses[i]
        for i in range(len(expenses)):
            x = expenses[i]
            for i in range(len(expenses)):
                z = expenses[i]
                y = x + z + a
                answer = x * z * a
                if y == 2020:
                    print(x)
                    print(z)
                    print(a)
                    return ("THE ANSWER FOR PART 2 IS... " + str(answer))


print(puzzle1())
print(puzzle2())