import React from 'react';
import { useSelector, useDispatch } from 'react-redux';

import Todos from '../component/Todos';
import { changeInput, insert, toggle, remove } from '../redux/todos';
import { useCallback } from 'react';
import useActions from '../hooks/useActions';

const TodosContainer = () => {
  const {input, todos} = useSelector(state => state.todos);

  // const dispatch = useDispatch();
  // const onChangeInput = useCallback(
  //   value => dispatch(changeInput(value)), [dispatch]);
  // const onInsert = useCallback(content => dispatch(insert(content)), [dispatch]);
  // const onToggle = useCallback(id => dispatch(toggle(id)), [dispatch]);
  // const onRemove = useCallback(id => dispatch(remove(id)), [dispatch]);
  const[onChangeInput, onInsert, onToggle, onRemove] = useActions(
    [changeInput, insert, toggle, remove], []);

  return (
    <Todos
      input={input}
      todos={todos}
      onChangeInput={onChangeInput}
      onInsert={onInsert}
      onToggle={onToggle}
      onRemove={onRemove}
    />
  );
}

// export default connect(
//   // mapStateToProps
//   ({todos}) => ({
//     input: todos.input,
//     todos: todos.todos,
//   }),
//
//   // mapDispatchToProps
//   // : 액션 생성 함수들로 구성된 객체로 전달 시 bindActionCreator 가 내부적으로 호출되어 dispatch 와 바인딩
//   {
//     changeInput,
//     insert,
//     toggle,
//     remove
//   }
// )(TodosContainer);

// redux hook 사용 시, connect 와 다르게 성능 최적화 필요
export default React.memo(TodosContainer);