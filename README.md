# compare-tool
Compare-tool is a file comparison tool written by golang that compares the difference lines between source and target files. The main feature is reading large files fast benefiting from golang's concurrent features.

## Installation

```bash
go mod tidy
go build
```

## Usage
```golang
./compare-tool -h
Usage of ./compare-tool:
  -dest string
        another input file
  -src string
        one of the input file
  -top N int
        show top N lines (default 10)
```
Input.txt and output.txt is used to compare.
```
./compare -src ./input.txt -dest ./output.txt 
```
The result is below:
```
./input.txt has 42 lines...
./output.txt has 44 lines...
in ./input.txt but not in ./output.txt: 2 lines

diff top 2 lines:
nobody cares about it
something went wrong with you

./output.txt has 44 lines...
./input.txt has 42 lines...
in ./output.txt but not in ./input.txt: 4 lines

diff top 4 lines:
who's the man here.
shell the name
here now
come true with the file
```


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

Licensed under the [GNU General Public License v3.0 or later](https://github.com/Magic-Fy/compare-tool/blob/main/LICENSE).
