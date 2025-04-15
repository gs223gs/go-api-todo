const Login = ({ username, password, setUsername, setPassword, handleLoginSubmit }) => {
  return  (
    <div>
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
  </div>
  );
};

export default Login;
