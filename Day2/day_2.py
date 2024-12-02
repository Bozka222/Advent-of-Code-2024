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

with open("input_data.txt") as data:
    lines = data.read().strip().split("\n")

for line in lines:
    nums = [int(i) for i in line.split()]
    ans += is_safe(nums)

print(ans)
