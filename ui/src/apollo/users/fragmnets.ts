import gql from 'graphql-tag'

export const RegularUser = gql`
  fragment RegularUserDetails on User {
    id
    username
    email
  }
`