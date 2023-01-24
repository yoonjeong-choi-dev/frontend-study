import { combineReducers } from 'redux';
import { all } from 'redux-saga/effects';

import counter, { counterSaga } from './counter';
import jsonHolder from './jsonHolder';
import jsonHolderForSaga, {jsonHolderSaga} from './jsonHolderForSaga';
import loading from './loading';

const rootReducer = combineReducers({
  counter,
  jsonHolder,
  jsonHolderForSaga,
  loading,
});

export default rootReducer;

export function* rootSaga() {
  yield all([counterSaga(), jsonHolderSaga()]);
}