import { useQuery } from '@apollo/client'
import { GetArticle } from './apollo/articles/queries'
import styles from './App.module.scss'

const App = () => {
  const { data, error, loading } = useQuery(GetArticle, {
    variables: { id: '00000000-0200-4c1b-4e12-1ba74bff4a4b' }
  })
  console.log('Data ', data, error, loading)

  return <div className={styles.app}>{JSON.stringify(data, null, 2)}</div>
}

export default App
