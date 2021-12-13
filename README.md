## Introduction

This tool scrapes the LLVM-C API and (currently) prints all available API
functions to `stdout`. It came to be from the lack of search options in the
[llvm-c](https://llvm.org/doxygen/group__LLVMC.html) official documentation, and
difficulty to navigate the hyperlinks as a beginner to llvm.

To find, for example, the `sext` corresponding instruction, instead of clicking
all apparently related hyperlinks of sub-modules, simply run the following (in
fact, this exact function was the "last-drop", and the reason why I wrote
program)
```
grep -i "sext" version_0.1.txt
```

---

## v0.1 Preview

```
Bit Writer
----------
(Functions)
int LLVMWriteBitcodeToFile (LLVMModuleRef M, const char *Path)
Writes a module to the specified path.  More...
int LLVMWriteBitcodeToFD (LLVMModuleRef M, int FD, int ShouldClose, int Unbuffered)
Writes a module to an open file descriptor.  More...
int LLVMWriteBitcodeToFileHandle (LLVMModuleRef M, int Handle)
Deprecated for LLVMWriteBitcodeToFD.  More...
LLVMMemoryBufferRef LLVMWriteBitcodeToMemoryBuffer (LLVMModuleRef M)
Writes a module to a new memory buffer and returns it.  More...

Transforms
----------
Aggressive Instruction Combining transformations
------------------------------------------------
(Functions)
void LLVMAddAggressiveInstCombinerPass (LLVMPassManagerRef PM)
See llvm::createAggressiveInstCombinerPass function.  More...

Coroutine transformations
-------------------------
(Functions)
void LLVMAddCoroEarlyPass (LLVMPassManagerRef PM)
See llvm::createCoroEarlyLegacyPass function.  More...
```
