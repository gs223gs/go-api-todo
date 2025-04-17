import Input from "../common/Input";
import Button from "../common/Button";

const LoginForm = ({
  username,
  password,
  setUsername,
  setPassword,
  setIsLogin,
}) => {

  const handleLoginSubmit = (e) => {
    e.preventDefault();
    console.log(username, password);
    setUsername("");
    setPassword("");
    //APIを叩く
    //responseが200ならば成功とポップアップ
    //JWTを保存
    //GET todos/id にリクエスト JWTを使用
    //responseが200ならば成功とポップアップ
    //todoListを更新

    //成功ならIsLoginをfalseにする
    setIsLogin(false);
  };
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
      <Button text="login" onClick={handleLoginSubmit} />
    </form>
  );
};

export default LoginForm;
