package algorithms;

public class MaxSubSum {

    public static void main(String[] args) {
        System.out.println(maxSum(new int[]{-2, 11, -4, 13, -5, -2}));
        System.out.println(maxSum2(new int[]{-2, 11, -4, 13, -5, -2}));
    }


    public static int maxSum(int[] arr) {
        int maxSum = 0;

        for (int i = 0; i < arr.length; i++) {
            int tempSum = 0;
            for (int j = i; j < arr.length; j++) {
                tempSum += arr[j];
                if (tempSum > maxSum) {
                    maxSum = tempSum;
                }
            }
        }

        return maxSum;
    }

    public static int maxSum2(int[] arr) {
        int maxSum = 0;
        int tempSum = 0;
        for (int i = 0; i < arr.length; i++) {
            tempSum += arr[i];
            if (tempSum > maxSum) {
                maxSum = tempSum;
            } else if (tempSum < 0) {
                tempSum = 0;
            }
        }

        return maxSum;
    }
}
