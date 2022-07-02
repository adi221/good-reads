import { FC } from 'react'
import { ApolloError, useMutation } from '@apollo/client'
import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom'
import { yupResolver } from '@hookform/resolvers/yup'
import TextField from '@mui/material/TextField';
import { signUpFormSchema } from './SignUpForm.schema'
import {
  StyledForm,
  CallToAction,
  HighlightedLink,
  FormContainer,
  StyledHeader,
  StyledButton
} from './SignUpForm.styles';
import { Subtitle1 } from '../../../styles/typography';
import { SignUpUser } from '../../../apollo/users/mutations';
import { RoutesDict } from '../../../utils/enums';
import { getGraphQLErrors } from '../../../utils/graphql';
import { UserErrCodes } from '../../../services/server/users'
import { useMessage } from '../../../contexts/MessageContext'
import { MutationSignUpUserArgs, User } from '../../../generated/graphql'

interface SignUpFormFields {
  username: string
  email: string
  password: string
}

export const formVariant = {
  exit: { y: '5rem', opacity: 0 },
};

export const transition = {
  duration: 0.2,
};

interface SignUpFormFields {
  username: string
  email: string
  password: string
}


const SignUpForm: FC = () => {
  const navigate = useNavigate()
  const { showMessage } = useMessage()

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<SignUpFormFields>({
    resolver: yupResolver(signUpFormSchema),
  })

  const onSignUpCompleted = () => {
    navigate(RoutesDict.ARTICLES)
  }

  const onSignUpError = (error: ApolloError) => {
    const errCodes = getGraphQLErrors(error).map(err => err.code)
    if (errCodes.includes(UserErrCodes.ERR_USERNAME_ALREADY_EXISTS)) {
      showMessage({ text: 'Please select a different username' })
    } else if (errCodes.includes(UserErrCodes.ERR_EMAIL_ALREADY_EXISTS)) {
      showMessage({ text: 'Please select a different email'})
    } else {
      showMessage({ text: 'Failed to sign you up! Please try again'})
    }
  }

  const [signUpUserMutation] = useMutation<User, MutationSignUpUserArgs>(SignUpUser, { 
    onCompleted: onSignUpCompleted,
    onError: onSignUpError,
  })

  const onSubmit = (data: SignUpFormFields) => {
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
      <StyledHeader>Sign up</StyledHeader>
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
            Sign up
          </Subtitle1>
        </StyledButton>
        <CallToAction>
          Already have an account?{' '}
          <HighlightedLink
            to={RoutesDict.LOGIN}
          >
            Log in 
          </HighlightedLink>
        </CallToAction>
      </StyledForm>
    </FormContainer>
  )
}

export default SignUpForm