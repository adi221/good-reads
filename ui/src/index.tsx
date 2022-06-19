import ReactDOM from 'react-dom/client'
import App from './App'
import GraphQLProvider from './contexts/GraphQLContext'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(
  <GraphQLProvider>
    <App />
  </GraphQLProvider>
)
