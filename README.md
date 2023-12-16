<!-- ctrl + shift + v to preview -->
# file-binder

## Table of Contents
- [file-binder](#file-binder)
  - [Table of Contents](#table-of-contents)
  - [General Information](#general-information)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Team \& My Work](#team--my-work)
  - [Main Learnings](#main-learnings)
  - [Setup](#setup)
  - [Example](#example)

## General Information
This project was made as a school project in [kood/Jõhvi](https://kood.tech/) (12.07.2023)

The project required me to create a a file binder which allows users to bind together 2 or more binary executables into a new executable. It works by first creating a mother executable which we append all the other binarys onto. Then, the mother executable reads itself and executes those executable bytes at runtime and printing the result to stdout.

**FILE BINDER** - File binders are utility software that allow a user to "bind" multiple files together, resulting in a single executable. They are commonly used by hackers to insert other programs such as Trojan horses into otherwise harmless files, making them more difficult to detect.

**BINARY EXECUTABLE** - Binary executable files contain executable code that is represented in specific processor instructions. These instructions are executed by a processor directly.

**DISCLAIMER**: This task was for educational purposes only. Do not use any of the methods described to cause harm. I do not take any accountability on any harm caused using these tools.

**NB! Different source control platform was used hence no commit history.**

## Features
- Simple file binder

## Technologies Used

[Golang](https://go.dev/)

## Team & My Work
This was a solo project.

Everything was made by me.

## Main Learnings
- Dangers of file-binders
- Inner workings of a typical file-binder tool

## Setup
Clone the repository
```
git clone https://github.com/JesusKris/file-binder.git
```
Bind 2 binary executables
```
go run injector.go [TARGET NAME] [EXECUTABLE] [EXECUTABLE]... EX: go run injector.go helloworld.exe hello.exe world.exe"
```

## Example

In this example, I will bind 2 simple "Hello, World" programs written in Golang and C.

Golang source:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World from Go file!")
}
```

C source:
```C
#include <stdio.h>
int main() {
   printf("Hello World from C file!\n");
   return 0;
}
```

Next, we need to compile both of those programs into binary executable:

Golang:
```
go build -o ./goExe main.go
```

C:
```
gcc main.c -o cExe
```

Now we will bind those 2 executables into 1 executable called cgoExe:
```
go run injector.go cgoExe goExe cExe
```

![Alt text](./assets/images/result.png?raw=true "result")

Now we can run the executable:
```
./cgoExe
```

![Alt text](./assets/images/result2.png?raw=true "result2")

