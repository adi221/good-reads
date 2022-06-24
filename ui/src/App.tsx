import GlobalStyle, { AppContainer } from './App.styles';
import AlertCenter from './components/AlertCenter/AlertCenter';
import Routes from './routes/Routes';

const App = () => (
  <AppContainer>
    <GlobalStyle />
    <Routes />
    <AlertCenter/>
  </AppContainer>
)

export default App
