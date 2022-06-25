import gql from 'graphql-tag'
import { RegularUser } from './fragmnets'

export const SignUpUser = gql`
  ${RegularUser}
  mutation signUpUser($username: String!, $email: String!, $password: String!) {
    signUpUser(username: $username, email: $email, password: $password) {
      ...RegularUserDetails
    }
  }
`

export const LoginUser = gql`
  ${RegularUser}  
  mutation loginUser($usernameOrEmail: String!, $password: String!) {
    loginUser(usernameOrEmail: $usernameOrEmail, password: $password) {
      ...RegularUserDetails
    }
  }
`