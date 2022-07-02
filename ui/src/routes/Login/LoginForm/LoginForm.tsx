import { FC } from 'react'
import { ApolloError, useMutation } from '@apollo/client'
import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom';
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
import { LoginUser } from '../../../apollo/users/mutations';
import { RoutesDict } from '../../../utils/enums';
import { getGraphQLErrors } from '../../../utils/graphql';
import { UserErrCodes } from '../../../services/server/users';
import { useMessage } from '../../../contexts/MessageContext'
import { MutationLoginUserArgs, User } from '../../../generated/graphql';

export const formVariant = {
  exit: { y: '5rem', opacity: 0 },
};

export const transition = {
  duration: 0.2,
};

interface LoginFormFields {
  usernameOrEmail: string
  password: string
}

const LoginForm: FC = () => {
  const navigate = useNavigate()
  const { showMessage } = useMessage()

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<LoginFormFields>({
    resolver: yupResolver(loginFormSchema),
  })

  const onLoginCompleted = () => {
    navigate(RoutesDict.ARTICLES)
  }

  const onLoginError = (error: ApolloError) => {
    const errCodes = getGraphQLErrors(error).map(err => err.code)
    if (errCodes.includes(UserErrCodes.ERR_USER_NON_EXIST) || errCodes.includes(UserErrCodes.ERR_INCORRECT_CREDS)) {
      showMessage({ text: 'Username/email or password are incorrect' })
    } else {
      showMessage({ text: 'Failed to log in you up! Please try again'})
    }
  }

  const [loginUserMutation] = useMutation<User, MutationLoginUserArgs>(LoginUser, {
    onCompleted: onLoginCompleted,
    onError: onLoginError
  })

  const onSubmit = (data: LoginFormFields) => {
    loginUserMutation({
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
          label='Username or email'
          type='text'
          {...register('usernameOrEmail')}
          error={!!errors?.usernameOrEmail}
          helperText={errors?.usernameOrEmail?.message}
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