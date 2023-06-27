package avltrees

// In a balanced tree the difference between the height of the left tree and right tree should be less 1, 
// 			height(left) - height(right) <= 1

// AVL Trees is an example of self balancing trees, below are the list of self balancing trees inview
// AVL ==> (Adelson-Vesky and Landis)(Easiest to understand)
// Red-black trees, B-trees, Splay Trees, 2-3 Trees

// AVL trees are a special type of BST that automatically rebalance every time we add or remove nodes
// They do this by ensuring the defference between the height of the left and right sub trees are never more than one(1)

// Rotations:
//  ===>> Left(LL)
//  ===>> Right(RR)
//  ===>> Right-Left(LR)
//  ===>> Right-Left(RL)
