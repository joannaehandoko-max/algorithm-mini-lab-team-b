# Algorithm Mini Lab - Team B

## Team Members
| Student | GitHub username |
|---|---|
| Student A: _Satria Danish Khan_ | @_sluxis-spec_ |
| Student B: _Najiib Rahmawan_ | @_najiibrahmawan2008-blip!_ |
| Student C: _Joana Evangeline Handoko_ | @_joannaehandoko-max_ |

## Our Experiments
1. **Insertion Sort** (`experiments/insertion_sort/main.go`)
2. **Depth-First Search (DFS)** (`experiments/dfs/main.go`)
3. **Dijkstra Algorithm** (`experiments/dijkstra/main.go`)

## Roles (rotating)
| Experiment | Satria Danish Khan | Joana Evangeline Handoko | Najiib Rahmawan |
|---|---|---|---|
| 1. Insertion Sort | Builder | Tester | Reviewer |
| 2. DFS | Reviewer | Builder | Tester |
| 3. Dijkstra | Tester | Reviewer | Builder |

## How to Run
Requires Go installed.

```
go run experiments/insertion_sort/main.go
go run experiments/dfs/main.go
go run experiments/dijkstra/main.go
```

## Repository Structure
```
algorithm-mini-lab-team-b/
|-- README.md
|-- experiments/
|-- screenshots/
|-- AI-NOTES.md
`-- REFLECTION.md
```

## Experiment Results

Each experiment follows: Predict -> Run -> Observe -> Change -> Explain.

### 1. Insertion Sort
| | Input | My prediction | Actual result |
|---|---|---|---|
| Original | `[5, 2, 8, 1, 4]` | `[2, 1, 4, 5, 8]` | `[1, 2, 4, 5, 8]` |
| Changed | `[9, 3, 7, 2, 1]` | `[3, 1, 2, 7, 9]` | `[1, 2, 3, 7, 9]` |

**What changed and why?**
_The input values were changed from [5, 2, 8, 1, 4] to [9, 3, 7, 2, 1]. The insertion sort algorithm itself stayed the same. I changed the numbers to test how the algorithm works with a different input. Even though the values were different, the final list was still arranged from the smallest number to the largest number. The algorithm compared each value with the numbers before it and moved the larger values to the right position. This shows that insertion sort can work with different types of input while still producing the correct result. The original order of the numbers did not affect the final sorting. The process was slightly different because the values had changed, but the main logic remained the same. After running the program, the list was sorted correctly in ascending order._

Screenshot: `screenshots/insertion-sort1.png`
Screenshot: `screenshots/insertion-sort2.png`
Screenshot: `screenshots/insertion-sort-terminal1.png`
Screenshot: `screenshots/insertion-sort-terminal2.png`

### 2. DFS
| | Graph | My prediction | Actual traversal |
|---|---|---|---|
| Original | A-B, A-C, B-D, B-E | A -> B -> D -> E -> C | A -> B -> D -> E -> C |
| Changed | + F under C | A -> B -> D -> E -> C -> F | A -> B -> D -> E -> C -> F |

**What changed and why?**
_The weight between nodes C and D was changed from 3 to 5. The Dijkstra algorithm itself was not changed. I modified the edge weight to see how increasing the distance between two nodes would affect the shortest path and its total cost. After changing the weight, the shortest path remained A → C → D → E, but the total cost increased from 7 to 9. The algorithm still compared the available routes and selected the one with the lowest total distance. Even though the C-D connection became longer, the other possible route was still more expensive. This shows that changing an edge weight can affect the cost of a path without necessarily changing the selected shortest path. The main logic of Dijkstra remained the same and the result was still correct._

Screenshot: `screenshots/DFS1.png`
Screenshot: `screenshots/DFS2.png`
Screenshot: `screenshots/DFS-terminal1.png`

### 3. Dijkstra
| | Network | My prediction | Shortest path | Cost |
|---|---|---|---|---|
| Original | C-D = 3 | A → C → D → E | A → C → D → E | 7 |
| Changed | C-D = 5 | A → C → D → E | A → C → D → E | 9 |

**What changed and why?**
_The graph was changed by adding a new node F under node C. The DFS algorithm itself was not changed. I added F to see what would happen to the traversal when a new node was added. Before the change, DFS visited A, B, D, E, and C. After adding F, the algorithm also visited F after reaching C. This made the traversal one step longer. DFS still worked in the same way, going deeper into a node before going back to the previous one. The main logic stayed the same, but the result changed because there was a new node to visit. This shows that DFS can follow a different graph without changing the algorithm itself._

Screenshot: `screenshots/djikstra1.png`
Screenshot: `screenshots/dijkstra2.png`
Screenshot: `screenshots/dijkstra3.png`
Screenshot: `screenshots/djikstra-terminal1.png`
Screenshot: `screenshots/djikstra-terminal2.png`

## Links
- AI usage: see [AI-NOTES.md](AI-NOTES.md)
- Individual reflections: see [REFLECTION.md](REFLECTION.md)
