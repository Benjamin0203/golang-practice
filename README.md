# Go Tutorial

---

[Create module](https://go.dev/doc/tutorial/create-module) Creates a `greetings.go` file contains function `Hello(name String)` and in [call module code](https://go.dev/doc/tutorial/call-module-code) we import it into `hello.go`.

---

[Return and handle an error](https://go.dev/doc/tutorial/handle-errors): import `errors` package and return error in `greetings.go` file. In `hello.go` file we handle the error.

---

[Return random greeting](https://go.dev/doc/tutorial/random-greeting): import `math/rand` package and use `rand.Intn` function to return random greeting. In `hello.go` file we call `greetings.Hello` function with random greeting by running `go run .` command in terminal.

---

[get a map with multiple people](https://go.dev/doc/tutorial/greetings-multiple-people): in `greetings.go` file we create `Hellos` function that takes a slice of names and returns a map that associates each of those names with a greeting message. In `hello.go` file we call `greetings.Hellos` function with random greeting by running `go run .` command in terminal.

---

[Testings](https://go.dev/doc/tutorial/add-a-test): Add two testing functions in `greetings_test.go` file. Run `go test` command in terminal to run the tests and see the failed tests only. Run `go test -v` command in terminal to run the tests and see all the tests.

---

[Build](https://go.dev/doc/tutorial/compile-install): Run `go build` command in terminal to build the executable file. Run `./hello` command in terminal to run the executable file.

1. go to the file that contains `func main()` , in this example, our `hello.go` will be the file that needs to be built
2. run command

```bash
go list -f '{{.Target}}'
```

to **get the Go install path**
in this case we get path `/Users/ben/go/bin/hello`

3. Add the**Go install directory** to your system's shell path.
   **Note**: the path has `/bin` as the end and **no** `hello` file!
   Like this:

```bash
export PATH=$PATH:/Users/ben/go/bin
```

4. run `go install`
5. In whereever directory run `hello` command and you will get the same result as running `go run .` inside `/hello` directory:

```bash
─    ~/Fu/golangPractice/hello    main !1 ─────── ✔  01:04:29 am  ─╮
╰─ hello                                                                     ─╯
Great to see you, Gladys!
map[Darrin:Hail, Darrin! Well met! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met!]
```
