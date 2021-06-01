*****************************
Structs, Methods & Interfaces
*****************************

Structs
=======

A struct is just a named collection of fields where we can store data.

.. code-block:: go

   type Rectangle struct {
   	Width  float64
   	Height float64
   }

   type Circle struct {
   	Radius float64
   }

If a symbol name (i.e. variables, types, functions, etc..) starts with a
lowercase, then it is private outside the package it's defined in.

Methods
=======

A method is a function with a receiver. A method declaraion binds an identifier,
the method name, to a method, and associates the method with the receiver's base
type.

``func (receiverName ReceiverType) MethodName(args) returnType {}``

They are very similar to functions but they are called by invoking them on an
instance of a particular type.

.. code-block:: go

   func (r Rectangle) Area() float64 {
   	return r.Width * r.Height
   }

By default, golang copies values when we pass in the ``ReceiverType`` while
calling the method.

Interfaces
==========

.. code-block:: go

   type Shape interface {
   	Area() float64
   }

In golang, **interface resolution is explicit**. If the type we have matches
what the interface is asking for, it will compile.

Anonymous structs
=================

.. code-block:: go

   	areaTests := []struct {
   		shape Shape
   		want  float64
   	}{
   		{Rectangle{12, 6}, 72.0},
   		{Circle{10}, 314.1592653589793},
   		{Triangle{12, 6}, 36.0},
   	}

Here we're declaring a slice of structs by using ``[]struct`` with two fields,
the ``shape`` and the ``want``.

Table driven tests and ``t.Run()``
==================================

We're using anonymous structs to capture test data, and then drive the tests
from that.

.. code-block:: go

   func TestArea4(t *testing.T) {
   	areaTests := []struct {
   		name    string
   		shape   Shape
   		hasArea float64
   	}{
   		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
   		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
   		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
   	}

   	for _, tt := range areaTests {
   		t.Run(tt.name, func(t *testing.T) {
   			got := tt.shape.Area()
   			if got != tt.hasArea {
   				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
   			}
   		})
   	}
   }

Here we're also using ``t.Run()`` within a for loop to have a more descriptive
output in case of failures as it shows separately as a test case within the
test.

.. code-block:: go

   	for _, tt := range areaTests {
   		t.Run(tt.name, func(t *testing.T) {

This also allows us to run a specific test within

.. code-block:: sh

   go test -run TestArea4/Rectangle
