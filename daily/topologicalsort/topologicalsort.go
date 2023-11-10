package topologicalsort

/*
MaximumInvitations2127
2127. Maximum Employees to Be Invited to a Meeting
*/
func MaximumInvitations2127(favorite []int) int {
	// i -> favorite[i]
	n := len(favorite)
	indeg := make([]int, n) // 记录每个结点的入度
	for _, x := range favorite {
		indeg[x]++
	}

	used := make([]bool, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1 // 到节点 i 为（包括i）止的最长「游走」路径经过的节点个数
	}

	// 拓扑排序
	q := []int{} // 队列中只保存入度为0的结点，表示可以遍历，且遍历过后，它指向的点的入度减1
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		used[u] = true
		q = q[1:]
		v := favorite[u] // u -> v
		// 状态转移
		f[v] = max(f[v], f[u]+1)
		indeg[v]--
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}

	// ring 表示最大的环的大小
	// total 表示所有环大小为 2 的 【基环内向树】 上的最长的 【双向游走】路径之和
	ring, total := 0, 0
	for i := 0; i < n; i++ {
		if !used[i] { // i 在环内
			j := favorite[i]
			if favorite[j] == i { // i, j 所在的环的大小为2
				total += f[i] + f[j]
				used[i], used[j] = true, true
			} else {
				// i ->j -> ...
				u, cnt := i, 0
				for {
					cnt++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = max(ring, cnt)
			}
		}
	}
	return max(ring, total)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
