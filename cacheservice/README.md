# Key Value Project

Before starting on this training document, make sure you have completed:

The reading and exercises outlined in the go training.
The basic setup of your Linux Environment.
The purpose of this task it to show how to use standard programming concepts and how to apply them to go.



What we're looking for.

SOLID principles implemented in go.

Professional code - treat this as you would a customer facing product. This should include good naming conventions/error handling/logging.



You will be creating an interactive cache service that allows users to update values in a cache from an http interface and impact a running project.



The project will consist of the following parts.

- An initial setup and logging.
- An http handler layer that receives http requests, performs necessary validation, and applies error handling back to the user.
A data access layer, in the form an in memory cache. This layer needs to be able to handle multiple threads
In order to complete this project you need to be able to perform CRUD operations on the in memory cache via the http interfaces.

When storing the values in the cache, don't just store them as a string, but pass in the type as part of the http request.
If a real service was using the cache it doesn't want to be performing data conversions every time it requests data from the cache to the type it needs, as that would slow down the operation.

## 1.1       Git Setup
***
Create a private git project under your name

Initialize it with a readme.md file.

Clone it to your machine, and create a new branch.

Create a merge request for your new branch.

This way all comments can be seen and tracked properly as per any QA.

Make sure to share this repo with your mentor and get them to review the initial setup.



## 1.2       Go Project  Setup
***
Create the go.mod file and main.go files and make sure you have a buildable hello world script.

Create a very basic logger function.

Extend the logger functionality so you can have log levels of debug, info, error

Move the logger functionality into its own file, confirm it still works.

Move the logger file into its own directory (which then becomes it's own package if you rename the package at the top of the file to match the new directory name)

Import the logger package into main.go and use that for your hello world script.

Submit this to your mentor for a quick review.



## 1.3       Serve from HTTP
***
Host an HTTP server in your code (do not use third party systems e.g. GorillaMux to do this, only use native go code)

Implement a simple get request, add a logger and confirm it's working.

Implement your CRUD handlers (don't implement the caching layer yet)

Think about what you want the incoming requests to look like e.g. post vs get and json vs query string

Add logging, validation, and http response codes

Move the handler to a new package, in your main.go make it so you can call a simple setup() and start() function.

Submit this to your mentor for a quick review.



## 1.4       Implement the cache
***
Thinking back to your go tour, what is the best structure for storing data in an easy to lookup by name way.

How are you going to protect the cache from asynch read/writes

Implement CRUD functions on your cache

Make sure you have validation and useful error handling in each layer



## 1.5       Integration
***
Instantiate your cache into the main.go and inject it into the http setup.

Use an interface. The http service should't care about implementation, just about the footprint of the methods.

Wire it all up.



## 1.6       Testing
***
Implement unit testing at the cache layer.

Make sure if you insert a type of int64 and request the object back, that the type matches.

