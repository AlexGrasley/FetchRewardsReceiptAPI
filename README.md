## READ ME

### Author
Alex Grasley

alexgrasley@gmail.com

### Description
Sample API written in Go to fulfill the requirements of the Fetch Rewards coding challenge. 

https://github.com/fetch-rewards/receipt-processor-challenge

My background is in C# and .NET API development. It was a lot of fun learning some of the basics of Go, 
and I'm excited to continue learning more about the language and its ecosystem. I wrote and structured the code
in a way that made sense to me, but might not be idiomatic Go. I am open to any and all feedback so that I can learn
to do better next time! 

### Running the API
There should be no external dependencies to run the API. The only requirement is to have Go installed on your machine.

#### Endpoints 
- POST /receipts/process
  - returns the Id of the receipt that was created
- GET /receipts/{id}/points
  - returns the number of points for the receipt with the given Id

### Testing
The project contains a set of basic unit tests that should validate most of the core logic of the API.

### Final Notes
While the API is simple, there are still things that could, and would need to be improved in a production environment.
Improvements include: Security, logging, error handling, performance, and more robust testing.

Security: Need to include some form of authentication and authorization, likely in the form of JWT tokens.

Logging: Need to include logging to help with debugging and monitoring.

Error Handling: Need to include more robust error handling, and return more informative error messages to the client,
assuming that this API is not public facing and only being accessed by a trusted client. If it were public facing,
we would need to be more careful about what information we return in the error messages.

Performance: The current implementation is plenty fast for the test cases, but if we were to scale this up to a 
production environment that is being accessed by a large number of clients as once, we would need to ensure that
everything is optimized. This could include the use of caching, re-use of object instances (such as items), background 
workers for longer running tasks, such as calculating points on a large receipt, and adding multi-threading when
calculating points, or performing other necessary processing on receipts. The current implementation by no means 
"needs" concurrency for point calculations, but it could be useful for receipts with a large number of items. 

Testing: The current tests are very basic and only cover the core logic of the API. Ideally we would include full integration 
tests as well as additional tests for edge cases and error conditions.