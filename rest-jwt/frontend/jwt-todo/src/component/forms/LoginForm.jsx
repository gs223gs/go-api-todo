import Input from "../common/Input";
import Button from "../common/Button";

const LoginForm = ({
  username,
  password,
  setUsername,
  setPassword,
  handleLoginSubmit,
}) => {
  return (
    <form action="">
      <p>login</p>
      <span>username : </span>
      <Input
        type="text"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <br />
      <span>password : </span>
      <Input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <br />
      <Button text="Create" onClick={handleLoginSubmit} />
    </form>
  );
};

export default LoginForm;
