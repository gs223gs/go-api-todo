const Input = ({ value, onChange, type="text" }) => {
  return (
    <input
      type={type}
      value={value}
      onChange={onChange}
    />
  );
};

export default Input;

