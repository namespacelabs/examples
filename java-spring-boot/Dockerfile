FROM openjdk:18.0.2.1-jdk

WORKDIR /app

COPY .mvn/ .mvn
COPY mvnw pom.xml ./
COPY src ./src

RUN ./mvnw clean package spring-boot:repackage
RUN cp ./target/*.jar ./target/app.jar
