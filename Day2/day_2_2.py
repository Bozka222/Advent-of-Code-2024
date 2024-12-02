ans = 0

def is_safe(nums):
    inc = nums[1] > nums[0]
    if inc:
        for i in range(1, len(nums)):
            diff = nums[i] - nums[i - 1]
            if not 1 <= diff <= 3:
                return False
        return True
    else:
        for i in range(1, len(nums)):
            diff = nums[i] - nums[i - 1]
            if not -3 <= diff <= -1:
                return False
        return True

def is_really_safe(nums):
    if is_safe(nums):
        return True
    for i in range(len(nums)):
        if is_safe(nums[:i] + nums[i + 1:]):
            return True
    return False

with open("input_data.txt") as data:
    lines = data.read().strip().split("\n")

for line in lines:
    nums = [int(i) for i in line.split()]
    ans += is_really_safe(nums)

print(ans)
