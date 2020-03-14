# Author: Hang Yu
# Create Time: 13/3/2020
import sys

# Check whether the characters in s can be replaced to get t
def can_map(s: str, t: str):
    # empty string and length mismatch always return false
    if not s or not t or len(s) != len(t): 
        return False
    mapping = {}
    for char_s, char_t in zip(s, t):
        # return False if this mapping conflicts with former one
        if char_s in mapping:
            if char_t != mapping[char_s]:
                return False
        # create new mapping if this mapping is a new one
        else:
            mapping[char_s] = char_t
    return True


if __name__ == "__main__":
    # Given less than two string
    if len(sys.argv) < 3:
        print("false")
    else:
        s, t = sys.argv[1], sys.argv[2]
        if can_map(s, t):
            print("true")
        else:
            print("false")

    