import "./App.css";
import { useState } from "react";
function App() {
  const handleLoginSubmit = (e) => {
    e.preventDefault();
    console.log(username, password);
    setUsername("");
    setPassword("");
    setIsLogin(false);
  };
  const handleRegisterSubmit = (e) => {
    e.preventDefault();
    console.log(username, password);
    setUsername("");
    setPassword("");
    setIsRegister(false);
    setIsLogin(true);
  };
  const handleTodoSubmit = (e) => {
    e.preventDefault();
  };

  const handleRegister = () => {
    setIsRegister(!isRegister);
    setIsLogin(!isLogin);
  };
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLogin, setIsLogin] = useState(false);
  const [isRegister, setIsRegister] = useState(true);
  //todolist
  //構造体 id todo isdone
  const todolist = [
    { id: 1, todo: "todo", isdone: false },
    { id: 2, todo: "todo2", isdone: true },
    { id: 3, todo: "todo3", isdone: false },
  ];
  const [todoList, setTodoList] = useState(todolist);
  return (
    <div>
      <button onClick={handleRegister}>
        {isRegister ? "login" : "register"}
      </button>
      {isRegister ? (
        <div>
          <p>register</p>
          <form action="">
            <span>username : </span>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <br />
            <span>password : </span>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <br />
            <button type="submit" onClick={handleRegisterSubmit}>
              Create
            </button>
          </form>
        </div>
      ) : isLogin ? (
        <form action="">
          <p>login</p>
          <span>username : </span>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <br />
          <span>password : </span>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <br />
          <button type="submit" onClick={handleLoginSubmit}>
            Create
          </button>
        </form>
      ) : (
        <div>
          <button onClick={() => setIsLogin(false)}>logout</button>
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
      )}
    </div>
  );
}

export default App;

/*
コンポーネント化するもの
フォーム
  input
  button





*/
