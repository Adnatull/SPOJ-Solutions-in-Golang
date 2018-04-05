package main	

import "fmt"

var cnt int
var visit[] bool
var ans bool

func dfs(graph [][]int, source , parent int) { 	
	cnt++
	visit[source] = true
	for j:=0; j<len(graph[source]); j++ { 
		var v int
		v = graph[source][j]
		if visit[v] == false  {
			dfs(graph, v, source)
		} else if v != parent { 
			ans = false
		}
	}	
}

func main() {
	var n, e int
	fmt.Scanf("%d %d", &n, &e)
	
	graph := make([][]int, n+1)
	ans = true
	var x, y int
	for i:= 0; i<e; i++ { 
		fmt.Scanf("%d %d", &x, &y)
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	cnt = 0
	for i:=0; i<=n; i++ { 
		visit = append(visit, false)
	}
	dfs(graph,1, -1)
	if cnt == n && ans { 
		fmt.Printf("YES\n")
	} else { 
		fmt.Printf("NO\n")
	}
}
