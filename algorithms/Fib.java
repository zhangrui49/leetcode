package algorithms;

public class Fib {

    public static void main(String[] args) {
        System.out.println(fib(2));
    }

    public static long fib(int n) {
        if (n <= 1) {
            return 1;
        }
        return fib(n - 1) + fib(n - 2);
    }

    static long left = 0;

    public static long fib2(int n) {

        if (n <= 1) {
            return 1;
        }
        left = fib2(n - 2);
        return left + fib(n - 1);
    }
}


