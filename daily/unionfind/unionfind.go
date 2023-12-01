package unionfind

//type UnionFind struct {
//	parents []int
//	sizes   []int
//}
//
//func NewUnionFind(n int) *UnionFind {
//	parents, sizes := make([]int, n), make([]int, n)
//	for i := 0; i < n; i++ {
//		parents[i], sizes[i] = i, 1
//	}
//	return &UnionFind{
//		parents: parents,
//		sizes:   sizes,
//	}
//}
//
//func (uf *UnionFind) Find(x int) int {
//	if uf.parents[x] == x {
//		return x
//	}
//	uf.parents[x] = uf.Find(uf.parents[x])
//	return uf.parents[x]
//}
//
//func (uf *UnionFind) Union(x, y int) {
//	rx, ry := uf.Find(x), uf.Find(y)
//	if rx != ry {
//		if uf.sizes[rx] > uf.sizes[ry] {
//			uf.parents[ry], uf.sizes[rx] = rx, uf.sizes[rx]+uf.sizes[ry]
//		} else {
//			uf.parents[rx], uf.sizes[ry] = ry, uf.sizes[ry]+uf.sizes[rx]
//		}
//	}
//}
//
//func (uf *UnionFind) GetSize(x int) int {
//	return uf.sizes[x]
//}

/*
CountPairs2316
2316. Count Unreachable Pairs of Nodes in an Undirected Graph
*/
//func CountPairs2316(n int, edges [][]int) int64 {
//	uf := NewUnionFind(n)
//	for _, edge := range edges {
//		uf.Union(edge[0], edge[1])
//	}
//
//	ans := int64(0)
//	for i := 0; i < n; i++ {
//		ans += int64(n - uf.GetSize(uf.Find(i)))
//	}
//
//	return ans / 2
//}

/*
765. Couples Holding Hands
*/
type UnionFind struct {
	parents []int
	rank    []int
	size    int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		rank[i] = 1
	}
	return &UnionFind{parents: parents, rank: rank, size: n}
}

// Find root of x
func (uf *UnionFind) Find(x int) int {
	if x == uf.parents[x] {
		return x
	}
	uf.parents[x] = uf.Find(uf.parents[x])
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.size--
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parents[rootY] = rootX
		uf.rank[rootX] += uf.rank[rootY]
	} else {
		uf.parents[rootX] = rootY
		uf.rank[rootY] += uf.rank[rootX]
	}
}

func minSwapsCouples(row []int) int {
	uf := NewUnionFind(len(row))
	for i := 0; i < len(row); i += 2 {
		uf.Union(i, i+1)
		uf.Union(row[i], row[i+1])
	}
	return len(row)/2 - uf.size
}
