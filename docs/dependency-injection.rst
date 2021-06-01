********************
Dependency Injection
********************

The ``fmp.printf`` calls ``Fprintf`` with ``os.Stdout``, which is an
``io.Writer``

.. code-block:: go

   type Writer interface {
   	Write(p []byte) (n int, err error)
   }

Apart from ``os.Stdout`` this interface is also implemented by the
``bytes.Buffer`` type, which we can use to write test for something which
implements the ``io.Writer`` interface.


