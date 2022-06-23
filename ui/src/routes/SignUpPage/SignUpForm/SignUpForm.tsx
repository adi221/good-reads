import { FC } from 'react'
import { useMutation } from '@apollo/client'
import { useForm } from 'react-hook-form'
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
import { RoutesDict } from '../../../@types/enums';

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
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<SignUpFormFields>({
    resolver: yupResolver(signUpFormSchema),
  })

  const [signUpUserMutation] = useMutation(SignUpUser)

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