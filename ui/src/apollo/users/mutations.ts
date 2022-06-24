import gql from 'graphql-tag'

export const SignUpUser = gql`
  mutation signUpUser($username: String!, $email: String!, $password: String!) {
    signUpUser(username: $username, email: $email, password: $password) {
      id
      username
      email
    }
  }
`

export const LoginUser = gql`
  mutation loginUser($usernameOrEmail: String!, $password: String!) {
    loginUser(usernameOrEmail: $usernameOrEmail, password: $password) {
      id
      username
      email
      createdAt
    }
  }
`