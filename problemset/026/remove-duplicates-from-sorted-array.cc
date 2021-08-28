class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int length = nums.size();
        if (length < 2) return length;
        int res = 1;
        int pre = nums[0];
        for (int i = 1; i < length; i++) {
            if (nums[i] != pre) {
                nums[res] = nums[i];
                res++;
                pre = nums[i];
            }
        }
        return res;
    }
};
