import Input from "../common/Input";
import Button from "../common/Button";

const TodoForm = ({ handleTodoSubmit, todo, setTodo }) => {

  return (
    <form action="">
      <span>todo名 : </span>
      <Input type="text" value={todo} onChange={(e) => setTodo(e.target.value)} />
      <br />
      <Button text="Todo作成" onClick={handleTodoSubmit} />
    </form>
  );
};

export default TodoForm;
