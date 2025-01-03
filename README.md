# Go Sort Algorithms

<img src=".\assets\Go-Sort.png#pic_center" alt="Go-Sort" width="50%" />

## 1 Introduction

This project implements various sorting algorithms in Go. It includes implementations of common sorting algorithms such as Merge Sort, Quick Sort, Heap Sort, and more. The project also includes benchmark tests to compare the performance of these algorithms.

## 2 Features

- [x] **Merge Sort:** An efficient, stable, comparison-based, divide and conquer sorting algorithm.
- [x] **Quick Sort:** A highly efficient sorting algorithm, serving as a systematic method for placing the elements of an array in order.
- [x] **Heap Sort:** A comparison-based sorting technique based on a binary heap data structure.
- [x] **Insertion Sort:** A simple sorting algorithm that builds the final sorted array one item at a time.
- [x] **Selection Sort:** An in-place comparison sorting algorithm.
- [x] **Shell Sort:** An in-place comparison sort that generalizes insertion sort to allow the exchange of items that are far apart.
- [x] **Bucket Sort:** A distribution sort that works by distributing the elements into a number of buckets.
- [x] **Radix Sort:** A non-comparative integer sorting algorithm.
- [x] **Bogo Sort:** A highly ineffective sorting algorithm based on the generate and test paradigm.
- [x] **Bubble Sort:** A simple comparison-based sorting algorithm.
- [x] **Sleep Sort:** A humorous and inefficient sorting algorithm.



## 3 Benchmark

The project includes benchmark tests to measure the performance of each sorting algorithm. These benchmarks can be run using the Go testing framework.

To run the benchmarks:

```sh
go test -bench . -benchmem
```

<img src=".\assets\benchmark.png" alt="Benchmark" />