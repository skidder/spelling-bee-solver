## Spelling Bee Solver
Simple program to help me solve the NY Times Spelling Bee puzzle because we all need a little help some times. Uses a word list from [https://github.com/dwyl/english-words](https://github.com/dwyl/english-words).

### Installing
```
go get github.com/skidder/spelling-bee-solver
```

Then download a world file, either from this repo (`words.txt`), from [https://github.com/dwyl/english-words](https://github.com/dwyl/english-words), or your own.

### Running

```
spelling-bee-solver -availableLetters dineupz -requiredLetter u
```

### Flags
```
  -availableLetters string
        Letters that can be used
  -minLength int
        Minimum length of words (default 4)
  -requiredLetter string
        Letters that must be used
  -words string
        Path to words file (default "words.txt")
```
