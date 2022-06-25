import { ReactNode } from 'react'
import { ApolloError, GraphQLErrors } from '@apollo/client/errors';

interface HttpErrorModel {
  message: String
  code: string
  statusCode: number
}

export const getGraphQLErrors = (err: ApolloError): HttpErrorModel[] => {
  // @ts-ignore
  return (err?.networkError?.result?.errors as GraphQLErrors).map((e: any) => e?.extensions)
}

export interface GQLResponsePattern<T> {
  Loading: () => ReactNode
  Error: (err: ApolloError | Error) => ReactNode
  Data: (data: T) => ReactNode
}

export function matchResponse<T>(
  p: GQLResponsePattern<T>
): (loading: boolean, data?: T, error?: ApolloError | Error) => ReactNode {
  return (loading: boolean, data?: T, error?: ApolloError | Error): ReactNode => {
    if (loading) {
      return p.Loading()
    }
    if (error !== undefined) {
      return p.Error(error)
    }
    if (data !== undefined) {
      return p.Data(data)
    }
    return null
  }
}
