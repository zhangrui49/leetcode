package algorithms;

import java.util.*;
import java.util.function.Consumer;

public class MyArrayList<T> implements List<T> {
    private static final int DEFAULT_CAPACITY = 16;
    private int size = 0;
    private T[] items;


    public MyArrayList() {
        clear();
    }

    public MyArrayList(int capacity) {
        if (capacity >= 0) {
            this.items = (T[]) new Object[capacity];
        } else {
            throw new IllegalArgumentException("Illegal argument,capacity must larger than zero");
        }
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public boolean contains(Object o) {


        return indexOf(o) > -1;
    }

    @Override
    public Iterator<T> iterator() {
        return new ArrayListIterator();
    }

    @Override
    public Object[] toArray() {

        return Arrays.copyOf(items, size);
    }

    @Override
    public <T> T[] toArray(T[] a) {
        return Arrays.copyOf((T[]) items, size);
    }

    @Override
    public boolean add(T t) {
        add(size(), t);
        return true;
    }

    @Override
    public boolean remove(Object o) {
        int idx = indexOf(o);
        remove(idx);
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
        this.items = (T[]) new Object[]{};
        size = 0;
    }

    @Override
    public T get(int index) {

        if (index < 0 || index >= size) {
            throw new IndexOutOfBoundsException();
        }

        return items[index];
    }

    @Override
    public T set(int index, T element) {

        T old = get(index);
        items[index] = element;
        return old;
    }

    @Override
    public void add(int index, T element) {
        if (items.length == size()) {
            ensureCapacity(size * 2);
        }
        for (int i = size; i > index; i--) {
            items[i] = items[i - 1];
        }
        items[index] = element;
        size++;
    }

    @Override
    public T remove(int index) {
        T remove = items[index];
        for (int i = index; i < size() - 1; i++) {
            items[i] = items[i + 1];
        }
        size--;
        return remove;
    }

    @Override
    public int indexOf(Object o) {
        for (int i = 0; i < items.length; i++) {
            if (o.equals(items[i])) {
                return i;
            }
        }
        return -1;
    }

    @Override
    public int lastIndexOf(Object o) {
        for (int i = items.length - 1; i >= 0; i--) {
            if (o.equals(items[i])) {
                return i;
            }
        }

        return -1;
    }

    @Override
    public ListIterator<T> listIterator() {
        return null;
    }

    @Override
    public ListIterator<T> listIterator(int index) {
        return null;
    }

    @Override
    public List<T> subList(int fromIndex, int toIndex) {
        return null;
    }

    private void ensureCapacity(int newSize) {
        if (newSize < size) {
            return;
        }
        T[] old = items;
        items = (T[]) new Object[newSize];
        for (int i = 0; i < old.length; i++) {
            items[i] = old[i];
        }
    }


    public class ArrayListIterator implements Iterator<T> {

        private int current = 0;

        @Override
        public boolean hasNext() {
            return current < size();
        }

        @Override
        public T next() {
            if (!hasNext()) {
                throw new NoSuchElementException();
            }
            return items[current++];
        }

        @Override
        public void remove() {
            MyArrayList.this.remove(--current);
        }

        @Override
        public void forEachRemaining(Consumer<? super T> action) {

        }
    }
}
