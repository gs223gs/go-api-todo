import TodoForm from "../forms/TodoForm";
import TodoList from "../todos/TodoList";
import Button from "../common/Button";
const Todos = ({ todoList, handleTodoSubmit, handleLogout }) => {

  //TODO
  //formにstateを送る
  return (
    <div>
      <Button text="logout" onClick={handleLogout} />
      <TodoForm handleTodoSubmit={handleTodoSubmit} />
      <TodoList todoList={todoList} />
    </div>
  );
};

export default Todos;
