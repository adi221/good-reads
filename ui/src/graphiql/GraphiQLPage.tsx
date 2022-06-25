import React, { Suspense, useCallback, FC } from 'react'
import { API_BASE_URL } from '../constants'
import 'graphiql/graphiql.min.css';
import { PageContainer } from './GraphiQLPage.styles';
import { getAccessToken } from '../utils/localStorage';

const GraphiQL = React.lazy(() => import('graphiql'))

const GraphiQLPage: FC = () => {
  const fetcher = useCallback(
    async (graphQLParams: any) => {
      try {
        const headers: HeadersInit = new Headers()
        headers.set('Content-Type', 'application/json')
        const token = getAccessToken()
        if (token) {
          headers.set('authorization', `Bearer ${token}`)
        }
        const response = await fetch(API_BASE_URL + '/graphql', {
          method: 'post',
          headers,
          credentials: 'same-origin',
          body: JSON.stringify(graphQLParams),
        })
        return await response.json()
      } catch (error) {
        console.log('Failed to fetch query/mutation', error)
      }
    },
    []
  )

  return (
    <PageContainer>
      <Suspense fallback={<div>loading...</div>}>
        <GraphiQL fetcher={fetcher} />
      </Suspense>
    </PageContainer>
  )
}

export default GraphiQLPage
