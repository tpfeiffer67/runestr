# What is runestr?

Runestr is a library designed for manipulating strings at the rune level.

## Functions

- `Length`: Returns the count of Unicode code points (runes) within a string.
- `PadRight`: Pads a string on the right side with a specified fill pattern until it reaches the provided width.
- `Left`: Retrieves the leftmost n runes from a string.
- `Right`: Retrieves the rightmost n runes from a string.
- `SplitOnNearestSpace`: Divides a string into two parts at the nearest space character to the n-th position.
- `RuneAtPosition`: Fetches the Unicode code point (rune) at the specified position in a string or returns the specified default value if the position is out of bounds.

## Usage

```golang
import (
    "github.com/tpfeiffer67/runestr"
)
```

## License

This repository is licensed under the MIT License. See the LICENSE file for more information.


## Author

- [@TPfeiffer67](https://www.github.com/TPfeiffer67)
