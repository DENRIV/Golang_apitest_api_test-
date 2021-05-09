# Golang_apitest_api_test

Golang_apitest_api_test 


> Using apitest

https://apitest.dev/

use some third-party libraries, such as apitest

installing the new packages

go get github.com/steinfletcher/apitest

go get github.com/steinfletcher/apitest-jsonpath

[run]

go run main.go

  localhost:8080 
  
Similar to last main_test.go, using apitest :

main_test.go

[CODE][Changed]

   jsonpath "github.com/steinfletcher/apitest-jsonpath"

   apitest.New().
   
	//..
  
   t.Run("not found", func(t *testing.T) {
   
      apitest.New().
      
	//..
  
   t.Run("found", func(t *testing.T) {
   
      apitest.New().
      
	//..


> Reports

These are the generated reports:

reports generation

add:

	Report(apitest.SequenceDiagram())
  
example:

apitest.New().

   Report(apitest.SequenceDiagram()).
   
   Handler(r).
   
   Get("/v1/bookmark").
   
   Expect(t).
   
   Status(http.StatusOK).
   
   End()


[run]:

go test -v

show in screen

=== RUN   Test_bookmarkIndex

Created sequence diagram (3157381659_2166136261.html): D:\@COURSES\GO\@GO_API\golang_api_test_apitest\apitest_test\.sequence\3157381659_2166136261.html

--- PASS: Test_bookmarkIndex (0.01s)

=== RUN   Test_bookmarkFind

=== RUN   Test_bookmarkFind/not_found

Created sequence diagram (1543772695_2166136261.html): D:\@COURSES\GO\@GO_API\golang_api_test_apitest\apitest_test\.sequence\1543772695_2166136261.html

=== RUN   Test_bookmarkFind/found

Created sequence diagram (1560550314_2166136261.html): D:\@COURSES\GO\@GO_API\golang_api_test_apitest\apitest_test\.sequence\1560550314_2166136261.html

--- PASS: Test_bookmarkFind (0.03s)

    --- PASS: Test_bookmarkFind/not_found (0.00s)
    
    --- PASS: Test_bookmarkFind/found (0.03s)

PASS

ok      github.com/eminetto/post-apitest        0.963s

CONCLUSION :

-  Using only the standard library, the project gains speed in the execution of tests, besides not depending on third-party libraries, which can be a problem in some teams.

- With a library like apitest you gain in productivity and ease of maintenance, but you lose in speed of execution. 

   MAYBE NO big is the difference compared to the standard library





