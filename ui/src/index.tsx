import ReactDOM from 'react-dom/client'
import { ThemeProvider } from 'styled-components';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App'
import GraphQLProvider from './contexts/GraphQLContext'
import theme from './styles/theme';
import MessageProvider from './contexts/MessageContext';
import AuthProvider from './contexts/AuthContext';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(
  <GraphQLProvider>
    <ThemeProvider theme={theme}>
      <AuthProvider>
        <MessageProvider>
          <Router>
            <App />
          </Router>
        </MessageProvider>
      </AuthProvider>
    </ThemeProvider>
  </GraphQLProvider>
)
