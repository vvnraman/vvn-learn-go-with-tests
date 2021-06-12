*****
Maths
*****

Make an SVG of a analog clock.

SVG
===

Clock represented as xml

.. code-block:: xml

   <?xml version="1.0" encoding="UTF-8" standalone="no"?>
   <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
   <svg xmlns="http://www.w3.org/2000/svg"
        width="100%"
        height="100%"
        viewBox="0 0 300 300"
        version="2.0">

     <!-- bezel -->
     <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

     <!-- hour hand -->
     <line x1="150" y1="150" x2="114.150000" y2="132.260000"
           style="fill:none;stroke:#000;stroke-width:7px;"/>

     <!-- minute hand -->
     <line x1="150" y1="150" x2="101.290000" y2="99.730000"
           style="fill:none;stroke:#000;stroke-width:7px;"/>

     <!-- second hand -->
     <line x1="150" y1="150" x2="77.190000" y2="202.900000"
           style="fill:none;stroke:#f00;stroke-width:3px;"/>
   </svg>

Assumptions
===========

- Origin = 0, 0, at top, left hand corner

- Clock centre = 150, 150

- Hour hand = 50 units long

- Minute hand = 80 units long

- Second hand = 90 units long

  Example second hand at midnight
  ::

        0,0
    00  ................................
    01  ................................
    02  ...........150,(150-90).........
    03  ..............^.................
    04  ..............|.................
    05  ..............|.................
    06  ..............|.................
    07  ..............|.................
    08  ...........150,150..............
    09  ................................
    10  ................................
    11  ................................
    12  ................................
    13  ................................
    14  ................................
    15  ................................

floats are horrible
===================

.. code-block:: go

   float64(t.Second()) * (math.Pi / 30)         // does not work
   // vs
   (math.Pi / (30 / (float64(t.Second()))))       // works

main.go
=======

I wasn't able to get ``clockface/main.go`` building as instructed.

::

  $ tree
  .
  ├── clockface
  │   └── main.go
  ├── clockface.go
  └── clockface_test.go

.. code-block:: sh

   cd clockface
   go build


