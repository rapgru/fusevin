# fusevin

Fusevin is a FUSE filesystem that provides special files to notify you about read access and allows you to choose the returned bytes of the read syscall over gRPC.

## Why Fusevin?

fusevin was created to address the challenge of monitoring programs that read input from standard input (stdin). The difficulty arises when it's essential to detect when input is required by the program.
This problem emerged in the context of a Jupyter kernel for a compiled language, where the kernel needs to explicitly signal the frontend to enable user input for stdin.
While fusevin is an unconventional approach to a very specific problem, it provides a highly effective solution. However, there may be simpler alternatives that exist.

Regardless, fusevin works and its potential use cases are yet to be fully explored.
