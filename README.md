GoMonday
========

Golang API connection examples with Monday.com v2 API based on GraphQL.

Originally created by Patrick Haugen at Production Media Design under SoftAngle w/ Trust Encryption for use at Disney Studios and for other companies seeking Monday.com connection assistance.

Building
--------

Ensure Golang is installed on your system:
1. https://golang.org/dl/
2. > go version

Windows:
> go build -o gomonday.exe && gomonday.exe -verbose -key [apikey] -board [boardid]

Linux:
> go build -o gomonday && gomonday -verbose -key [apikey] -board [boardid]

Running
-------

Windows:
> gomonday.exe
> gomonday.exe -h
> gomonday.exe -version
> gomonday.exe -verbose -debug -webport 80 -key [apikey] -board [boardid]

Linux:
> gomonday.exe
> gomonday.exe -h
> gomonday.exe -version
> gomonday.exe -verbose -debug -webport 80 -key [apikey] -board [boardid]
