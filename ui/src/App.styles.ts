import styled, { createGlobalStyle } from 'styled-components';

export const AppContainer = styled.div`
  min-height: 100vh;
`

const GlobalStyle = createGlobalStyle`
  *,
  *::before,
  *::after {
    padding: 0;
    margin: 0;
    box-sizing: inherit;
  }

  html {
    font-family: 'Open Sans', sans-serif;
    font-weight: 600;
  }

  body {
    box-sizing: border-box;
  }
`;

export default GlobalStyle;
