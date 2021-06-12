****
Sync
****

Make a counter which is safe to use concurrently.

``sync.WaitGroup``
==================

https://golang.org/pkg/sync/#WaitGroup

A ``WaitGroup`` waits for a collection of goroutines to finish. The main
goroutine calls ``Add()`` to set the number of goroutines to wait for. Then each
fo the goroutines runs and calls ``Done()`` when finished. At the same time,
``Wait()`` can be used to block until all goroutines are finished.

.. code-block:: go

   wantedCount := 1000
   counter := Counter{}

   var wg sync.WaitGroup
   wg.Add(wantedCount)

   for i := 0; i < wantedCount; i++ {
   	go func(w *sync.WaitGroup) {
   		counter.Inc()
   		wg.Done()
   	}(&wg)
   }

   wg.Wait()

Including ``sync.Mutex`` in type
================================

.. code-block:: go

   type Counter struct {
   	mu    sync.Mutex
   	value int
   }

   // vs

   type Counter struct {
   	sync.Mutex
   	value int
   }

In the first case, we do what we've seen in other languages. In the 2nd case,
the methods of ``sync.Mutex`` become part of the public interface of our
``Counter`` type. Hence it should be avoided as it would expose it for abuse by
clients of our type as they can lock/unlock it themselves.

It may make sense in other cases if we're designing our own types.

Do not copy ``sync.Mutex``
==========================

Pass in types which have ``sync.Mutex`` as a data member by pointer as a mutex
must not be copied after first use.

This shows up as an error when we run ``go vet``.
