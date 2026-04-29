func twoSum(nums []int, target int) []int {
    solve  := []int{}
    found := false
    for ind1, value1 := range nums {
        for ind2, value2 :=  range nums {
            if value1 + value2 == target  && ind1 != ind2{
                solve = append(solve, ind1, ind2)
                found = true
                break
            }
        }
        if(found) {
            break
        }
    }
    return solve
}