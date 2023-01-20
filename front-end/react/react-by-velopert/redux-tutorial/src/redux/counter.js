import { createAction, handleActions } from 'redux-actions';

// Define Action Type
const INCREASE = 'INCREASE';
const DECREASE = 'DECREASE';

// Define Action Creator
//export const increase = () => ({type: INCREASE});
//export const decrease = () => ({type: DECREASE});
export const increase = createAction(INCREASE);
export const decrease = createAction(DECREASE);

// Define Reducer
const initialState = {
  number: 0,
};

function counterLegacy(state = initialState, action) {
  switch (action.type) {
    case INCREASE:
      return {
        number: state.number + 1
      };
    case DECREASE:
      return {
        number: state.number - 1
      };
    default:
      return state;
  }
}

const counter = handleActions(
  {
    [INCREASE]: (state, action) => ({number: state.number + 1}),
    [DECREASE]: (state, action) => ({number: state.number - 1})
  },
  initialState
)

export default counter;