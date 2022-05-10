package javal.algor;

import java.util.Arrays;

public class J037 {

    public static void main(String[] args) {
        J037 j037 = new J037();
        j037.solveSudoku(new char[][]{
                "187654329".toCharArray(),
                "2........".toCharArray(),
                "3........".toCharArray(),
                "4........".toCharArray(),
                "5........".toCharArray(),
                "6........".toCharArray(),
                "7........".toCharArray(),
                "8........".toCharArray(),
                "9........".toCharArray(),
        });
    }

    public void solveSudoku(char[][] board) {
        sudoku(board, 0, 0);
        for (int i = 0; i < 9; i++) {
            System.out.println(Arrays.toString(board[i]));
        }

    }

    public boolean sudoku(char[][] board, int i, int j) {
        if (j >= 9) {
            for (int k = 0; k < i; k++) {
                System.out.println(k+":"+Arrays.toString(board[k]));
            }
            System.out.println("");
            System.out.println("");
            return sudoku(board, i + 1, 0);
        }
        if (i >= 9) {
            return true;
        }
        char c = board[i][j];
        //  System.out.println("i is " + i + "  j is  " + j + "  ch is: " + c);
        if (c != '.') {
            return sudoku(board, i, j + 1);
        } else {
            for (int k = 1; k <= 9; k++) {
                char cc = (char) (k + '0');
                board[i][j] = cc;
                // System.out.println("ch new " + cc);
                if (isValidSudoku(i, j, board)) {
                    if (sudoku(board, i, j + 1)) {
                        return true;
                    }
                }
                board[i][j] = '.';
            }
        }
        return false;
    }

    public boolean isValidSudoku(int i, int j, char[][] board) {
        for (int k = 0; k < 9; k++) {
            if (k != j && board[i][k] == board[i][j])
                return false;
        }
        for (int k = 0; k < 9; k++) {
            if (k != i && board[k][j] == board[i][j])
                return false;
        }
        for (int row = i / 3 * 3; row < i / 3 * 3 + 3; row++) {
            for (int col = j / 3 * 3; col < j / 3 * 3 + 3; col++) {
                if ((row != i || col != j) && board[row][col] == board[i][j])
                    return false;
            }
        }
        return true;
    }
}
