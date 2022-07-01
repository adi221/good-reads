import { useLocation } from 'react-router-dom'
import GlobalStyle, { AppContainer } from './App.styles';
import AlertCenter from './components/AlertCenter/AlertCenter';
import Sidebar from './components/Sidebar/Sidebar';
import Routes from './routes/Routes';
import { RoutesDict } from './utils/enums';

const routesWithoutSidebar = new Set<string>([RoutesDict.LOGIN, RoutesDict.SIGN_UP])

const App = () => {
  const { pathname } = useLocation()
  const shouldShowSidebar = !routesWithoutSidebar.has(pathname)

  return (
    <AppContainer>
      <GlobalStyle />
      {shouldShowSidebar && <Sidebar />}
      <Routes />
      <AlertCenter/>
    </AppContainer>
  )
}


export default App
