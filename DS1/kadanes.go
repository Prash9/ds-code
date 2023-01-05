import (
    "math"
)
func maxSubArray(nums []int) int {
    max_sum := math.MinInt
    current_sum := 0
    max_element := math.MinInt
    for _, val := range nums{
        current_sum += val
        if current_sum < 0 {
            current_sum = 0
        }
        if current_sum > max_sum{
            max_sum = current_sum
        }
        // max_element = math.Max(max_element, val)
        max_element = int(math.Max(float64(max_element), float64(val)))

    }
    if max_sum == 0 {
        return max_element
    }
    return max_sum
}