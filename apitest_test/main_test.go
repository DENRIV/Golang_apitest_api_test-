package main

import (
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/gorilla/mux"
   "github.com/steinfletcher/apitest"
   
   jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func Test_bookmarkIndex(t *testing.T) {
   r := mux.NewRouter()
   r.Handle("/v1/bookmark", bookmarkIndex())
   ts := httptest.NewServer(r)
   defer ts.Close()
   apitest.New().
      Report(apitest.SequenceDiagram()).
      Handler(r).
      Get("/v1/bookmark").
      Expect(t).
      Status(http.StatusOK).
      End()
}

func Test_bookmarkFind(t *testing.T) {
   r := mux.NewRouter()
   r.Handle("/v1/bookmark/{id}", bookmarkFind())
   ts := httptest.NewServer(r)
   defer ts.Close()
   t.Run("not found", func(t *testing.T) {
      apitest.New().
	     Report(apitest.SequenceDiagram()).
         Handler(r).
         Get("/v1/bookmark/1").
         Expect(t).
         Status(http.StatusNotFound).
         End()
   })
   t.Run("found", func(t *testing.T) {
      apitest.New().
	     Report(apitest.SequenceDiagram()).
         Handler(r).
         Get("/v1/bookmark/2").
         Expect(t).
         Assert(jsonpath.Equal(`$.link`, "https://apitest.dev")).
         Status(http.StatusOK).
         End()
   })
}