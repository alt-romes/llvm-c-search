## Introduction

This tool scrapes the LLVM-C API and (currently) prints all available API
functions to `stdout`. It came to be from the lack of search options in the
[llvm-c](https://llvm.org/doxygen/group__LLVMC.html) official documentation, and
difficulty to navigate the hyperlinks as a beginner to llvm.

To find, for example, the `sext` corresponding instruction, instead of clicking
all apparently related hyperlinks of sub-modules (in
fact, this exact function was the "last-drop", and the reason why I wrote
program), simply run the program, and then filter for `sext` by tapping `/` when the TUI loads.

## v0.2 Preview

Searching for the function that returns the LLVM pointer type of an LLVM type
![Screenshot 2021-12-24 at 23 39 30](https://user-images.githubusercontent.com/21295306/147374322-b5833e04-1300-4b0f-b152-d2f91e970118.png)
