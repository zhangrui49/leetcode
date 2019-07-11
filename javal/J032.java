package javal;


public class J032 {

    public static void main(String[] args) {
        J032 j032 = new J032();
       // System.out.println(j032.longestValidParentheses("())())))))))))()(()))(()))())(()))()))(((())))()))(((())()())()(()))(()((())))((()(())))(())(()()))))())((())())))()()(()))())(((())((()((()(()()))(()(()((()())((())(()(()((((())())()(()()()"));
        //System.out.println(j032.longestValidParentheses("(()"));
        System.out.println(j032.longestValidParentheses(")))()())())()())))()((()))(()))((())(((())())(())(())((((()))(()()(()())()())((((())))()()())())))))(((((())((((()())(()))(()))))(())())(((("));
    }

    public int longestValidParentheses(String s) {
        StringBuilder stringBuilder = new StringBuilder(s);
        String temp = "";
        int len = s.length();
        for (int j = len; j > 0; j--) {
            for (int k = 0; k <= len - j; k++) {
                temp = stringBuilder.substring(k, k + j);
                if (isValid(temp)) {
                    return temp.length();
                }
            }
        }
        return 0;
    }

    public boolean isValid(String s) {

        char[] characters = new char[s.length() + 1];
        int index = 1;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(') {
                characters[index++] = c;
            } else if (c == ')' && characters[--index] != '(') {
                return false;
            }
        }
        return index == 1;
    }
}


