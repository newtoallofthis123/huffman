# Huffman

Huffman encoding library implemented in Go

## Installation

This is a native implementation of the Huffman encoding algorithm in Go. To install it, you can use the `go get` command:

```bash
go get github.com/newtoallofthis123/huffman
```

## Example

```go
package main

func main(){
 chars := []rune{'A', 'A', 'A', 'B', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D', 'E', 'E'}
 // This is needed since the library is implemented on generic slices
 interfaces := huffman.ConvertToInterface(chars)

 h := huffman.FromList(interfaces)
 res := h.Encode()
 r := h.Decode(res)

 fmt.Println(r, res)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
