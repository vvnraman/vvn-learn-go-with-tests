*****************
Pointers & Errors
*****************

Strongly typed typedefs
=======================

``type MyName OriginalType``

This is one of the amazing features of golang IMHO.

.. code-block:: go

   type Bitcoin int

   func (b Bitcoin) String() string {
   	return fmt.Sprintf("%d BTC", b)
   }

This makes our typedef implement the ``Stringer`` interface from the fmt
package and allows us to define how our type is printed when used with the
``%s`` format string.

Pointers
========

.. code-block:: go

   type Stringer interface {
   	String() string
   }

.. code-block:: go

   func (w *Wallet) Deposit(amount Bitcoin) {
   	w.balance += amount
   }

The ``ReceiverType`` here is a ``pointer`` to a ``Wallet``.

.. code-block:: go

   func (w *Wallet) Deposit(amount Bitcoin) {
   	w.balance += amount
   }

- In ``w.balance += amount``, it should have ideally been ``(*w).balance``, but
  these pointers to structs are automatically dereferenced.

``nil``
=======

Pointers can be ``nil`` and must be checked for the same. Otherwise it
will ``panic``.

Errors
======

Use the ``errors`` package.

.. code-block:: go

   var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

