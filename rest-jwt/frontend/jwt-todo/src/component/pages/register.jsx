import RegisterForm from "../forms/RegisterForm";
const Register = ({ username, password, setUsername, setPassword, handleRegisterSubmit }) => {
  return (
    <div>
    <p>register</p>
    <RegisterForm
      username={username}
      password={password}
      setUsername={setUsername}
      setPassword={setPassword}
      handleRegisterSubmit={handleRegisterSubmit}
    />
  </div>
  );
};
export default Register;
