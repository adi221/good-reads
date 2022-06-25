import React, { FC } from 'react'
import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache, from, ApolloLink, ServerError } from '@apollo/client'
import { onError } from '@apollo/client/link/error'
import { setContext } from '@apollo/client/link/context';
import { API_BASE_URL } from '../constants'
import { getAccessToken, removeAccessToken, setAccessToken } from '../utils/localStorage';

const httpLink = new HttpLink({
  uri: API_BASE_URL + '/graphql',
})

const afterwareLink = new ApolloLink((operation, forward) => {
  return forward(operation).map((response) => {
    const context = operation.getContext()
    const { response: { headers } } = context;
    const token = headers.get('authorization')
    if (token) {
      setAccessToken(token)
    }
    return response
  })
})

const authLink = setContext((_, { headers }) => {
  const token = getAccessToken()
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
})

const cache = new InMemoryCache()

const errorLink = onError((err) => {
   console.error(err)
   if (err.networkError) {
    if ((err.networkError as ServerError).statusCode === 401) {
      removeAccessToken()
      // TODO: Redirect to login page..
    }
  }
})


const client = new ApolloClient({
  link: from([errorLink, authLink, afterwareLink, httpLink]),
  cache
})


interface Props {
  children: React.ReactNode
}

const GraphQLProvider: FC<Props> = ({ children }) => <ApolloProvider client={client}>{children}</ApolloProvider>
export default GraphQLProvider