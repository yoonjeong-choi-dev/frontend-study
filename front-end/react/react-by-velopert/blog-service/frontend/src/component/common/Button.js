import styled, { css } from 'styled-components';

import palette from '../../lib/styles/palette';
import { Link } from 'react-router-dom';

const buttonStyle = css`
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: bold;
  padding: 0.25rem 1rem;
  color: white;
  outline: none;
  cursor: pointer;

  background: ${ palette.gray[8] };

  &:hover {
    background: ${ palette.gray[6] };
  }

  ${ props =>
          props.fullWidth &&
          css`
            padding: 0.75rem 0;
            width: 100%;
            font-size: 1.125rem;
          `
  }

  ${ props =>
          props.indigo &&
          css`
            background: ${ palette.indigo[5] };

            &:hover {
              background: ${ palette.indigo[4] };
            }
          `
  }
`;

const SCButton = styled.button`
  ${ buttonStyle }
`;

const SCLinkButton = styled(Link)`
  ${ buttonStyle }
`;

const Button = (props) => {
  return props.to ? (
    <SCLinkButton { ...props } indigo={ props.indigo ? 1 : 0 } />
  ) : (
    <SCButton { ...props } />
  );
};

export default Button;