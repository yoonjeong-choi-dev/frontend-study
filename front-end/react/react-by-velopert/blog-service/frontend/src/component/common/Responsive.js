import styled from 'styled-components';

const SCWrapper = styled.div`
  padding: 0 1rem;
  width: 1024px;
  margin: 0 auto;

  @media (max-width: 1024px) {
    width: 768px;
  }
  @media (max-width: 768px) {
    width: 100%;
  }
`;

const Responsive = ({ children, ...rest }) => {
  return <SCWrapper { ...rest }>{ children }</SCWrapper>;
};

export default Responsive;