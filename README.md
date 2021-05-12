# rcal cli

The CLI allows you to analyze rectangles and features that exist among rectangles. Your implementation is required to cover
the following:

- Intersection
- Adjacency
- Containment

```text
rcal cli calculates intersection, containment and adjacency between two rectangles.

Usage:
  rcal [command]

Available Commands:
  adjacency    calculates the adjacency between two rectangles
  containment  calculates containment between two rectangles.
  help         Help about any command
  intersection calculates the intersection between two rectangles.

Flags:
      --config string   config file (default is $HOME/.rcal.yaml)
  -h, --help            help for rcal
  -t, --toggle          Help message for toggle

Use "rcal [command] --help" for more information about a command.
```

## Intersection

```text
$ rcal intersection
Error: required flag(s) "rectangleA", "rectangleB" not set
Usage:
  rcal intersection [flags]

Flags:
  -h, --help                help for intersection
  -a, --rectangleA string   rectangle A coords BottomLeft (x,y) and Top Right (x,y) as -a 6,4,14,10 -b  4,2,7,6
  -b, --rectangleB string   rectangle B coords BottomLeft (x,y) and Top Right (x,y) as -a 6,4,14,10 -b  4,2,7,6

Global Flags:
      --config string   config file (default is $HOME/.rcal.yaml)

Error: required flag(s) "rectangleA", "rectangleB" not set

```

### Command
```text
$rcal intersection -a 6,4,14,10 -b 4,2,7,6
```
### Output

```text
 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {6 10}         {14 10}         {6 4}           {14 4}          6       14      4       10      

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {4 6}          {7 6}           {4 2}           {7 2}           4       7       2       6       

+--------Reference--------+
+-------- Intersection in Block #1 --------+
        +-----------+
        |        R1 |
+-----------+       |
| R2        |       |
|           |-------+
+-----------+ 


*** Intersection between Rectangle A and Rectanle b ***
        Intersection in Block No. 1 : [{6 6} {7 4}]

```

### Command

```text
$ rcal intersection -a 6,4,14,10 -b 2,5,4,8
```

### Output

```text
 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {6 10}         {14 10}         {6 4}           {14 4}          6       14      4       10      

 TopLeft        TopRight        BottomLeft      BottomRight     xMin    xMax    yMin    yMax    
 --------       --------        --------        --------        --------------------------------
 {2 8}          {4 8}           {2 5}           {4 5}           2       4       5       8       


*** Intersection between Rectangle A and Rectanle b: None Identified ***

```
