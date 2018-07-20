Benchmark results
=================

This file contains the output of example/example.go run on an i7 3770k with 16GB of RAM.

Current results
---------------

Radix sorted  10  in  10.167µs
Sort sorted  10  in  2.22µs
Quicksort sorted  10  in  1.624µs
Insertion sorted  10  in  476ns
Merge sorted  10  in  890ns
Success:
Radix:  true
QuickSort:  true
Insertion:  true
Merge:  true

Radix sorted  100  in  6.517µs
Sort sorted  100  in  15.09µs
Quicksort sorted  100  in  10.403µs
Insertion sorted  100  in  11.817µs
Merge sorted  100  in  22.652µs
Success:
Radix:  true
QuickSort:  true
Insertion:  true
Merge:  true

Radix sorted  1000  in  64.092µs
Sort sorted  1000  in  178.735µs
Quicksort sorted  1000  in  105.591µs
Insertion sorted  1000  in  582.955µs
Merge sorted  1000  in  187.103µs
Success:
Radix:  true
QuickSort:  true
Insertion:  true
Merge:  true

Radix sorted  10000  in  583.762µs
Sort sorted  10000  in  2.327078ms
Quicksort sorted  10000  in  1.243841ms
Insertion sorted  10000  in  34.99612ms
Merge sorted  10000  in  1.210353ms
Success:
Radix:  true
QuickSort:  true
Insertion:  true
Merge:  true

Radix sorted  100000  in  3.553859ms
Sort sorted  100000  in  18.349581ms
Quicksort sorted  100000  in  7.81994ms
Insertion sorted  100000  in  3.234176203s
Merge sorted  100000  in  19.119832ms
Success:
Radix:  true
QuickSort:  true
Insertion:  true
Merge:  true

Radix sorted  1000000  in  47.965834ms
Sort sorted  1000000  in  184.037864ms
Quicksort sorted  1000000  in  93.071287ms
Insertion sort skipped due to high set size
Merge sorted  1000000  in  163.817966ms
Success:
Radix:  true
QuickSort:  true
Merge:  true

Radix sorted  10000000  in  449.192062ms
Sort sorted  10000000  in  2.132433473s
Quicksort sorted  10000000  in  1.055007709s
Insertion sort skipped due to high set size
Merge sorted  10000000  in  1.644670057s
Success:
Radix:  true
QuickSort:  true
Merge:  true

Radix sorted  100000000  in  4.575019916s
Sort sorted  100000000  in  24.310221216s
Quicksort sorted  100000000  in  11.906609603s
Insertion sort skipped due to high set size
Merge sorted  100000000  in  17.889905109s
Success:
Radix:  true
QuickSort:  true
Merge:  true

Later tests with Radix sort
---------------------------

Radix sorted  10  in  13.319µs
Sort sorted  10  in  2.353µs
Quicksort sorted  10  in  19.316µs
Success:
Radix:  true
QuickSort:  true

Radix sorted  100  in  18.454µs
Sort sorted  100  in  18.288µs
Quicksort sorted  100  in  12.535µs
Success:
Radix:  true
QuickSort:  true

Radix sorted  1000  in  85.955µs
Sort sorted  1000  in  232.067µs
Quicksort sorted  1000  in  140.521µs
Success:
Radix:  true
QuickSort:  true

Radix sorted  10000  in  772.48µs
Sort sorted  10000  in  2.473692ms
Quicksort sorted  10000  in  1.212938ms
Success:
Radix:  true
QuickSort:  true

Radix sorted  100000  in  5.012217ms
Sort sorted  100000  in  21.121836ms
Quicksort sorted  100000  in  8.046821ms
Success:
Radix:  true
QuickSort:  true

Radix sorted  1000000  in  46.740532ms
Sort sorted  1000000  in  187.233799ms
Quicksort sorted  1000000  in  95.314782ms
Success:
Radix:  true
QuickSort:  true

Radix sorted  10000000  in  462.379556ms
Sort sorted  10000000  in  2.147001491s
Quicksort sorted  10000000  in  1.079043131s
Success:
Radix:  true
QuickSort:  true

Radix sorted  100000000  in  4.606175081s
Sort sorted  100000000  in  24.272163308s
Quicksort sorted  100000000  in  12.093175678s
Success:
Radix:  true
QuickSort:  true

First tests of Quicksort
------------------------

Sort sorted  10  in  1.232µs
Quicksort sorted  10  in  885ns

Sort sorted  100  in  7.723µs
Quicksort sorted  100  in  5.462µs

Sort sorted  1000  in  96.258µs
Quicksort sorted  1000  in  61.588µs

Sort sorted  10000  in  1.256535ms
Quicksort sorted  10000  in  739.672µs

Sort sorted  100000  in  15.43422ms
Quicksort sorted  100000  in  8.831228ms

Sort sorted  1000000  in  190.096395ms
Quicksort sorted  1000000  in  104.110686ms

Sort sorted  10000000  in  2.139816277s
Quicksort sorted  10000000  in  1.181530688s

Sort sorted  100000000  in  24.435276898s
Quicksort sorted  100000000  in  13.315615248s

Sort sorted  1000000000  in  11m32.177990751s
Quicksort sorted  1000000000  in  6m18.776106009s