# ASCII Art 

A **Go-based ASCII Art Generator** that converts input text into ASCII art using customizable font files. This project demonstrates modular programming in Go and handles edge cases like newlines, empty input, and multiple font styles.


## Features

* Convert any text into ASCII art using a **selected font** (`shadow.txt`, `standard.txt`, `thinkertoy.txt`).
* Handles multi-line input using `\n`.
* Handles empty input and single newline inputs.
* Modular Go design:
  * `main.go` – orchestrates the program
  * `readfont.go` – reads the ASCII font file
  * `parseinput.go` – processes input into words/lines
  * `parseascii.go` – loops through words and prints ASCII
  * `printword.go` – prints each word character by character


## File Structure

```bash
ascii-art-project/
├── main.go
├── asciiart/
│   ├── readfont.go
│   ├── parseinput.go
│   ├── parseascii.go
│   └── printword.go
├── standard.txt
├── thinkertoy.txt
├── shadow.txt
└── README.md
```

## Usage

### Run the program

By default, the program reads `standard.txt`:

```bash
go run . "Hello"
```

You can switch fonts by changing the file in `ReadFont` in `main.go`:

```go
fileLines, err := asciiart.ReadFont("thinkertoy.txt")
```

or

```go
fileLines, err := asciiart.ReadFont("shadow.txt")
```

## Examples (Standard Font)

1. **Single word:**

```bash
go run . "Hello" | cat -e
```

Output:

```text
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

2. **Multi-line input using `\n`:**

```bash
go run . "Hello\nThere" | cat -e
```

Output:

```text
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

3. **Empty input:**

```bash
go run . "" | cat -e
```

Output:

```text

```

4. **Single newline input:**

```bash
go run . "\n" | cat -e
```

Output:

```text
$
```


## How It Works

1. `main.go` reads the input from the command line.
2. `readfont.go` loads the selected ASCII font file and splits it into lines.
3. `parseinput.go` converts the input string into words/lines, handling literal `\n`.
4. `parseascii.go` loops through each word and prints its ASCII representation or a blank line if empty.
5. `printword.go` prints each character row by row using the font file.


## ASCII Font File

* Each font file contains **all printable ASCII characters**.
* Each character is **8 lines tall + 1 empty line**, forming a **9-line block**.
* The program maps each character to the correct block using:

```go
asciiIndex := int(char) - 32
start := asciiIndex * 9
```

* Changing the font file changes the **style of ASCII art** output.

## Project Goals

* Learn modular programming in Go
* Practice file I/O and string manipulation
* Handle edge cases like empty input and newlines
* Build a reusable ASCII Art generator with multiple font styles


