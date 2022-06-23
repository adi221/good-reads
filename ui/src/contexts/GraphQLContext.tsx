import React, { FC } from 'react'
import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache, from } from '@apollo/client'
import { onError } from '@apollo/client/link/error'
import { API_BASE_URL } from '../constants'

const httpLink = new HttpLink({
  uri: API_BASE_URL + '/graphql'
})

const cache = new InMemoryCache()

// Error interceptor
const errorLink = () => {
 return onError((err) => {
   console.error(err)
 })
}

const client = new ApolloClient({
  link: from([errorLink(), httpLink]),
  cache
})


interface Props {
  children: React.ReactNode
}

const GraphQLProvider: FC<Props> = ({ children }) => <ApolloProvider client={client}>{children}</ApolloProvider>
export default GraphQLProvider