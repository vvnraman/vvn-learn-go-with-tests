**********
Reflection
**********

From twitter - https://twitter.com/peterbourgon/status/1011403901419937792?s=09

Write the following function, which takes a struct ``x`` and calls ``fn`` for
all string fields found inside.

.. code-block:: go

   walk(x interface{}, fn func(string))

``reflect.ValueOf``
===================

The ``reflect`` package's ``ValueOf()`` method returns us a ``Value`` of a given
variable.

