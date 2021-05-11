# rcal cli

Calculates the following features between two rectangles R1 and R2

- Intersection
- Containment
- Adjacency

## Commands

### Intersection

```text
$ rcal intersection -a 6,4,14,10 -b  4,2,7,6

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {6 10}         {14 10}         {6 4}           {14 4}          6       14      4       10      

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {4 6}          {7 6}           {4 2}           {7 2}           4       7       2       6       
scanning for intersection(s) @ block: 1

+--------Reference--------+
        +-----------+
        |        R1 |
+-----------+       |
| R2        |       |
|           |-------+
+-----------+ 
Points of Intersection in Block No. 1 : [{6 6} {7 4}] 
```

### Containment

```text
$ rcal containment -a 6,4,14,10 -b 7,6,11,8

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {6 10}         {14 10}         {6 4}           {14 4}          6       14      4       10      

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {7 8}          {11 8}          {7 6}           {11 6}          7       11      6       8       

+--------Reference--------+
+---------------+
|   +-------+   |
|   |   R2  |   |
|   +-------+   |
+---------------+ 
```

### Adjacency

```text
$  rcal adjacency -a 6,4,14,10 -b 8,10,16,4

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {6 10}         {14 10}         {6 4}           {14 4}          6       14      4       10      

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {8 4}          {16 4}          {8 10}          {16 10}         8       16      10      4       

+--------Reference--------+
Adjacency (Partial)
+-------+
| R1    |
|       |
|       |
+----+-------+
     | R2    |
     |       |
     |       |
     +-------+ 
```

