class Solution {
public:
    int maxArea(vector<int>& height) {
        // 双指针
        int res = 0, temp = 0;
        int i = 0, j = height.size() - 1;
        while (i < j) {
            if (height[i] < height[j]) {
                temp = height[i] * (j-i);
                i++;
            } else {
                temp = height[j] * (j-i);
                j--;
            }
            if (temp > res) 
                res = temp;
        }
        return res;
    }
};
