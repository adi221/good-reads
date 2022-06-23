import { FC } from 'react'
import { useMutation } from '@apollo/client'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import TextField from '@mui/material/TextField';
import { loginFormSchema } from './LoginForm.schema'
import {
  StyledForm,
  CallToAction,
  HighlightedLink,
  FormContainer,
  StyledHeader,
  StyledButton
} from './LoginForm.styles';
import { Subtitle1 } from '../../../styles/typography';
import { SignUpUser } from '../../../apollo/users/mutations';
import { RoutesDict } from '../../../@types/enums';

export const formVariant = {
  exit: { y: '5rem', opacity: 0 },
};

export const transition = {
  duration: 0.2,
};

interface LoginFormFields {
  username: string
  email: string
  password: string
}

const LoginForm: FC = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<LoginFormFields>({
    resolver: yupResolver(loginFormSchema),
  })

  const [signUpUserMutation] = useMutation(SignUpUser)

  const onSubmit = (data: LoginFormFields) => {
    signUpUserMutation({
      variables: data
    })
  }

  return (
    <FormContainer
      variants={formVariant}
      exit="exit"
      transition={transition}
    >
      <StyledHeader>Log in</StyledHeader>
      <StyledForm onSubmit={handleSubmit(onSubmit)}>
        <TextField
          label='Username'
          type='text'
          {...register('username')}
          error={!!errors?.username}
          helperText={errors?.username?.message}
        />
        <TextField
          label='Email'
          type='text'
          {...register('email')}
          error={!!errors?.email}
          helperText={errors?.email?.message}
        />
        <TextField
          label='Password'
          type='password'
          {...register('password')}
          error={!!errors?.password}
          helperText={errors?.password?.message}
        />
        <StyledButton
          variant="contained"
          type="submit"
        >
          <Subtitle1>
            Log in
          </Subtitle1>
        </StyledButton>
        <CallToAction>
          Don&apos;t have an account?{' '}
          <HighlightedLink
            to={RoutesDict.SIGN_UP}
          >
            Sign up
          </HighlightedLink>
        </CallToAction>
      </StyledForm>
    </FormContainer>
  )
}

export default LoginForm