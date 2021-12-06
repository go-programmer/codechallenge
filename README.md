


“ For me, great algorithms are the poetry of
  computation. Just like verse, they can be terse,
  allusive, dense, and even mysterious.

  But once unlocked, they cast a brilliant new
  light on some aspect of computing. ”
— Francis Sullivan


“ An algorithm must be seen to be believed. ”
— Donald Knuth


“ I will, in fact, claim that the difference
  between a bad programmer and a good one is
  whether he considers his code or his data
  structures more important.
  
  Bad programmers worry about the code.
  
  Good programmers worry about data structures
  and their relationships. ” 
— Linus Torvalds


“ Algorithms + Data Structures = Programs. ” 
— Niklaus Wirth`



# Code Challenge

Code challenge solutions from different sources
in Golang by @codeanit.

I do this because I want to improve my problem 
solving skills and keep my competitive edge.


### How to become good at code challenges?

##### Understand The Basics 
Don't skip basics, mathematics, data structures 
and algorithms. Mathematics helps build a solution. 
The data structures are the tools and the algorithms 
are the techniques that are the arsenal that every
good programmer must have, more the better. Else, 
you will only see `a hammer and a nail`.

##### Know The Process:
To solve the challenge, start with trivial, slow 
ideas to form a heuristic technique, and then 
improve towards creative, fast algorithms which 
could be solved with specific techniques. So just 
solve as you can first even the exponential solution 
if it works it's fine, be greatful.

Start by solving easy problems, then medium, and 
finally the difficult ones. Try different types 
of problems from different sources.

Learn from other's solution and compare with your 
own. Try to understand what other did differently
and analyse what can be improved, both in your 
solution as well as others. This will help add more 
dimensions to problem analysis and solutions ideas. 

Improve your understanding by trying to answer 
Why was it done this way?. 

Review. Realize. Refactor. Re-engineer. 


##### Estimate The Complexity, If Not Provided:
The time limit set for online tests is usually 
from 1 to 10 seconds. We can therefore estimate 
the expected complexity. During contests, we are 
often given a limit on the size of data, and 
therefore we can guess the time complexity within 
which the task should be solved. This is usually 
a great convenience because we can look for a 
solution that works in a specific complexity instead 
of worrying about a faster solution. 
For example, if:
• n <= 1 000 000, the expected time complexity is O(n) or O(nlogn),
• n <= 10 000, the expected time complexity is O(n^2),
• n <= 500, the expected time complexity is O(n^3).

Of course, these limits are not precise. They are 
just approximations, and will vary depending on the 
specific task.



# Folders
`resource` folder contains learning materials.


# Test The Project Via Command Line
Compiles the package sources and tests found in 
the current directory and then runs the resulting 
test binary. 

Without any arguments, also called local directory 
mode. This mode disables caching.

`go test`  

When go tests is invoked with arguments it is called 
package list mode and caches the results.
Test all packages in a directory tree:
`go test ./...` 

Run all tests in the current directory
`go test .`

Run tests for a package
`go test [PACKAGE]`

Run tests verbose (-v)  and stop at the first test that fails (-failfast)
`go test -v -failfast`

Run test of a package passing the package file
`go test FILE.go FILE_test.go`

Run a specific test for a package
`go test -v [PACKAGE] -run [TestName]`

View code coverage
`go test -cover -v`

Create a coverage profile to a file
`go test -coverprofile=coverage.out`

View the generated coverage report coverage.out in browser.
This will open in a browser automatically.
`go tool cover -html=coverage.out`

Skip long running tests (-short)
`go test -short -v`

Test function must have the condition implemented
```go
if testing.Short() {
    t.Skip("Skipping long-running test.")
}
```

Executes the files with a tag `short_test`
`go test -v -tags=sort_test`


go test has two running modes. Understanding them is 
essential to having an easy time working with the tool:
- Local directory mode, or running without arguments.
- Package list mode, or running with arguments.


In the local directory mode, go test compiles the package 
sources and tests found in the current directory and then 
runs the resulting test binary. This mode disables caching. 
After the package test finishes, go test prints a summary 
line showing the test status (‘ok’ or ‘FAIL’), the package 
name, and elapsed time.

To run your tests in this mode, run go test in your project's 
root directory.

In the package list mode, go test compiles and tests each of 
the packages listed as arguments to the command. If a package 
test passes, go test prints only the final ‘ok’ summary line. 
If a package test fails, go test prints the full test output.

To run your test in this mode, run go test with explicit 
package arguments. For example, we can run go test PACKAGE_NAME 
to test a specific package or go test ./... to test all 
packages in a directory tree or go test . to run all tests 
in the current directory.

In our daily work with go test, the difference between 
the two modes is caching.

When we run go test in package list mode it will cache 
successful package test results to avoid unnecessary 
reruns. When it can find the result of a test in the 
cache, go test will redisplay the cached result instead 
of running the tests again. When this happens, go test 
will annotate the test results with (cached) in place 
of the elapsed time in the summary line.

-short in other ways besides skipping a test. For example, 
mocking networks calls instead of opening a connection or 
loading simple fixtures instead of loading them from a 
database. The options are many, it all depends on the 
function that the test is covering.

go test supports tags that are used to run (or skip) a 
special type of test files from our test suite.

There are a couple of rules on build tags.
build tags are special comments, with the format // +build TAGNAME
we have to place the tag on the first line of the file
we have to add an empty line after the tag
the tag name comment cannot have a dash, but it allows underscores

Some other useful features:
    -timeout 120s

    List tests with -list: to list tests, benchmarks, or examples matching a regular expression passed as the argument. The usage of the -list flags is analogous to the -run flag that we discussed before. When we use this flag go test will not run any tests, benchmarks or examples.

    Disabling caching with -count: as we discussed before, go test by default caches test results for packages. It then uses them to skip running tests that have are not modified between two runs. Although can improve the performance, there are times when we would like to disable the caching using the -count=1 flag.

    Get JSON output using -json: if we would like to take the output of go test and process it using a program, having it in a JSON format is better than normal text. This flag will convert the output to JSON, so processing it with a program is less cumbersome.

    Use more CPU cores using -cpu: this flag will set the GOMAXPROCS variable to the argument that we pass to the flag. It limits the number of operating system threads that can execute user-level Go code simultaneously.

    Detect race conditions using -race: since it's very easy to write concurrent code in Go, race conditions are always a risk. Go's tooling is great in this regard - it has a fully integrated race conditions detector in its toolchain.
    https://blog.golang.org/race-detector


t.Errorf() // Test executes after failure
t.Fatalf() // Other test execution halts 