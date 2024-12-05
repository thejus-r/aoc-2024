# --- Day 5: Print Queue ---

## Solution: Check for topological order

Given a graph G and a sequence seq:

1. Build an in-degree map: Track the number of incoming edges for each node.
2. Simulate processing the nodes in seq:
   - For each node in seq, check if its in-degree is 0 (i.e., it is ready to be processed).
   - If the in-degree is not 0, the sequence is invalid.
   - Remove the node from the graph by decrementing the in-degrees of its neighbors.

If all nodes in seq pass these checks, the sequence is a valid topological order.
