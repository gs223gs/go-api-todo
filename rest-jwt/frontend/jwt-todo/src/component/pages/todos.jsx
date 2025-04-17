import TodoForm from "../forms/TodoForm";
import TodoList from "../todos/TodoList";
import Button from "../common/Button";
import { useState } from "react";
const Todos = ({ handleLogout }) => {
  const handleTodoSubmit = (e) => {
    e.preventDefault();
    //APIを叩く
    //responseが200ならば成功とポップアップ
    //todoListを更新

    setTodo("");
  };

  const [todo, setTodo] = useState("");
  const todolist = [
    { id: 1, todo: "todo", isdone: false },
    { id: 2, todo: "todo2", isdone: true },
    { id: 3, todo: "todo3", isdone: false },
  ];

  const [todoList, setTodoList] = useState(todolist);

  //TODO
  //formにstateを送る
  return (
    <div>
      <Button text="logout" onClick={handleLogout} />
      <TodoForm handleTodoSubmit={handleTodoSubmit} todo={todo} setTodo={setTodo} />
      <TodoList todoList={todoList} />
    </div>
  );
};

export default Todos;
