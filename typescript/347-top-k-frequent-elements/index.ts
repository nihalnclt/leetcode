console.log("-- 347. Top K Frequent Elements --");

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

function topKFrequent(nums: number[], k: number): number[] {
    let map = new Map<number, number>();

    for (let num of nums) {
        map.set(num, (map.get(num) || 0) + 1);
    }

    return Array.from(map.entries())
        .sort((a, b) => {
            return b[1] - a[1];
        })
        .slice(0, k)
        .map((arr) => {
            return arr[0];
        });
}

console.log(topKFrequent([2, 3, 2, 1, 4, 1, 1], 2));
