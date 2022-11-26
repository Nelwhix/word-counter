# Word-Counter
    A CLI tool that can count words, lines or bytes
    in a string or file

## Usage
    To count words in a string

```bash
    echo "Hello I am a boy" | ./word-counter
```
    Expected answer should be 5. Or if you have a file write:

```bash
    cat test.txt | ./word-counter
```
    Expected answer should be 64.

    To Count Lines, add the "-lines" flag to your command.
    So for the previous examples we have:
```
    echo "Hello I am a boy" | ./word-counter -lines
```

To Count Bytes, add the "-bytes" flag to your command

## Testing
    You must have go version 1.19.3 installed
```bash
    go test -v
```

## Builds
    Windows and MacOS binaries are on this repo so download and enjoy!