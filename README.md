# Mazes Chapter 2

Chapter 2 is all about building a grid class that contains cells. We can
populate our maze and print it out.

Running this produces a random maze with the btree algorithm

```
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
|       |       |       |       |   |               |   |               |       |
+---+   +---+   +---+   +---+   +   +---+---+---+   +   +---+---+---+   +---+   +
|       |   |   |   |                           |   |   |                       |
+---+   +   +   +   +---+---+---+---+---+---+   +   +   +---+---+---+---+---+   +
|   |                   |       |                       |   |   |   |   |       |
+   +---+---+---+---+   +---+   +---+---+---+---+---+   +   +   +   +   +---+   +
|       |   |               |   |   |   |   |       |   |   |   |   |       |   |
+---+   +   +---+---+---+   +   +   +   +   +---+   +   +   +   +   +---+   +   +
|   |           |   |   |   |   |   |   |   |   |   |   |   |   |   |       |   |
+   +---+---+   +   +   +   +   +   +   +   +   +   +   +   +   +   +---+   +   +
|       |       |   |   |   |       |   |       |   |   |   |           |       |
+---+   +---+   +   +   +   +---+   +   +---+   +   +   +   +---+---+   +---+   +
|   |       |       |       |           |   |       |   |       |   |   |       |
+   +---+   +---+   +---+   +---+---+   +   +---+   +   +---+   +   +   +---+   +
|   |   |           |       |                   |           |   |               |
+   +   +---+---+   +---+   +---+---+---+---+   +---+---+   +   +---+---+---+   +
|   |           |   |       |   |       |   |       |   |           |   |       |
+   +---+---+   +   +---+   +   +---+   +   +---+   +   +---+---+   +   +---+   +
|   |   |   |   |           |                   |   |                   |       |
+   +   +   +   +---+---+   +---+---+---+---+   +   +---+---+---+---+   +---+   +
|   |                       |   |                   |           |       |   |   |
+   +---+---+---+---+---+   +   +---+---+---+---+   +---+---+   +---+   +   +   +
|       |                   |   |           |                   |           |   |
+---+   +---+---+---+---+   +   +---+---+   +---+---+---+---+   +---+---+   +   +
|   |               |       |   |           |   |   |           |       |       |
+   +---+---+---+   +---+   +   +---+---+   +   +   +---+---+   +---+   +---+   +
|       |   |   |   |       |   |   |                           |   |       |   |
+---+   +   +   +   +---+   +   +   +---+---+---+---+---+---+   +   +---+   +   +
|               |   |                   |   |           |           |   |   |   |
+---+---+---+   +   +---+---+---+---+   +   +---+---+   +---+---+   +   +   +   +
|   |   |       |       |       |       |   |   |   |   |   |       |       |   |
+   +   +---+   +---+   +---+   +---+   +   +   +   +   +   +---+   +---+   +   +
|   |   |       |               |   |       |           |   |   |       |       |
+   +   +---+   +---+---+---+   +   +---+   +---+---+   +   +   +---+   +---+   +
|       |           |   |                                       |   |           |
+---+   +---+---+   +   +---+---+---+---+---+---+---+---+---+   +   +---+---+   +
|       |           |   |               |   |       |               |       |   |
+---+   +---+---+   +   +---+---+---+   +   +---+   +---+---+---+   +---+   +   +
|                                                                               |
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
```

The following code is just a port of the ruby code in the book

