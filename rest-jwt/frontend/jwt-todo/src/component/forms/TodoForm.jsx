import Input from "../common/Input";
import Button from "../common/Button";

const TodoForm = ({ handleTodoSubmit }) => {
  return (
    <form action="">
      <span>todo名 : </span>
      <Input type="text" />
      <br />
      <Button text="Todo作成" onClick={handleTodoSubmit} />
    </form>
  );
};

export default TodoForm;
