# simple HTTP server in GO (from scratch)
a simple multithreaded HTTP server from scratch with go

HTTP/1.1 server based on tcp connections

# Features 

1. adding mux handlers (same as go ServeMux but simpler)

2. being able to parse url params and use them in handlers 

3. multithreaded 

4. all the functionality you expect from a super simple HTTP server  

# TODO 

- add method check (users be able to add method before the path => "POST /test/hello" )