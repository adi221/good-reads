import { FC } from 'react'
import {  Routes as Switch, Route } from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute/ProtectedRoute';
import GraphiQLPage from '../graphiql/GraphiQLPage';
import { RoutesDict } from '../utils/enums';
import HomePage from './HomePage/HomePage';
import LoginPage from './LoginPage/LoginPage';
import SignUpPage from './SignUpPage/SignUpPage';

const Routes: FC = () => (
  <Switch>
    <Route path={RoutesDict.HOME} element={
      <ProtectedRoute>
        <HomePage/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.GRAPHIQL} element={
      <ProtectedRoute>
        <GraphiQLPage/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.LOGIN} element={<LoginPage/>}/>
    <Route path={RoutesDict.SIGN_UP} element={<SignUpPage/>}/>
  </Switch>
)

export default Routes