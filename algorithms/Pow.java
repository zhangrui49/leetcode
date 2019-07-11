package algorithms;

public class Pow {

    public static void main(String[] args) {
        System.out.println(Math.pow(2, 10));
        System.out.println(pow(2L, 9));
    }


   static long pow(long value, int pow) {
        if (pow == 0) {
            return 1;
        }
        if (pow == 1) {
            return value;
        }
        if (pow % 2 == 0) {
            return pow(value * value, pow / 2);
        } else {
            return pow(value * value, pow / 2)*value;
        }

    }

}
