package algorithms;

public class Fib {

    public static void main(String[] args) {
        long date1 = System.currentTimeMillis();
        System.out.println(fib(40));
        long date2 = System.currentTimeMillis();
        System.out.println("fib1:" + (date2 - date1));
        System.out.println(fib2(40));
        System.out.println("fib2:" + (System.currentTimeMillis() - date2));
    }

    public static long fib(int n) {

        if (n <= 1) {
            return 1;
        }
        return fib(n - 1) + fib(n - 2);

    }


    public static long fib2(int n) {
        long[] c = new long[n + 2];
        c[0] = 1L;
        c[1] = 1L;
        for (int i = 2; i < n + 2; i++) {
            c[i] = c[i - 1] + c[i - 2];
        }
        return c[n];
    }
}


