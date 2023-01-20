import { useDispatch } from 'react-redux';
import { useMemo } from 'react';
import { bindActionCreators } from 'redux';

export default function useActions(actionCreators, deps) {
  const dispatch = useDispatch();
  return useMemo(() => {
      if (Array.isArray(actionCreators)) {
        return actionCreators.map(action => bindActionCreators(action, dispatch));
      }
      return bindActionCreators(actionCreators, dispatch);
    },
    deps ? [dispatch, ...deps] : deps
  );
}