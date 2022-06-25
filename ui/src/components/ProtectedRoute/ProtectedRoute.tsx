import { FC } from 'react'
import { Navigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { RoutesDict } from '../../utils/enums';

interface Props {
  children: JSX.Element
}

const ProtectedRoute: FC<Props> = ({ children }) => {
  const { user } = useAuth()
  if (!user) {
    return <Navigate to={RoutesDict.LOGIN} />
  }
  return children
}

export default ProtectedRoute