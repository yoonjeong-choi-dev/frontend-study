import { combineReducers } from 'redux';
import loading from './loading';
import auth from './auth';
import user from './user';

const rootReducer = combineReducers({
  loading,
  auth,
  user,
});

export default rootReducer;