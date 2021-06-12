**********
Reflection
**********

Write the following function, which takes a struct ``x`` and calls ``fn`` for
all string fields found inside.

.. code-block:: go

   walk(x interface{}, fn func(string))

``reflect.ValueOf``
===================

The ``reflect`` package's ``ValueOf()`` method returns us a ``Value`` of a given
variable.

.. code-block:: go

   type Value struct {
   	// contains filtered or unexported fields
   }

``Value`` is the reflection interface to a Go value.

Multiple versions of functions and tests
========================================

We've tried to do something unique with the names of the function and tests here
in order to not destructively refactor code and lose what we had before.

- Use the ``_vX`` suffix to denote a new version of the function or which
  adds add additional functionality. eg. ``walk()`` vs ``walk_v1()``

- Use the ``_rY`` suffix to denote that we've refactored a function or a test
  while retaining functionality. eg.  ``walk()`` vs ``walk_r1()`` vs
  ``walk_r2()`` vs ``walk_v1_r1()`` vs ``walk_v1_r2()``.

- Tests are always specific to a given version. They are agnostic of the
  refactored function.

We're using the following construct in the tests to test mutliple
refactorings (``_rY``) of a function with the same functionality (``_vX``).

.. code-block:: go

   functionsToTest := []func(interface{}, func(string)){
     walk_v5,
     walk_v5_r1,
     walk_v5_r2,
   }


Look at the final ``walk_v8`` for the complete code

.. code-block:: go

    func walk_v8(x interface{}, fn func(input string)
