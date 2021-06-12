*******
Context
*******

Helps us manage long running, resource intensive processes.

Context tree
============

https://blog.golang.org/context

  The context package provides functdions to derive new Context values from
  existing ones. These values form a tree: when a Context is cancelled, all
  Contexts derived from it are also cancelled.

  - ``Context.WithCancel()``
  - ``Context.WithDeadline()``
  - ``Context.WithTimeout()``
  - ``Context.WithValue()``

.. code-block:: go

   data := "hello, world"
   store := &SpyStore{response: data}
   server := Server_v2(store)

   request := httptest.NewRequest(http.MethodGet, "/", nil)
   response := httptest.NewRecorder()

   cancellingCtx, cancel := context.WithCancel(request.Context())
   time.AfterFunc(5*time.Millisecond, cancel)
   request = request.WithContext(cancellingCtx)


   server.ServeHTTP(response, request)

Where we're doing above is derive a new ``cancellingCtx`` from our ``request``,
which returns us a ``cancel`` function. We then schedule that function to be
called in 5 miliseconds. Finally we use this new context in our request by
calling ``request.WithContext``.

``Context.Done()``
==================

``Context`` has a ``Done()`` method which returns a channel which gets sent a
signal when the context is either "done" or "cancelled". We can listen to that
signal and call ``store.Cancel()`` if we get it.

.. code-block:: go

   ctx := r.Context()

   data := make(chan string, 1)

   go func() {
   	data <- store.Fetch()
   }()

   select {
   case d := <-data:
   	fmt.Fprint(w, d)
   case <-ctx.Done():
   	store.Cancel()
   }

