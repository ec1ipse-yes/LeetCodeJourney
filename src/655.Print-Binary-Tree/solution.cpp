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
    vector<vector<string>> printTree(TreeNode* root) {
        // Get height and width of Matrix
        int h = getHeight(root);
        int w = (1 << h) - 1;
        // Initialize matrix
        vector<vector<string>> matrix(h, vector<string>(w, ""));
        // Fill matrix according to biselection
        fillMatrix(root, matrix, 0, 0, w-1);
        return matrix;
    }
private:
    int getHeight(TreeNode* root) {
        if(!root) return 0;
        return max(getHeight(root->left), getHeight(root->right)) + 1;
    }

    void fillMatrix(TreeNode *root, vector<vector<string>>& matrix, int h, int l, int r) {
        if(!root) return;
        int mid = (l + r) / 2;
        matrix[h][mid] = std::to_string(root->val);
        fillMatrix(root->left, matrix, h+1, l, mid-1);
        fillMatrix(root->right, matrix, h+1, mid+1, r);
    }
};
