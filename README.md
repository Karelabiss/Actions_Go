# Hardest module you can find on github
### !! I did not use any best practises !!

Debuging status messages for project that don't use best practises.

### Use of this module (Be careful very hard!):

```go
//Prints info status to output
func main(){
  actions.Info("Hello World");
}
```

```
Output: [?] - Hello World
```

It's the same for all functions

```
  .Info(msg) //[?] - msg
  .Debug(msg) //[DEBUG] - msg
  .Failure(msg)  //[-] - msg
  .Error(msg)  //[!] - msg
  .Success(msg)  //[+] - msg
```
