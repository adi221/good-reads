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