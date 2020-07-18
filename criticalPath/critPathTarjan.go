// Copyright 2020 Juan Tellez All rights reserved.

package critpath

/*
   From LeetCode Critical Path detector

   https://leetcode.com/problems/critical-connections-in-a-network/

   Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
   Output: [[1,3]]
   Explanation: [[3,1]] is also accepted.

   This uses Tarjan's algorith.  The idea is to iterate over all the nodes
   each time, while marking each node as we execute a DFS recursion.  As we
   reach each node it is marked with the lowest visible node seen.  This
   has the effect of detecting cycles in a graph.  Once the dfs recursion
   returns, we compare the returned id and if it doesn't match our lowest
   node visible it means that there is no cycle. seen.


*/

var (
	lowTimes []int
	time     int
	critical [][]int
)

// tdfs implements Tarjan's algorithm using depths first search
func tdfs(graph [][]int, visited []int, node int, parent int) {

	if visited[node] != -1 {
		return
	}

	min := func(a int, b int) int {
		switch {
		case a < b:
			return a
		case b < a:
			return b
		default:
			return a
		}
	}

	time = time + 1
	visited[node] = time
	lowTimes[node] = time

	neighbors := graph[node]
	for _, nbr := range neighbors {
		if nbr == parent {
			continue
		}
		if visited[nbr] == -1 {
			tdfs(graph, visited, nbr, node)
			lowTimes[node] = min(lowTimes[node], lowTimes[nbr])
			if visited[node] < lowTimes[nbr] {
				critical = append(critical, []int{node, nbr})
			}
		} else {
			lowTimes[node] = min(lowTimes[node], visited[nbr])
		}
	}
}

func criticalConnections(n int, connections [][]int) [][]int {

	critical = make([][]int, 0)
	for i := range critical {
		critical[i] = make([]int, 0)
	}

	lowTimes = make([]int, n)

	// make a graph of neighbors for each node
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	// make neighbors graph
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}

	visited := make([]int, n)
	for i := range visited {
		visited[i] = -1
	}
	tdfs(graph, visited, 0, -1)

	return critical
}
