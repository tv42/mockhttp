=========================================================
 mockhttp.go -- Go package for unit testing HTTP serving
=========================================================

Unit testing HTTP services written in Go means you need to call their
``ServeHTTP`` receiver. For this, you need something that fulfills the
``http.ResponseWriter`` interface, and you need to populate a
``http.Request`` struct with suitable-looking data. ``mockhttp.go``
helps you do these tasks, without excessive copy-pasting.

See ``mockhttp_test.go`` for an example of usage.
