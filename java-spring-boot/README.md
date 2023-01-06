# Java Spring Boot application

This directory demonstrates how to model a Java Spring Boot application with Namespace.
For this example, we use Postgres for the persistence layer.

We're using Dockerfile for a simple integration with Java Spring Boot.
A native Java integration (on our roadmap) will allow best-in-class build performance and minimal image size.

The application server consumes the database as a Namespace resource provided by a shared Postgres provider.
The resource produces a typed instance object which provides a password along with the endpoint.
Another advantage is that resources have their own lifetime modeling and initialization only happens once.
In the case of a database, this means that the schema will be applied as part of the lifecycle and the Java application does not need to worry about it.
