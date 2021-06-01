******
Select
******

``httptest.Server``
===================

Go has a built-in ``net/http`` package for building web servers as well as an
``net/http/httptest`` which implements all the interfaces in the ``net/http``
package.

Here is how we create a test server

.. code-block:: go

   httptest.NewServer(http.HandlerFunc(
   	func(w http.ResponseWriter, r *http.Request) {
   		w.WriteHeader(http.StatusOK)
   	})
   )

``defer``
=========

By prefixing a function call with ``defer`` golang will now call that function
at the end of the containing function.

Always make channels
====================

The zero value for channels is ``nil``, and if we try to send to it via ``<-``,
it will block forever by design.

So we cannot use ``var ch chan struct{}`` to create a channel, and we should
always use ``ch := make(chan struct{})``

.. code-block:: go

   func ping(url string) chan struct{} {
   	ch := make(chan struct{})
   	go func() {
   		http.Get(url)
   		close(ch)
   	}()
   	return ch
   }

``select-case``
===============

.. code-block:: go

   select {
   case <-ping(slow):
   	return slow, nil
   case <-ping(fast):
   	return fast, nil
   case <-time.After(timeout):
   	return "", fmt.Errorf("timed out waiting for %s and %s", slow, fast)
   }

- We can wait on multiple channels via ``select``

- The first one to sent a value "wins" and the code under the ``case`` executes.

``time.After``
==============

``time.After()`` returns a ``chan`` and will send a signal down it after the
amount of time we define.
