# Cpu Troll
Chaos Troll dedicated to raising CPU latency by the requested percentage and timespan.

This is part of the larger Chaos Troll Suite meant to provide FaaS (Failure as a Service) to improve the ease of adopting a Chaos Engineering model.

# Usage
Run the executable in any prompt and pass in the cpu usage percentage out of 100 and the length of time in seconds for the latency to last.

`./cpu-troll.exe 100 15` will run the CPU at 100% load for 15 seconds.

# Dependencies
Golang - https://golang.org/

# To Build Source
Ensure that Go is properly setup on your machine. [Golang -- Getting Started](https://golang.org/doc/install)

Clone the package and navigate to the directory it was cloned to in any prompt.

Run

`go build`

and the executable should be built and placed in your current working directory.