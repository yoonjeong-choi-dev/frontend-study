import styled from 'styled-components';

import palette from '../../lib/styles/palette';

const SCButton = styled.button`
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
`;

const Button = (props) => <SCButton { ...props } />;

export default Button;