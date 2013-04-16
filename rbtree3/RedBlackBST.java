import java.util.Iterator;
import java.util.NoSuchElementException;

public class RedBlackBST<Key extends Comparable<Key>, Value> 
                          implements Iterable<Key>
{
    private static final boolean RED   = true;
    private static final boolean BLACK = false;

    private Node root;      // root of the BST
    private int N;          // size of BST
    private int k;

    private class Node
    {
        Key key;            // key
        Value value;        // associated data
        Node left, right;   // left and right subtrees
        private double xc, yc;
        boolean color;      // color of parent link

        Node(Key key, Value value, boolean color)
        {
            this.key   = key;
            this.value = value;
            this.color = color;
        }
    }

    public boolean contains(Key key)
    {  return (get(key) != null);  }

    public Value get(Key key)
    {  return get(root, key);  }

    private Value get(Node x, Key key)
    {
        if (x == null)        return null;
        if (eq  (key, x.key)) return x.value;
        if (less(key, x.key)) return get(x.left,  key);
        else                  return get(x.right, key);
    }

    public void put(Key key, Value value)
    {
        if (!contains(key)) N++;
        root = insert(root, key, value);
        root.color = BLACK;
    }

   private Node insert(Node h, Key key, Value value)
   { 
      if (h == null) 
         return new Node(key, value, RED);

      if (isRed(h.left))
         if (isRed(h.left.left))
            h = splitFourNode(h);

      if (eq(key, h.key))
         h.value = value;
      else if (less(key, h.key)) 
         h.left = insert(h.left, key, value); 
      else 
         h.right = insert(h.right, key, value); 

      if (isRed(h.right))
         h = leanLeft(h);

      return h;
   }

   public void deleteMin()
   {
      root = deleteMin(root);
      root.color = BLACK;
      N--;
   }

   private Node deleteMin(Node h)
   { 
      if (h.left == null)
         return null;

      if (!isRed(h.left) && !isRed(h.left.left))
         h = moveRedLeft(h);

      h.left = deleteMin(h.left);

      if (isRed(h.right))
         h = leanLeft(h);

      return h;
   }

   public void deleteMax()
   {
      root = deleteMax(root);
      root.color = BLACK;
      N--;
   }

   private Node deleteMax(Node h)
   { 
      if (h.right == null)
      {  
         if (h.left != null)
            h.left.color = BLACK;
         return h.left;
      }

      if (isRed(h.left))
         h = leanRight(h);

      if (!isRed(h.right) && !isRed(h.right.left))
         h = moveRedRight(h);

      h.right = deleteMax(h.right);

      if (isRed(h.right))
         h = leanLeft(h);

      return h;
   }

   public void delete(Key key)
   { 
      root = delete(root, key);
      root.color = BLACK;
      N--;
   }

   private Node delete(Node h, Key key)
   { 
      if (less(key, h.key)) 
      {
         if (!isRed(h.left) && !isRed(h.left.left))
            h = moveRedLeft(h);
         h.left =  delete(h.left, key);
      }
      else 
      {
         if (isRed(h.left)) 
            h = leanRight(h);
         if (eq(key, h.key) && (h.right == null))
            return null;
         if (!isRed(h.right) && !isRed(h.right.left))
            h = moveRedRight(h);
         if (eq(key, h.key))
         {
            h.value = get(h.right, min(h.right));
            h.key = min(h.right);
	    h.right = deleteMin(h.right);
         }
         else h.right = delete(h.right, key);
      }
 
      if (isRed(h.right))
         h = leanLeft(h);

      return h;
   }

// Helper methods

    private boolean less(Key a, Key b) { return a.compareTo(b) <  0; }
    private boolean eq  (Key a, Key b) { return a.compareTo(b) == 0; }

    private boolean isRed(Node x)
    {
        if (x == null) return false;
        return (x.color == RED);
    }

    private Node rotR(Node h)
    {  // Rotate right.
       Node x = h.left;
       h.left = x.right;
       x.right = h;
       return x;
    }

    private Node rotL(Node h)
    {  // Rotate left.
       Node x = h.right;
       h.right = x.left;
       x.left = h;
       return x;
    }

   private Node splitFourNode(Node h)
   {  // Rotate right, then flip colors.
      h = rotR(h);
      //      h.color       = RED;
      h.left.color  = BLACK;
      return h;
   }

   private Node leanLeft(Node h)
   {  // Make a right-leaning 3-node lean to the left.
      h = rotL(h);
      h.color      = h.left.color;                   
      h.left.color = RED;                     
      return h;
   }

   private Node leanRight(Node h)
   {  // Make a left-leaning 3-node lean to the right.
      h = rotR(h);
      h.color       = h.right.color;                   
      h.right.color = RED;                     
      return h;
    }

    private Node moveRedLeft(Node h)
    {  // Assuming that h is red and both h.left and h.left.left
       // are black, make h.left or one of its children red.
       h.color      = BLACK;
       h.left.color = RED;  
       if (!isRed(h.right.left)) h.right.color = RED; else
       { 
          h.right = rotR(h.right);
          h = rotL(h);
       }
      return h;
    }

    private Node moveRedRight(Node h)
    {  // Assuming that h is red and both h.right and h.right.left
       // are black, make h.right or one of its children red.
       h.color      = BLACK;
       h.right.color = RED;  
       if (!isRed(h.left.left)) h.left.color = RED; else
       { 
          h = rotR(h);
          h.color = RED;
          h.left.color = BLACK;
       }
       return h;
    }

// Methods for tree drawing

    public void draw(double y, double lineWidth, double nodeSize)
    { 
      k = 1; 
      setcoords(root, y);
      StdDraw.setPenColor(StdDraw.BLACK);
      StdDraw.setPenRadius(lineWidth);
      drawlines(root); 
      StdDraw.setPenColor(StdDraw.WHITE);
      drawnodes(root, nodeSize); 
    }

    public void setcoords(Node x, double d)
    {
        if (x == null) return;
        setcoords(x.left, d-.04);
        x.xc = (1.0 * k++)/N; x.yc = d - .04;
        setcoords(x.right, d-.04);
    }

    public void drawlines(Node x)
    {
        if (x == null) return;
        drawlines(x.left);
        if (x.left != null)
        {
	    if (x.left.color == RED) StdDraw.setPenColor(StdDraw.RED);
	    else StdDraw.setPenColor(StdDraw.BLACK);
            StdDraw.line(x.xc, x.yc, x.left.xc, x.left.yc);
        }
        if (x.right != null)
        {
	    if (x.right.color == RED) StdDraw.setPenColor(StdDraw.RED);
	    else StdDraw.setPenColor(StdDraw.BLACK);
            StdDraw.line(x.xc, x.yc, x.right.xc, x.right.yc);
        }
        drawlines(x.right);
    }

    public void drawnodes(Node x, double nodeSize)
    {
        if (x == null) return;
        drawnodes(x.left, nodeSize);
	StdDraw.filledCircle(x.xc, x.yc, nodeSize);
        drawnodes(x.right, nodeSize);
    }

    public void mark(Key key)
    { 
      StdDraw.setPenColor(StdDraw.BLACK);
      marknodes(key, root); 
    }

    public void marknodes(Key key, Node x)
    {
        if (x == null) return;
        marknodes(key, x.left);
	if (eq(key, x.key))
           StdDraw.filledCircle(x.xc, x.yc, .004);
        marknodes(key, x.right);
    }

// Utility functions

    // return number of key-value pairs in symbol table
    public int size() { return N; }

    // height of tree (empty tree height = 0)
    public int height() { return height(root); }
    private int height(Node x) {
        if (x == null) return 0;
        return 1 + Math.max(height(x.left), height(x.right));
    }

    // return the smallest key
    public Key min() { 
        return min(root);
    }

    public Key min(Node x) { 
        Key key = null;
        for ( ; x != null; x = x.left)
            key = x.key;
        return key;
    }

    // return the largest key
    public Key max() {
        Key key = null;
        for (Node x = root; x != null; x = x.right)
            key = x.key;
        return key;
    }

// Tree iterator

   public Iterator<Key> iterator() 
   { return new BSTIterator(root); }

   private class BSTIterator implements Iterator<Key>
   {
      Stack<Node> stack = new Stack<Node>();

      BSTIterator(Node x)
      {
         while (x != null)
         {
            stack.push(x);
            x = x.left;
         }
      }

      public boolean hasNext()  
      { return !stack.isEmpty();                    }
      public void remove()      
      { throw new UnsupportedOperationException();  }

      public Key next()
      {
         if (!hasNext()) throw new NoSuchElementException();
         Node x = stack.pop();
         Key key = x.key;
         x = x.right;
         while(x != null)
         {
            stack.push(x);
            x = x.left;
         }
         return key;
      }
}



// Integrity checks

   public boolean check() 
   {  // Is this tree a red-black tree?
      return isBST() && is234() && isBalanced();
   }

   private boolean isBST() 
   {  // Is this tree a BST?
      return isBST(root, min(), max());
   }

   private boolean isBST(Node x, Key min, Key max)
   {  // Are all the values in the BST rooted at x between min and max,
      // and does the same property hold for both subtrees?
      if (x == null) return true;
      if (less(x.key, min) || less(max, x.key)) return false;
      return isBST(x.left, min, x.key) && isBST(x.right, x.key, max);
   } 

   private boolean is234() { return is234(root); }
   private boolean is234(Node x)
   {  // Does the tree have no red right links, and at most two (left)
      // red links in a row on any path?
      if (x == null) return true;
      if (isRed(x.right)) return false;
      if (isRed(x))
        if (isRed(x.left))
          if (isRed(x.left.left)) return false;
      return is234(x.left) && is234(x.right);
   } 

   private boolean isBalanced()
   { // Do all paths from root to leaf have same number of black edges?
      int black = 0;     // number of black links on path from root to min
      Node x = root;
      while (x != null)
      {
         if (!isRed(x)) black++;
            x = x.left;
      }
      return isBalanced(root, black);
   }

   private boolean isBalanced(Node x, int black)
   { // Does every path from the root to a leaf have the given number 
     // of black links?
      if      (x == null && black == 0) return true;
      else if (x == null && black != 0) return false;
      if (!isRed(x)) black--;
      return isBalanced(x.left, black) && isBalanced(x.right, black);
   } 


// Basic test client


   public static void main(String[] args)
   {

        // test some fixed inputs
        RedBlackBST<String, String> st = new RedBlackBST<String, String>();

        // insert some key-value pairs
        st.put("www.cs.princeton.edu",   "128.112.136.11");
        st.put("www.cs.princeton.edu",   "128.112.136.35");
        st.put("www.princeton.edu",      "128.112.130.211");
        st.put("www.math.princeton.edu", "128.112.18.11");
        st.put("www.yale.edu",           "130.132.51.8");
        st.put("www.amazon.com",         "207.171.163.90");
        st.put("www.simpsons.com",       "209.123.16.34");
        st.put("www.stanford.edu",       "171.67.16.120");
        st.put("www.google.com",         "64.233.161.99");
        st.put("www.ibm.com",            "129.42.16.99");
        st.put("www.apple.com",          "17.254.0.91");
        st.put("www.slashdot.com",       "66.35.250.150");
        st.put("www.whitehouse.gov",     "204.153.49.136");
        st.put("www.espn.com",           "199.181.132.250");
        st.put("www.snopes.com",         "66.165.133.65");
        st.put("www.movies.com",         "199.181.132.250");
        st.put("www.cnn.com",            "64.236.16.20");
        st.put("www.iitb.ac.in",         "202.68.145.210");

        // search for IP addresses given URL
        System.out.println("size = " + st.size());
        System.out.println(st.get("www.cs.princeton.edu"));
        System.out.println(st.get("www.harvardsucks.com"));
        System.out.println(st.get("www.simpsons.com"));
        System.out.println();

        // test out the iterator
        for (String s : st) 
            System.out.println(s);

        System.out.println();
        System.out.println("Valid red-black tree: " + st.check());
        System.out.println();


        // insert N elements in order if one command-line argument supplied
        if (args.length == 0) return;
        int N = Integer.parseInt(args[0]);
        RedBlackBST<Integer, Integer> st2 = new RedBlackBST<Integer, Integer>();
        for (int i = 0; i < N; i++) {
            st2.put(i, i);
            int h = st2.height();
            System.out.println("i = " + i + ", height = " + h);
            if (!st2.check()) System.out.println("integrity check failed");
        }
        System.out.println("size = " + st2.size());
        System.out.println("integrity check " + st2.check());
    }
}
