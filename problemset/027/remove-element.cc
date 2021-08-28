class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int res = 0;
        for (auto v : nums) {
            if (v != val) {
                nums[res] = v;
                res++;
            }
        }
        return res;
    }
};
