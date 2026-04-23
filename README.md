# ascii-art

# ASCII Art Generator (Go)

## Description

This project is a command-line program written in **Go** that converts input text into ASCII art using predefined banner templates.

Each character is represented using an 8-line graphical format, allowing you to display styled text directly in your terminal.

## Features

* Converts standard ASCII characters (32–126) into ASCII art
* Supports multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
    
* Handles:

  * Line breaks (`\n`)
  * Empty input
  * Invalid characters (replaced with space)
* Clean modular structure with separate functions
* Unit-test friendly design

## Project Structure

```
ascii-art/
│
├── main.go              # Entry point
├── buildAsciiMap.go     # Builds ASCII character map
├── printAscii.go        # Prints ASCII output
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── tests/
    ├── main_test.go
    ├── ascii_map_test.go
    └── print_ascii_test.go
```


## How It Works

1. The program reads a banner file.
2. Each ASCII character is mapped to its 8-line representation.
3. The input string is processed:

   * `\n` is converted into actual new lines
4. The program prints characters line-by-line to form ASCII art.


## Banner Format

* Each character has **8 lines**
* Characters are separated by **1 empty line**
* Total characters: **95** (ASCII 32–126)

Example (simplified):

```
(space)

        
        
        
        
        
        

!

 _ 
| |
| |
| |
|_|
(_)
    
    

"
 _ _  
( | ) 
 V V  
      
      
      
      
      
```

## Usage

### Basic Run

```bash
go run . "Hello"
```

### With Specific Banner

```bash
go run . "Hello" shadow.txt
```

### With New Lines

```bash
go run . "Hello\nWorld"
```

---

## Examples

```bash
go run . "Hello"
```

```bash
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Error Handling

The program handles:

* Invalid number of arguments
* Invalid banner name
* File read errors
* Unsupported ASCII characters


## Testing

Run tests using:

```bash
go test ./...
```

Test coverage includes:

* ASCII map generation
* Output rendering
* Edge cases (empty input, invalid input)


## Improvements (Future Work)

* Return errors instead of using `os.Exit`
* Add support for custom banner files
* Improve performance for large inputs
* Add file input/output support


## Developed as part of a Go learning project focused on:

* File handling
* String manipulation
* Data structures (maps, slices)
* Clean code practices

## License

This project is for educational purposes.
