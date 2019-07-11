package algorithms;


import sun.reflect.generics.tree.Tree;

public class BinarySearchTree<T extends Comparable<? super T>> {


    private BinaryNode<T> root;


    private class BinaryNode<T> {
        T item;
        BinaryNode<T> left;
        BinaryNode<T> right;

        public BinaryNode(T item) {
            this(item, null, null);
        }

        public BinaryNode(T item, BinaryNode<T> left, BinaryNode<T> right) {
            this.item = item;
            this.left = left;
            this.right = right;
        }
    }


}
