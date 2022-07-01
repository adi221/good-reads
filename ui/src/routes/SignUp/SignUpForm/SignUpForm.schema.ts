import * as yup from 'yup'

export const signUpFormSchema = yup.object().shape({
  username: yup.string()
    .min(3, 'Your password must be at least 3 characters.')
    .max(20, 'Your password must be at max 20 characters.')
    .required('Email is a required field.'),
  email: yup.string()
    .email('Invalid email address.')
    .required('Email is a required field.'),
  password: yup.string()
    .min(6, 'Your password must be at least 6 characters.')
    .max(20, 'Your password must be at max 20 characters.')
    .required('Password is a required field.'),
})