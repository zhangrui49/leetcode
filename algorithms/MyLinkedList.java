package algorithms;

import java.util.*;

public class MyLinkedList<T> implements List<T> {

    private int size;
    private Node<T> first;
    private Node<T> last;

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size() == 0;
    }

    @Override
    public boolean contains(Object o) {
        return indexOf(o) != -1;
    }

    @Override
    public Iterator<T> iterator() {
        return new LinkedIterator();
    }

    @Override
    public Object[] toArray() {
        Node start = first;
        Object[] array = new Object[size];
        for (int i = array.length - 1; i >= 0; i--) {
            array[i] = start.item;
            start = start.next;
        }
        return array;
    }

    @Override
    public <T> T[] toArray(T[] a) {
        Node<T> start = (Node<T>) first;
        T[] array = (T[]) new Object[size];
        for (int i = array.length - 1; i >= 0; i--) {
            array[i] = start.item;
            start = start.next;
        }
        return array;
    }

    @Override
    public boolean add(T t) {
        add(size(), t);
        return true;
    }

    @Override
    public boolean remove(Object o) {
        remove(indexOf(o));
        return true;
    }

    @Override
    public boolean containsAll(Collection<?> c) {
        return false;
    }

    @Override
    public boolean addAll(Collection<? extends T> c) {
        return false;
    }

    @Override
    public boolean addAll(int index, Collection<? extends T> c) {
        return false;
    }

    @Override
    public boolean removeAll(Collection<?> c) {
        return false;
    }

    @Override
    public boolean retainAll(Collection<?> c) {
        return false;
    }

    @Override
    public void clear() {
        first = new Node<>(null, null, null);
        last = new Node<>(first, null, null);
        first.next = last;
        size = 0;
    }

    @Override
    public T get(int index) {
        return getNode(index).item;
    }

    @Override
    public T set(int index, T element) {
        Node<T> node = getNode(index);
        T old = node.item;
        node.item = element;
        return old;
    }

    @Override
    public void add(int index, T element) {
        Node<T> node = getNode(index);
        Node<T> newNode = new Node<>(node.prev, node.next, element);
        newNode.prev.next = newNode;
        node.prev = newNode;
        size++;
    }

    @Override
    public T remove(int index) {
        Node<T> node = getNode(index);
        T old = node.item;
        node.prev.next = node.next;
        node.next.prev = node.prev;
        size--;
        return old;
    }

    @Override
    public int indexOf(Object o) {
        int index = 0;
        for (Node<T> x = first; x != null; x = x.next) {
            if (o.equals(x.item))
                return index;
            index++;
        }
        return -1;
    }

    @Override
    public int lastIndexOf(Object o) {
        int index = 0;
        for (Node<T> x = last; x != null; x = x.prev) {
            if (o.equals(x.item))
                return index;
            index++;
        }
        return -1;
    }

    @Override
    public ListIterator<T> listIterator() {
        return new LinkedIterator();
    }

    @Override
    public ListIterator<T> listIterator(int index) {
        return new LinkedIterator();
    }

    @Override
    public List<T> subList(int fromIndex, int toIndex) {
        return null;
    }


    private class Node<T> {
        private Node<T> prev;
        private Node<T> next;
        private T item;

        public Node(Node<T> prev, Node<T> next, T item) {
            this.prev = prev;
            this.next = next;
            this.item = item;
        }
    }

    private Node<T> getNode(int index) {
        Node<T> node;
        if (index < 0 || index > size()) {
            throw new ArrayIndexOutOfBoundsException();
        }
        if (index < size() / 2) {
            node = first.next;
            for (int i = 0; i < index; i++) {
                node = node.next;
            }
        } else {
            node = last.prev;
            for (int i = size(); i > index; i--) {
                node = node.prev;
            }
        }
        return node;
    }


    private class LinkedIterator implements ListIterator<T> {
        private Node<T> current = first.next;

        @Override
        public boolean hasNext() {
            return current != last;
        }

        @Override
        public T next() {
            T element = current.item;
            current = current.next;
            return element;
        }

        @Override
        public boolean hasPrevious() {
            return current != first;
        }

        @Override
        public T previous() {
            return current.prev.item;
        }

        @Override
        public int nextIndex() {
            return indexOf(current.next);
        }

        @Override
        public int previousIndex() {
            return indexOf(current.prev);
        }

        @Override
        public void remove() {
            MyLinkedList.this.remove(current.prev);
        }

        @Override
        public void set(T t) {
            MyLinkedList.this.set(indexOf(current), t);
        }

        @Override
        public void add(T t) {
            MyLinkedList.this.add(t);
        }
    }
}
