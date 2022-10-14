package leetcode

func maxProfit(prices []int) int {
    b,profit := 0,0
    for i, p := range prices {
        if p <= prices[b] {
            b = i
        } else {
            profit = max(profit, p - prices[b])
        }
    }
    return profit
}

func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}
