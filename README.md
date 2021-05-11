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
  