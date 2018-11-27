# hello-go
My first github respository is a golang hello world program. Nothing interesting in the example.
What is interesting is the apparent ease of cross compiling using go on Linux to target macOS and Windows.
Or cross compiling using Windows Subsystem for Linux.  This looks very attractive indeed:

```
GOOS=windows GOARCH=amd64 go build
```

