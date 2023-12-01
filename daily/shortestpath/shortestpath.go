package shortestpath

import "math"

/*
1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance
*/
func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
	// w[i][j] 记录 i 到 j 的路径长度weight
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)
		for j := range w {
			w[i][j] = math.MaxInt / 2 // 防止加法溢出
		}
	}
	// 无边的weight为无穷大
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		w[x][y], w[y][x] = wt, wt
	}

	// Floyd算法
	f := w
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}
	// 此时f[i][j]表示从i到j的最短路径长度

	maxCnt := n
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i != j && f[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= maxCnt {
			maxCnt = cnt
			ans = i
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
399. Evaluate Division
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) (res []float64) {
	mp := make(map[string]int)
	e := [40][40]float64{} // edges
	n := 0
	for i, equation := range equations {
		for _, x := range equation {
			if _, ok := mp[x]; !ok {
				mp[x] = n
				n++
			}
		}
		u, v, value := mp[equation[0]], mp[equation[1]], values[i]
		e[u][v], e[v][u] = value, 1.0/value
	}
	// n indicate number of points

	f := e // complete graph
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if f[i][k] != 0 && f[k][j] != 0 {
					f[i][j] = f[i][k] * f[k][j]
				}
			}
		}
	}

	for _, querie := range queries {
		x, ok1 := mp[querie[0]]
		y, ok2 := mp[querie[1]]
		if !(ok1 && ok2) || f[x][y] == 0 {
			res = append(res, -1)
		} else {
			res = append(res, f[x][y])
		}
	}

	return res
}
