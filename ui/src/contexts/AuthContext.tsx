import { createContext, FC, useContext } from 'react'
import { useQuery } from '@apollo/client'
import { User } from '../@types/users'
import { Me } from '../apollo/users/queries'
import { matchResponse } from '../utils/graphql'
import LoadingPage from '../components/LoadingPage/LoadingPage'

interface AuthContextType {
  user: User | null
}

const AuthContext = createContext<AuthContextType>({
  user: null
})

interface Props {
  children: React.ReactNode
}

interface MeResponse {
  me: User
}

// @ts-ignore
const AuthProvider: FC<Props> = ({ children }) => {
  const { loading, data, error } = useQuery(Me)

  const render = matchResponse<MeResponse>({
    Loading: () => <LoadingPage/>,
    Error: (err) => <div>{JSON.stringify(err)}</div>,
    Data: ({ me: user }) => <AuthContext.Provider value={{ user }}>{children}</AuthContext.Provider>,
  })

  return render(loading, data, error)
}

export default AuthProvider

export const useAuth = () => useContext(AuthContext)

