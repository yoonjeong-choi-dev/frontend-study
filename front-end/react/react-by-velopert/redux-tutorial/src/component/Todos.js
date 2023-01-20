const TodoItem = ({todo, onToggle, onRemove}) => {
  return (
    <div style={{display: 'flex', gap: '3px', marginBottom: '5px'}}>
      <input
        type="checkbox"
        onClick={() => onToggle(todo.id)}
        checked={todo.done}
        readOnly={true}
      />
      <span style={{textDecoration: todo.done ? 'line-through' : 'none'}}>
        {todo.content}
      </span>
      <button onClick={() => onRemove(todo.id)}>Delete</button>
    </div>
  );
};

const Todos = (
  {
    input,
    todos,
    onChangeInput,
    onInsert,
    onToggle,
    onRemove,
  }) => {
  const onSubmit = e => {
    e.preventDefault();
    onInsert(input);
    onChangeInput('');
  }

  const onChange = e => onChangeInput(e.target.value);

  return (
    <div>
      <form onSubmit={onSubmit} style={{display: 'flex', gap: '3px', marginBottom: '5px'}}>
        <input value={input} onChange={onChange}/>
        <button type="submit">Register</button>
      </form>
      <div>
        {todos.map(todo => (
          <TodoItem
            key={todo.id}
            todo={todo}
            onToggle={onToggle}
            onRemove={onRemove}
          />
        ))}
      </div>
    </div>
  );
}

export default Todos;