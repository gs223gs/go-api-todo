
const Input = ({ value, onChange, type="text" }) => {
  return (
    <input
      type={type}
      value={value}
      onChange={(e) => onChange(e.target.value)}
    />
  );
};

export default Input;

