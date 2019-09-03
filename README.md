# Demo Parrot Service
This is a basic HTTP API that will be used to demonstrate Helm and Kind.
There is one endpoint with two methods:
- GET /parrots
- POST /parrots
To run, ``````go build main.go
``This requires MongoDB as that is where it stores the parrot data, and also I needed an external dependency ^^;
After the demo, feel free to poke around with it. This was built with clean architecture in mind, which I haven't done in a while.

