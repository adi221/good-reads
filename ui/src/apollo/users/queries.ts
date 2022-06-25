import gql from 'graphql-tag'
import { RegularUser } from './fragmnets'

export const Me = gql`
  ${RegularUser}
  query Me {
    me {
      ...RegularUserDetails
    }
  }
`