func containsDuplicate(nums []int) bool {
    mapping := make(map[int]int)
    for _, num := range nums{
        _, exists := mapping[num]
        if exists{
            return true
        }
        mapping[num] = 1
    }
    return false
}