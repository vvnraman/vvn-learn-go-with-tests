***********
Concurrency
***********

Goroutines
==========

An operation that does not block in golang will run in a separate process called
a **goroutine**. We tell golang to start a new **goroutine** by putting the
keyword ``go`` in front of it.

.. code-block:: go

   go doSomething()

This is often done by using an anonymous function, i.e. one without a name.

.. code-block:: go

   go func(u string) {
   	results[u] = wc(u)
   }(url)

Anonymous functions
===================

- They can be executed at the same time that they're declared, by putting ``()``
  at the end of the anonymous function.

- They maintain access to the lexical scope they are defined in - all the
  variables that are available at the point when you declare the anonymous
  function are also available in the body of the function.

Channels
========

When we start a **goroutine**, the main program doesn't wait for it to complete
before existing. We may unexpectedly either get a panic or we'll have empty data
in any variable we were populating in the **goroutine**.

To solve such issues (and possible data races), we can co-ordinate our
**goroutines** using channels.

.. code-block:: go

   // Send statement
   resultChannel <- result{u, wc(u)}

   // Receive expression
   r := <- resultChannel



