// Copyright 2020 Juan Tellez All rights reserved.

package critpath

/*
   From LeetCode Critical Path detector

   https://leetcode.com/problems/critical-connections-in-a-network/

   Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
   Output: [[1,3]]
   Explanation: [[3,1]] is also accepted.

   This is a brute force attack.  The idea is to iterate over all the nodes
   each time, while marking a single node as expendable.  This solution works
   but has O(v*e*v), Leetcode website rejects it with time exceeded
*/

func dfs(graph [][]int, visited []bool, node int) int {

	if visited[node] == true {
		return 0
	}

	visited[node] = true
	nodes := 1
	neighbors := graph[node]
	for _, n := range neighbors {
		nodes = nodes + dfs(graph, visited, n)
	}

	return nodes
}

func bfCriticalConnections(n int, connections [][]int) [][]int {

	critical := make([][]int, 0)
	for i := range critical {
		critical[i] = make([]int, 0)
	}

	for exclude := range connections {
		// make a graph of neighbors for each node
		graph := make([][]int, n)
		for i := range graph {
			graph[i] = make([]int, 0)
		}

		visited := make([]bool, n)

		for node := range graph {
			graph[node] = make([]int, 0)
			for c := range connections {
				if c == exclude {
					continue
				}
				if connections[c][0] == node {
					graph[node] = append(graph[node], connections[c][1])
				}
				if connections[c][1] == node {
					graph[node] = append(graph[node], connections[c][0])
				}
			}
		}

		for i := range visited {
			visited[i] = false
		}
		count := dfs(graph, visited, 0)
		if count != n {
			critical = append(critical, connections[exclude])
		}
	}

	return critical
}
