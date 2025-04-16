import LoginForm from "../forms/LoginForm";

const Login = ({ username, password, setUsername, setPassword, handleLoginSubmit }) => {
  return  (
    <div>
    <LoginForm
      username={username}
      password={password}
      setUsername={setUsername}
      setPassword={setPassword}
      handleLoginSubmit={handleLoginSubmit}
    />
  </div>
  );
};

export default Login;
