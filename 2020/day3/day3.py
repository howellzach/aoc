# Day 3 Challenge

with open("day3.txt", "r") as f:
    slope = [line.rstrip('\n') for line in f]
    
maxsl = len(slope[0])-1

def sled(right, down):
    trees = 0
    m_r = 0
    m_d = 0
    
    for i in range(len(slope)):
        m_r = m_r + right
        m_d = m_d + down
        try:
            if m_r > maxsl:
                m_r = m_r - maxsl-1
                if slope[m_d][m_r] == "#":
                    trees +=1
            else:
                if slope[m_d][m_r] == "#":
                    trees +=1     
        except IndexError:
            pass
    return trees

pt1 = sled(3,1)

r1d1 = sled(1,1)
r3d1 = sled(3,1)
r5d1 = sled(5,1)
r7d1 = sled(7,1)
r1d2 = sled(1,2)

pt2 = r1d1 * r3d1 * r5d1 * r7d1 * r1d2

print("Part 1: " + str(pt1))
print("Part 2: " + str(pt2))