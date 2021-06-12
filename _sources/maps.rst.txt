****
Maps
****

Syntax
======

.. code-block:: go

   type Dictionary map[string]string

Map lookup
==========

.. code-block:: go

   // definition
   var d Dictionary

   // later
   definition, ok := d[word]

The second value is a boolean which indicates if the key was found successfully.

Reference like semantics
========================

Each map value is a pointer to a ``runtime.hmap`` structure. So when we pass
them around, we're copying the internal pointer along. This means that when we
modify a copied map value, the change is reflected outside the scope as well.

``nil`` map
===========

- This creates a ``nil`` map

  .. code-block:: go

     var d Dictionary

- Instead use the following to create an empty map

  .. code-block:: go

     var dictionary = map[string]string{}

     // or

     var dictionary = make(map[string]string)
