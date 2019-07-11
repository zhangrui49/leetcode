package algorithms;

public class BinarySearch {
    private static final int NOT_FOUND = -1;


    public static void main(String[] args) {
        System.out.println(binarySearch(new int[]{1, 2, 3, 4, 5, 6, 7, 8}, 7));
    }

    public static int binarySearch(int[] arr, int value) {
        int start = 0;
        int end = arr.length - 1;
        while (start <= end) {
            int mid = (start + end) / 2;
            if (arr[mid] < value) {
                start = mid + 1;
            } else if (arr[mid] > value) {
                end = mid - 1;
            } else {
                return mid;
            }
        }
        return NOT_FOUND;
    }
}
