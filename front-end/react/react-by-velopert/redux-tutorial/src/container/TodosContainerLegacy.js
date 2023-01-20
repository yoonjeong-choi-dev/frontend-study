import { connect } from 'react-redux';

import Todos from '../component/Todos';
import { changeInput, insert, toggle, remove } from '../redux/todos';

const TodosContainer = (
  {
    input,
    todos,
    changeInput,
    insert,
    toggle,
    remove,
  }) => {
  return (
    <Todos
      input={input}
      todos={todos}
      onChangeInput={changeInput}
      onInsert={insert}
      onToggle={toggle}
      onRemove={remove}
    />
  );
}

export default connect(
  // mapStateToProps
  ({todos}) => ({
    input: todos.input,
    todos: todos.todos,
  }),

  // mapDispatchToProps
  // : 액션 생성 함수들로 구성된 객체로 전달 시 bindActionCreator 가 내부적으로 호출되어 dispatch 와 바인딩
  {
    changeInput,
    insert,
    toggle,
    remove
  }
)(TodosContainer);