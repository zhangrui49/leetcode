package java;

/**
 * J001
 */
public class J001 {

    public static void main(String[] args) {
        int nums[] = {11, 7, 11, 2,4,67,9,33};
        int target = 44;
        int[] result = twoSum(nums, target);
        if (result != null) {
            System.out.println(result[0] + "--" + result[1]);
        }
    }

    static int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if ((nums[i] + nums[j]) == target) {
                    return new int[]{i, j};
                }
            }
        }
        return null;
    }
}