# rcal cli

Calculates the following features between two rectangles R1 and R2

- Intersection
- Containment
- Adjacency

## Commands

### Intersection

```text
 rcal intersection -a 4,2,7,6 -b 6,4,14,10

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {7 2}          {4 2}           {7 6}           {4 6}           7       4       6       2       

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {14 4}         {6 4}           {14 10}         {6 10}          14      6       10      4  
```

### Containment

```text
rcal containment -a 4,2,7,6 -b 6,4,14,10

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {7 2}          {4 2}           {7 6}           {4 6}           7       4       6       2       

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {14 4}         {6 4}           {14 10}         {6 10}          14      6       10      4 
```

### Adjacency

```text
rcal adjacency -a 4,2,7,6 -b 6,4,14,10

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {7 2}          {4 2}           {7 6}           {4 6}           7       4       6       2       

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {14 4}         {6 4}           {14 10}         {6 10}          14      6       10      4 
```

