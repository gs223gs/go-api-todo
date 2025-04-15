
const Register = ({ username, password, setUsername, setPassword, handleRegisterSubmit }) => {
  return (
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
  );
};
export default Register;
