
import random

length = 100

targetSize = 100
target = []


def generateBigNum():
    return random.randint(0, length**3)


left = targetSize
for num in range(length):
    if generateBigNum() % (length-num) < left:
        target.append(num)
        left -= 1

print(target)
