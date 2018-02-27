def twoSum(nums, target):
    d = {}
    for index in range(len(nums)):
        v = target - nums[index]
        if v in d:
            return [d[v], index]
        d[nums[index]] = index


print(twoSum([2, 7, 11, 15], 9))
