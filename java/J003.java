package java;

public class J003 {

    public static void main(String[] args) {
        J003 p3 = new J003();
        System.out.println(p3.lengthOfLongestSubstring("bbbbbb"));
    }

    public int lengthOfLongestSubstring(String s) {
        int len = s.length();
        int[] hash = new int[256];
        int left = 0, max = 0;
        for (int i = 0; i < len; i++) {
            char c = s.charAt(i);
            if (hash[c] > left) {
                left = hash[c];
            }
            int l = i - left +1;
            hash[c] = i + 1;
            if (l > max) {
                max = l;
            }
        }
        return max;
    }
}
