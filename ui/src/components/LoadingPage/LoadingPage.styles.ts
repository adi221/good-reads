import styled from 'styled-components';
import { motion } from 'framer-motion';

export const PageContainer = styled(motion.div)`
  background-color: ${(props) => props.theme.colors.purewhite};
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: ${(props) => props.theme.zIndex.navigation};
  width: 100%;
`;

export const Logo = styled.h1`
  color: ${(props) => props.theme.colors.primary2};
  font-size: ${(props) => props.theme.fontSize['3xlarge']};
`;
