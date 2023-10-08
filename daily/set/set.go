package set

import "github.com/emirpasic/gods/trees/redblacktree"

/*
StockPrice
2034. 股票价格波动
*/
type StockPrice struct {
	prices       *redblacktree.Tree // price: times	某价格出现的次数
	timePriceMap map[int]int        // time: price	某时间的股票价格
	maxTimestamp int                // 当前时间
}

func Constructor() StockPrice {
	return StockPrice{prices: redblacktree.NewWithIntComparator(), timePriceMap: map[int]int{}, maxTimestamp: 0}
}

func (this *StockPrice) Update(timestamp int, price int) {
	// 两种情况
	// 1. 新增时间 map直接放 tree直接put(price, times+1)
	// 2. 对已有时间的更正(删掉再添加) 需在tree中去除原有时间对应的价格 若次数多次则减一 一次则remove 最后map再放新时间

	if prePrice := this.timePriceMap[timestamp]; prePrice > 0 {
		// 已有当前时间 需更正价格: 去除原有价格prePrice
		if times, _ := this.prices.Get(prePrice); times.(int) > 1 {
			this.prices.Put(prePrice, times.(int)-1)
		} else {
			this.prices.Remove(prePrice)
		}
	}
	times := 0
	if val, ok := this.prices.Get(price); ok {
		times = val.(int)
	}
	this.prices.Put(price, times+1)
	this.timePriceMap[timestamp] = price

	// 更新最新时间
	if timestamp >= this.maxTimestamp {
		this.maxTimestamp = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.timePriceMap[this.maxTimestamp]
}

func (this *StockPrice) Maximum() int {
	return this.prices.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.prices.Left().Key.(int)
}
