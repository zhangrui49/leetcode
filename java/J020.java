package java;

public class J020 {

    public static void main(String[] args) {
        J020 p20 = new J020();
        System.out.println(p20.isValid("{()}()()"));

    }

    public boolean isValid(String s) {

        char[] characters = new char[s.length()+1];
        int index = 1;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(' || c == '[' || c == '{') {
                characters[index++]= c;
            } else if (c == ')' && characters[--index] != '(') {
                return false;
            } else if (c == ']' && characters[--index] != '[') {
                return false;
            } else if (c == '}' && characters[--index] != '{') {
                return false;
            }
        }
        return index == 1;
    }
}
