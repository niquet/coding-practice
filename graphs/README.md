# Graphs

- Graph: non-linear data structure used to represent relationships between discrete objects/entities
- Non-linear: do not follow a sequential order
- Consists of:
  - Nodes/vertices representing the individual objects within the graph
  - Edges representing the connections/relationships between the nodes
- Edges can be directed or undirected and can optionally have values (a weighted graph)
- Trees are undirected graphs in which any two vertices are connected by exactly one edge and there can be no cycles in the graph
- Commonly used to model relationship between unordered entities

## Graph representations

- Common graph representations:
  - Adjacency matrix
  - Adjacency list
  - Hash table of hash tables
- _Hash table of hash tables is the simplest approach during interviews_
- _In interviews, graphs are commonly given in the input as 2D matrices where cells are the nodes and each cell can traverse to its adjacent cells (up/down/left/right) Hence it is important that you be familiar with traversing a 2D matrix. When traversing the matrix, always ensure that your current position is within the boundary of the matrix and has not been visited before.

---

## Types of graphs

- Undirected graph: edges have no direction (connection between two nodes is bidirectional)
- Directed graph: edges have a specific direction (indicates a one-way relationship from one node to another)
- Weighted graph: edges have associated numerical values (weights) representing costs, distances, or capacities

## Definitions

- Connectivity: graphs can be connected (all nodes are reachable from each other) or disconnected
- Cycles: graphs can contain cycles (a path starts and ends at the same node) or be acyclic (no cycles)
- Shortest path algorithms: find the most efficient path betwen two nodes

## Applications

- Dependency management: represent task dependencies
- Recommendation systems: suggest products or content based on user preferences and connections

## Credits

- [https://www.techinterviewhandbook.org/algorithms/graph/](https://www.techinterviewhandbook.org/algorithms/graph/)
