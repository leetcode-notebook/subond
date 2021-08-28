class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        if (0 == nums.size()) {
            return 0;
        }
        int res = 0;
        for (auto v : nums) {
            if (target > v) {
                res++;
            }
        }
        return res;
    }
};
