import { createAction, handleActions } from 'redux-actions';
import produce from 'immer';

// Define Action Type
const CHANGE_INPUT = 'todos/CHANGE_INPUT';
const INSERT = 'todos/INSERT';
const TOGGLE = 'todos/TOGGLE';
const REMOVE = 'todos/REMOVE';


// Define Action Creator
// export const changeInput = (input) => ({
//   type: CHANGE_INPUT,
//   input,
// });
export const changeInput = createAction(CHANGE_INPUT, input => input);


let id = 3; // initialState.todos.length + 1
// export const insert = (content) => ({
//   type: INSERT,
//   todo: {
//     id: id++,
//     content,
//     done: false,
//   }
// });
export const insert = createAction(INSERT, (content) => ({
  id: id++,
  content,
  done: false
}));

// export const toggle = (id) => ({
//   type: TOGGLE,
//   id
// });
export const toggle = createAction(TOGGLE, id => id);

// export const remove = (id) => ({
//   type: REMOVE,
//   id
// });
export const remove = createAction(REMOVE, id => id);


// Define Reducer
const initialState = {
  input: '',
  todos: [
    {
      id: 1,
      content: 'Learn React',
      done: false,
    },
    {
      id: 2,
      content: 'Learn Javascript',
      done: true,
    }
  ]
};

function todosLegacy(state = initialState, action) {
  switch (action.type) {
    case CHANGE_INPUT:
      return {
        ...state,
        input: action.input
      };
    case INSERT:
      return {
        ...state,
        todos: state.todos.concat(action.todo)
      };
    case TOGGLE:
      return {
        ...state,
        todos: state.todos.map(todo => todo.id === action.id ? {...todo, done: !todo.done} : todo)
      }
    case REMOVE:
      return {
        ...state,
        todos: state.todos.filter(todo => todo.id !== action.id)
      }
    default:
      return state;
  }
}

const todos = handleActions(
  // createAction 으로 만든 액션 생성 함수가 반환하는 액션의 데이터는 payload 로 접근해야 함
  {
    [CHANGE_INPUT]: (state, {payload: input}) =>
      produce(state, draft => {
        draft.input = input
      }),
    [INSERT]: (state, {payload: todo}) =>
      produce(state, draft => {
        draft.todos.push(todo)
      }),
    [TOGGLE]: (state, {payload: id}) =>
      produce(state, draft => {
        const toggleTodo = draft.todos.find(todo => todo.id === id);
        toggleTodo.done = !toggleTodo.done;
      }),
    [REMOVE]: (state, {payload: id}) => ({
      ...state,
      todos: state.todos.filter(todo => todo.id !== id)
    })
  },
  initialState
)

export default todos;
