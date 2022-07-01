import { FC } from 'react'
import {  Routes as Switch, Route } from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute/ProtectedRoute';
import GraphiQLPage from '../graphiql/GraphiQLPage';
import { RoutesDict } from '../utils/enums';
import Login from './Login/Login';
import SignUp from './SignUp/SignUp';
import Settings from './Settings/Settings';
import Articles from './Articles/Articles';
import Categories from './Categories/Categories';

const Routes: FC = () => (
  <Switch>
    <Route path={RoutesDict.ARTICLES} element={
      <ProtectedRoute>
        <Articles/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.CATEGORIES} element={
      <ProtectedRoute>
        <Categories/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.SETTINGS} element={
      <ProtectedRoute>
        <Settings/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.GRAPHIQL} element={
      <ProtectedRoute>
        <GraphiQLPage/>
      </ProtectedRoute>
    }/>
    <Route path={RoutesDict.LOGIN} element={<Login/>}/>
    <Route path={RoutesDict.SIGN_UP} element={<SignUp/>}/>
  </Switch>
)

export default Routes