import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { motion } from 'framer-motion';
import Button from '@mui/material/Button';
import { H4 } from '../../../styles/typography';

export const FormContainer = styled(motion.div)`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
`;

export const StyledHeader = styled(H4)`
  margin-bottom: ${(props) => props.theme.spacing['64']};
`

export const StyledButton = styled(Button)`
  width: 100%;
  height: ${(props) => props.theme.spacing['48']};
  text-align: center;
  border-radius: 100px;
  background-color: ${(props) => props.theme.colors.tertiary1};
  opacity: ${(props) => (props.disabled ? props.theme.opacity.medium : props.theme.opacity.high)};
  color: ${(props) => props.theme.colors.purewhite};
  font-family: inherit;
  font-weight: 400;
  border: none;
  cursor: ${(props) => (props.disabled ? 'not-allowed' : 'pointer')};
  border: 2px solid transparent;
  transition: border-color 200ms ease-out;
  position: relative;

  &:active,
  &:focus {
    outline: none;
    border-color: ${(props) => props.theme.colors.white};
  }`

export const StyledForm = styled.form`
  width: 50%;
  display: flex;
  flex-direction: column;
  gap: ${(props) => props.theme.spacing['16']};
`;

export const CallToAction = styled.h3`
  color: ${(props) => props.theme.colors.purewhite};
  font-weight: 400;
  text-align: center;
`;

export const HighlightedLink = styled(Link)`
  text-decoration: none;
  color: ${(props) => props.theme.colors.alert};
`;

export const ForgotPassword = styled(Link)`
  text-decoration: none;
  color: ${(props) => props.theme.colors.purewhite};
  font-weight: 400;
`;
