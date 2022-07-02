import gql from 'graphql-tag'

export const GetCategories = gql`
  query categories($filter: FilterSchemaInput) {
    categories(filter: $filter) {
      items {
        id
        userId
        title
        createdAt
      }
      total
    }
  }
`