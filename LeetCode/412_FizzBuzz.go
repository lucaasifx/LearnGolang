func fizzBuzz(n int) []string {
    result := make([]string, n)
    step := ""

    for x := 1; x <= n; x++ {
        if x % 3 == 0 && x % 5 == 0 {
            step = "FizzBuzz"
        } else if x % 3 == 0 {
            step = "Fizz"
        } else if x % 5 == 0 {
            step = "Buzz"
        } else {
            step = fmt.Sprintf("%d", x)
        }
        result[x - 1] = step
    }

    return result
}