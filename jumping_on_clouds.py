#!/usr/bin/python3

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c):
    # c = [0,0,1,0,0,1,0]

    jumps = 0
    current_position = 0
    while current_position < len(c) - 1:
        #                                      |
        # protect from going out of bounds ie. v
        #                             [0,0,0,1,0,0]
        if current_position + 2 <= len(c) - 1 and c[current_position + 2] == 0:
            current_position += 2
        else:
            current_position += 1
        jumps += 1
    print(jumps)
    return jumps


jumpingOnClouds([0, 1, 0, 0, 0, 1, 0])
jumpingOnClouds([0, 0, 1, 0, 0, 1, 0])
jumpingOnClouds([0, 0, 0, 0, 1, 0])
