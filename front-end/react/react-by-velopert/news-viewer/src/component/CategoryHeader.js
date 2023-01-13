import React from 'react';
//import styled, { css } from 'styled-components';
import styled from 'styled-components';
import { NavLink } from 'react-router-dom';

const categories = [
  {
    name: 'all',
    text: '전체보기',
  },
  {
    name: 'business',
    text: '비즈니스',
  },
  {
    name: 'entertainment',
    text: '엔터테인먼트',
  },
  {
    name: 'health',
    text: '건강',
  },
  {
    name: 'science',
    text: '과학',
  },
  {
    name: 'sports',
    text: '스포츠',
  },
  {
    name: 'technology',
    text: '기술',
  },
];

const SCCategoryContainer = styled.div`
  display: flex;
  width: 768px;
  margin: 0 auto;
  padding: 1rem;

  @media screen and (max-width: 768px) {
    width: 100%;
    overflow-x: auto;
  }
`;

// const SCCategory = styled.div`
//   font-size: 1.125rem;
//   white-space: pre;
//   text-decoration: none;
//   color: inherit;
//   padding-bottom: 0.25rem;
//
//   cursor: pointer;
//
//   :hover {
//     color: #495057;
//   }
//
//   ${props => props.active && css`
//     font-weight: 600;
//     border-bottom: 2px solid #22b8cf;
//     color: #22b8cf;
//
//     &:hover {
//       color: #3bc9db;
//     }
//   `}
//   & + & {
//     margin-left: 1rem;
//   }
// `;

const SCCategory = styled(NavLink)`
  font-size: 1.125rem;
  white-space: pre;
  text-decoration: none;
  color: inherit;
  padding-bottom: 0.25rem;

  cursor: pointer;

  :hover {
    color: #495057;
  }

  &.active {
    font-weight: 600;
    border-bottom: 2px solid #22b8cf;
    color: #22b8cf;

    &:hover {
      color: #3bc9db;
    }
  }
  
  & + & {
    margin-left: 1rem;
  }
`;

// const CategoryHeader = ({category: current, onSelect}) => {
//   return (
//     <SCCategoryContainer>
//       {categories && categories.map(category => (
//         <SCCategory
//           key={category.name}
//           active={current === category.name}
//           onClick={() => onSelect(category.name)}
//         >
//           {category.text}
//         </SCCategory>
//       ))}
//     </SCCategoryContainer>
//   );
// };

const CategoryHeader = () => {
  return (
    <SCCategoryContainer>
      {categories && categories.map(category => (
        <SCCategory
          key={category.name}
          className={({isActive}) => (isActive ? 'active' : undefined)}
          to={category.name === 'all' ? '/' : `/${category.name}`}
        >
          {category.text}
        </SCCategory>
      ))}
    </SCCategoryContainer>
  );
};

export default CategoryHeader;