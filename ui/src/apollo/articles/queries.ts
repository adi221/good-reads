import gql from 'graphql-tag'

export const GetArticle = gql`
  query article($id: ID!) {
    article(id: $id) {
      id
      title
    }
  }
`