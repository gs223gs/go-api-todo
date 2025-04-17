import LoginForm from "../forms/LoginForm";

const Login = ({ username, password, setUsername, setPassword, handleLoginSubmit, setIsLogin }) => {
  return  (
    <div>
    <LoginForm
      username={username}
      password={password}
      setUsername={setUsername}
      setPassword={setPassword}
      handleLoginSubmit={handleLoginSubmit}
      setIsLogin={setIsLogin}
    />
  </div>
  );
};

export default Login;
