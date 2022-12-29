package dev.namespacelabs.examples;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

import dev.namespacelabs.examples.model.Todo;
import dev.namespacelabs.examples.repository.TodoRepository;
import jakarta.validation.Valid;

@Controller
public class TodoController {

    @Autowired
    private TodoRepository todoRepository;

    @GetMapping(value = { "/", "/todos" })
    public String listTodos(Todo todo, Model model) {

        var todos = (List<Todo>) todoRepository.findAll();

        model.addAttribute("todos", todos);

        return "todos";
    }

    @PostMapping("/todos")
    public String createTodo(@Valid Todo todo, BindingResult result, Model model) {
        if (result.hasErrors()) {
            return "todos";
        }

        todoRepository.save(todo);

        return "redirect:/todos";
    }
}