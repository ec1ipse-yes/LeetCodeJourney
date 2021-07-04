/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    string tree2str(TreeNode* root) {
        if (root == nullptr) return "";
        const string s = std::to_string(root->val);
        const string l = tree2str(root->left);
        const string r = tree2str(root->right);

        // Special case1: left + right node == nullptr -> s
        if (root->left == nullptr && root->right == nullptr) return s;
        // Special case2: right == nullptr; left != nullptr -> s(l)
        if (root->left != nullptr && root->right == nullptr) return s+"("+l+")";
        // General case: s(l)(r)
        return s+"("+l+")"+"("+r+")";
    }
};