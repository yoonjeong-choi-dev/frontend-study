import { combineReducers } from 'redux';
import loading from './loading';
import auth from './auth/auth';
import user from './auth/user';
import write from './post/write';

const rootReducer = combineReducers({
  loading,
  auth,
  user,
  write,
});

export default rootReducer;