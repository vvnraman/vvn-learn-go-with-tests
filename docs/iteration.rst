*********
Iteration
*********

This introduces golang's builtin benchmarking capabilities. Its amazing how much
the ecosystem offers by default.

.. code-block:: sh

   go test -bench="."
   # goos: linux
   # goarch: amd64
   # pkg: learn-go-with-tests/src/iteration
   # cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
   # BenchmarkRepeat-16      10653862               109.7 ns/op
   # PASS
   # ok      learn-go-with-tests/src/iteration       1.290s

