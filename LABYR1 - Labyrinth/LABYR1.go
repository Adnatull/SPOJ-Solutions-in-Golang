package main

import "fmt"


var mlen int
var row, column int
var n1, n2 int

func dfs (i int, j int, graph[] string, distance[][] int, parent int)  { 
	if i<0 || i>= row || j<0 || j >= column || graph[i][j]=='#' || distance[i][j]!=0 { 
			return
	}	 
	distance[i][j] = parent + 1	
	if distance[i][j] > mlen { 
		mlen = distance[i][j]
		n1, n2 = i, j
	}
	dfs(i-1,j, graph, distance, distance[i][j])
	dfs(i, j-1, graph, distance, distance[i][j])
	dfs(i, j+1, graph, distance, distance[i][j])
	dfs(i+1, j, graph, distance, distance[i][j])
	return
} 

func main() { 
	var t int
	fmt.Scanf("%d", &t)
	for cases := 1; cases <= t; cases ++ { 
		
		fmt.Scanf("%d %d", &column, &row)
		var graph[] string
		var s string
		for i := 1; i<= row; i++ { 
			fmt.Scanln(&s)
			graph = append(graph, s)
		}
				

		var fond bool
		fond = false
		for i := 0; i<row; i++ { 
			for j := 0; j<column; j++ { 
				if graph[i][j] == '.' { 
					mlen = 0
					distance := make([][] int, row+1)
					for k := 0; k<=row; k++ { 
						distance[k] = make([] int, column+1)
					}
					dfs(i, j, graph, distance, -1)
					mlen = 0
					for k := 0; k<=row; k++ { 
						for l := 0; l<= column; l++ { 
							distance[k][l] = 0
						}
					}
					var p, q int
					p, q = n1, n2
					dfs(p, q, graph, distance, -1)
					fmt.Printf("Maximum rope length is %d.\n",mlen)
					
					fond = true
					break

				}
			}
			if fond { 
				break
			}
		}
		
	}
	
}
