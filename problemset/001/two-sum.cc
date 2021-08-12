class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> res = {};
        auto count = nums.size();
        for (auto i = 0; i < count; i++) {
            for (auto j = i+1; j < count; j++) {
                if (target == nums[i] + nums[j]) {
                    res.push_back(i);
                    res.push_back(j);
                }
            }
        }
        return res;
    }
};
