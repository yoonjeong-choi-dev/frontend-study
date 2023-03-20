import styled, { css } from 'styled-components';
import palette from '../../lib/styles/palette';
import { Link } from 'react-router-dom';

const SCWrapper = styled.div`
  ${ props =>
          props.hasMarginTop &&
          css`
            margin-top: 1rem;
          `
  }
  color: ${ palette.gray[6] };

  span + span:before {
    color: ${ palette.gray[5] };
    padding: 0 0.25rem;
    content: '\\B7';
  }
`;

const SubInfo = ({ name, createdAt, hasMarginTop }) => {
  return (
    <SCWrapper hasMarginTop={ hasMarginTop }>
       <span>
         <b>
           <Link to={ `/${ name }` }>{ name }</Link>
         </b>
       </span>
      <span>{ new Date(createdAt).toLocaleDateString() }</span>
    </SCWrapper>
  );
};

export default SubInfo;