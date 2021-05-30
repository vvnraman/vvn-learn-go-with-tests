*****************
Arrays and slices
*****************

Array initialization
====================

Arrays have a fixed capacity which we define when we declare the variable. There
are two ways to initialize:

- ``[N]type{value1, value2, ..., valueN}``

  .. code-block:: go

     numbers := [5]{1, 2, 3, 4, 5}

- ``[...]type{value1, value2, ..., valuen}``

  .. code-block:: go

     numbers := [...]int{1, 2, 3, 4, 5}

- Their type encodes the size within itself.

----

- Print format ``%v``

``range``
=========

``range`` lets us iterate over an array. Every time it is called, it returns two
values, the index and the value.

.. code-block:: go

   sum := 0
   for _, number := range numbers {
       sum += number
   }

We can ignore the index by using ``_``.

Slices
======

- Slice type ``[]type{value1, value2, ..., valuen}``

  .. code-block:: go

     sliceNumbers := []int{1, 2, 3, 4, 5}

Variable number of arguments
============================

Example

.. code-block:: go

   func SumAll(numbersToSum ...[]int) (sums [] int) {
       return
   }

A method which taken in a variable number of slices and returns a slice.

