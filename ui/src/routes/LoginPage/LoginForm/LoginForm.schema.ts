import * as yup from 'yup'

export const loginFormSchema = yup.object().shape({
  usernameOrEmail: yup.string()
    .min(3, 'Username or email must be at least 3 characters.')
    .required('Username or email is a required field.'),
  password: yup.string()
    .min(6, 'Your password must be at least 6 characters.')
    .max(20, 'Your password must be at max 20 characters.')
    .required('Password is a required field.'),
})