import styled from 'styled-components';
import { motion } from 'framer-motion';

export const PageContainer = styled.main`
  background-color: ${(props) => props.theme.primary2};
  min-height: 100vh;
  display: flex;
  overflow: hidden;
`;

export const Logo = styled.h1`
  font-size: ${(props) => props.theme.fontSize['2xlarge']};
  color: ${(props) => props.theme.primary2};
  margin-top: ${(props) => props.theme.spacing['64']};
`;

export const ShowcaseContainer = styled(motion.section)`
  background-color: ${(props) => props.theme.purewhite};
  flex-basis: 50%;
  padding: 0 ${(props) => props.theme.spacing['48']};
  display: flex;
  flex-direction: column;
  z-index: 2;
`;

export const ShowcaseImg = styled.img`
  height: auto;
  width: 85%;
  margin-bottom: ${(props) => props.theme.fontSize['xlarge']};
`;

export const Showcase = styled.article`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  flex-grow: 1;
`;

export const ShowcaseTitle = styled.h1`
  color: ${(props) => props.theme.colors.primary2};
  font-size: ${(props) => props.theme.fontSize['xlarge']};
`;

export const FormContainer = styled(motion.section)`
  flex-basis: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: relative;
  z-index: 1;
  background-color: ${(props) => props.theme.colors.primary2};
`;
