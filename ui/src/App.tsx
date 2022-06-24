import GlobalStyle from './App.styles';
import AlertCenter from './components/AlertCenter/AlertCenter';
import Routes from './routes/Routes';

const App = () => (
  <div>
    <GlobalStyle />
    <Routes />
    <AlertCenter/>
  </div>
)

export default App
