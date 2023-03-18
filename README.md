# Go Tutorial

---

[Create module](https://go.dev/doc/tutorial/create-module) Creates a `greetings.go` file contains function `Hello(name String)` and in [call module code](https://go.dev/doc/tutorial/call-module-code) we import it into `hello.go`.

---

[Return and handle an error](https://go.dev/doc/tutorial/handle-errors): import `errors` package and return error in `greetings.go` file. In `hello.go` file we handle the error.

---

[Return random greeting](https://go.dev/doc/tutorial/random-greeting): import `math/rand` package and use `rand.Intn` function to return random greeting. In `hello.go` file we call `greetings.Hello` function with random greeting by running `go run .` command in terminal.
