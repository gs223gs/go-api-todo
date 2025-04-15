const Todos = ({ todoList, handleTodoSubmit, handleLogout }) => {
  return (
    <div>
          <button onClick={handleLogout}>logout</button>
          <form action="">
            <span>todo名 : </span>
            <input type="text" />
            <br />
            <button type="submit" onClick={handleTodoSubmit}>
              Todo作成
            </button>
            <br />
          </form>

          <div>
            {todoList.map((todo) => (
              <div key={todo.id}>
                {todo.todo} : {todo.isdone ? "完了" : "未完了"}
              </div>
            ))}
          </div>
        </div>
  );
};

export default Todos;
