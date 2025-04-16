import Input from "../common/Input";
import Button from "../common/Button";
const RegisterForm = ({
  username,
  password,
  setUsername,
  setPassword,
  handleRegisterSubmit,
}) => {
  return (
    <form action="">
      <p>register</p>
      <span>username : </span>
      <Input
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <br />
      <span>password : </span>
      <Input
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <br />
      <Button text="Create" onClick={handleRegisterSubmit} />
    </form>
  );
};

export default RegisterForm;
