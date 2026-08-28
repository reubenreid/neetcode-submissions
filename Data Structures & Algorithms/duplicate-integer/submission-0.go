import "slices"

func hasDuplicate(nums []int) bool {
    seen := make([]int, 0)

    for i := 0; i < len(nums); i++ {
	    if (!slices.Contains(seen, nums[i])) {
           seen = append(seen, nums[i])
        }

        if (len(seen) == len(nums)) {
            break;
        }
    }

    return len(nums) != len(seen)
}
