class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const result = strs.reduce((acc, curr) => {
            const chars = curr.split("").sort().toString();
            if (acc[chars]) {
                acc[chars].push(curr)
            } else {
                acc[chars] = [curr];
            }

            return acc;
        }, {});

        return Object.values(result)
    }
}
