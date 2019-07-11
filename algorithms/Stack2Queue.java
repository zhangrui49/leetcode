package algorithms;

import java.util.Stack;

public class Stack2Queue {
    private Stack<Integer> popStack;
    private Stack<Integer> pushStack;

    public Stack2Queue() {
        this.popStack = new Stack<>();
        this.pushStack = new Stack<>();
    }

    private void pushToPop() {
        if (popStack.empty()) {
            while (!pushStack.empty()) {
                popStack.push(pushStack.pop());
            }
        }
    }

    public int poll() throws RuntimeException {
        pushToPop();
        if (popStack.isEmpty()) {
            throw new RuntimeException("queue is empty");
        }
        return popStack.pop();
    }

    public int peek() throws RuntimeException {
        pushToPop();
        if (popStack.empty()) {
            throw new RuntimeException("queue is empty");
        }
        return popStack.peek();
    }

    public void add(int element) {
        pushStack.push(element);
        pushToPop();
    }

}
