
const TodoItem = ({ todo }) => {
  return (
    <div>
      <p>{todo.todo}</p>
      <p>{todo.isdone ? "完了" : "未完了"}</p>
    </div>
  );
};

export default TodoItem;
