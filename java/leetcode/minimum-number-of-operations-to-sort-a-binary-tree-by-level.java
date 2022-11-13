/*
You are given the root of a binary tree with unique values.

In one operation, you can choose any two nodes at the same level and swap their values.

Return the minimum number of operations needed to make the values at each level sorted in a strictly increasing order.

The level of a node is the number of edges along the path between it and the root node.
*/
import java.util.LinkedList;
import java.util.Queue;

class Solution {
    private static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode() {}
        TreeNode(int val) { this.val = val; }
        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
       }
    }

    private int countOps(int[] level) {
        int ops = 0;
        for (int i = 0; i < level.length; i++) {
            int min = level[i];
            int minIndex = i;
            for (int j = i+1; j < level.length; j++) {
                if (level[j] < min) {
                    min = level[j];
                    minIndex = j;
                }
                min = Math.min(min, level[j]);
            }
            if (min != level[i]) {
                int tmp = level[i];
                level[i] = min;
                level[minIndex] = tmp;
                ops++;
            }
        }
        return ops;
    }
    
    public int minimumOperations(TreeNode root) {
        int minOps = 0;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.offer(root);
        while (!q.isEmpty()) {
            int n = q.size();
            int[] level = new int[n];
            for (int i = 0; i < n; i++) {
                TreeNode node = q.poll();
                if (node.left != null) q.offer(node.left);
                if (node.right != null) q.offer(node.right);
                level[i] = node.val;
            }
            minOps += countOps(level);
        }
        return minOps;
    }
}