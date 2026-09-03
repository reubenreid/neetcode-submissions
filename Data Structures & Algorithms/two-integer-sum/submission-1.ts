class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums: number[], target: number): number[] {
        const seen = {};
        for(let i = 0; i < nums.length; i++) {
            if (seen[nums[i]] !== undefined) {
                return [seen[nums[i]], i];
            }

            const aim = target - nums[i];
            seen[aim] = i;
        }
    }
}
