import "./App.css";
import { useState } from "react";
function App() {
  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(username, password);
    setUsername("");
    setPassword("");
    setIsLogin(true);
  };
  const handleTodoSubmit = (e) => {
    e.preventDefault();
  };
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLogin, setIsLogin] = useState(false);

  return (
    <>
      {isLogin ? (
        <div>
          <button onClick={() => setIsLogin(false)}>logout a</button>
          <form action="">
            <span>todo名 : </span>
            <input type="text" />
            <br />
            <button type="submit" onClick={handleTodoSubmit}>
              Todo作成
            </button>
            <br />
          </form>
        </div>
      ) : (
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
          <button type="submit" onClick={handleSubmit}>
            Create
          </button>
        </form>
      )}
    </>
  );
}

export default App;

/*
コンポーネント化するもの
フォーム
  input
  button





*/
