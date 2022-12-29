package dev.namespacelabs.examples.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import dev.namespacelabs.examples.model.Todo;

@Repository
public interface TodoRepository extends CrudRepository<Todo, String> {

}