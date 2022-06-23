import ReactDOM from 'react-dom/client'
import { ThemeProvider } from 'styled-components';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App'
import GraphQLProvider from './contexts/GraphQLContext'
import theme from './styles/theme';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(
  <GraphQLProvider>
    <ThemeProvider theme={theme}>
    <Router>
      <App />
    </Router>
    </ThemeProvider>
  </GraphQLProvider>
)
