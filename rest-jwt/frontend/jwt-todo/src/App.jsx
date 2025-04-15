import "./App.css";
import { useState } from "react";
import Register from "./component/register/register";
import Login from "./component/login/login";
import Todos from "./component/todos/todos";
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

  const handleLogout = () => {
    setIsLogin(!isLogin);
    setIsRegister(!isRegister);
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
        <Register
          username={username}
          password={password}
          setUsername={setUsername}
          setPassword={setPassword}
          handleRegisterSubmit={handleRegisterSubmit}
        />
      ) : isLogin ? (
        <Login
          username={username}
          password={password}
          setUsername={setUsername}
          setPassword={setPassword}
          handleLoginSubmit={handleLoginSubmit}
        />
      ) : (
        <Todos
          todoList={todoList}
          handleTodoSubmit={handleTodoSubmit}
          handleLogout={handleLogout}
        />
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
