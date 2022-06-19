import React, { FC } from 'react'
import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache } from '@apollo/client'
import { API_BASE_URL } from '../constants'

const httpLink = new HttpLink({
  uri: API_BASE_URL + '/graphql'
})

const cache = new InMemoryCache()

const client = new ApolloClient({
  link: httpLink,
  cache
})

interface Props {
  children: React.ReactNode
}

const GraphQLProvider: FC<Props> = ({ children }) => <ApolloProvider client={client}>{children}</ApolloProvider>
export default GraphQLProvider