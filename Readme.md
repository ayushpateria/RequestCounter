## Request Counter ##

The program demonstrates a basic server which counts the total number of requests received. To support multiple threads operate on the same counter, all the operations on it are thread-safe. As atomic operations typically incur some extra costs, the server is optimized for accuracy over performance. I think for a request counter it makes more sense to optimize for accuracy, otherwise, the actual count could be different by a factor of number of threads.

The QPS is calculated by a go-routine that captures the value of counter every second.

*Instructions*

- Build using `make build`.
- Run the server by executing `./bin/myserver'
- You should be able to reach the server at http://localhost:8080/ and see the number of requests at http://localhost:8080/count and the qps at http://localhost:8080/qps
