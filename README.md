# parallel-sorting-go

A training project built to practice Go by reading numbers from files, sorting them, and exploring how to process data in a structured and scalable way.

This project was inspired by another assignment from an Operating Systems course that originally had to be implemented in C, but here the idea is recreated in Go as a learning exercise.

## Project goals

The main goals of this project are:

- Practice the Go language with a real programming exercise. 
- Read numeric data from files.
- Sort the data efficiently. 
- Understand how a larger sorting task can be divided into smaller parts. 
- Experiment with concepts related to parallelism and external sorting. 

## How it works

The project reads numbers from input files, splits the data into smaller chunks, sorts each chunk, and then merges the sorted parts into a final ordered result. This structure is useful for studying sorting strategies and for understanding how the problem could later be adapted to parallel execution. 

## Tech stack

- Go
- File I/O
- Sorting algorithms
- Heap / Min-Heap for merging sorted chunks

## Project structure

A possible structure for the project is:

```text
parallel-sorting-go/
├── main.go
├── input/
├── output/
└── README.md
```

## Running the project

Make sure you have Go installed, then run:

```bash
go run main.go
```

If your project is organized into multiple files, you can run:

```bash
go run .
```

## Learning focus

This project is mainly focused on learning:

- Go syntax and project organization
- Reading and writing files
- Implementing sorting algorithms
- Working with heaps
- Thinking about concurrency and parallel sorting strategies

## Possible improvements

Some future improvements for the project could be:

- Add goroutines to sort chunks in parallel
- Benchmark different sorting strategies
- Support very large input files
- Add automated tests
- Improve input validation and error handling

## Notes

This is a study project, so the main purpose is to understand the ideas behind the implementation rather than build a production-ready tool. 