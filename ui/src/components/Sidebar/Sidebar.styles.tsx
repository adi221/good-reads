import styled from 'styled-components';
import { NavLink } from 'react-router-dom';
import { motion } from 'framer-motion';

interface Props {
  expanded: boolean | null;
}

export const ToggleExpandedButton = styled(motion.button)`
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  top: 50%;
  left: 100%;
  transform: translate(-50%, -50%);
  border-radius: ${(props) => props.theme.borderRadius.circle};
  border: none;
  color: white;
  background-color: ${(props) => props.theme.colors.primary2};
  cursor: pointer;
  transform-origin: center;
  z-index: ${(props) => props.theme.zIndex.raised};

  opacity: 0;
  transition: opacity 300ms ease-out;

  & > * {
    font-size: ${(props) => props.theme.fontSize.xlarge};
  }

  &:focus {
    outline: none;
  }
`;

export const Container = styled(motion.nav)<Props>`
  height: 100vh;
  background-color: ${(props) => props.theme.colors.primary1};
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding:  ${(props) => props.theme.spacing['32']};
  position: relative;

  &:hover ${ToggleExpandedButton} {
    opacity: 1;
  }
`;

export const TopContainer = styled.div<Props>`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;

  & > :not(:last-child) {
    margin-right: ${(props) => props.theme.spacing['24']};
  }
`;

export const Username = styled.h4`
  font-weight: 700;
  font-size: ${(props) => props.theme.spacing['24']};
  color: ${(props) => props.theme.white};
`;

export const ListItem = styled(motion.li)`
  list-style-type: none;
  margin-bottom: ${(props) => props.theme.spacing['40']};
`;

export const LinkText = styled(motion.h4)<{ selected: boolean }>`
  font-weight: 600;
`;

export const Link = styled(NavLink)<{ selected: boolean }>`
  position: relative;
  display: flex;
  align-items: center;
  text-align: left;
  text-decoration: none;
  font-size:  ${(props) => props.theme.fontSize.medium};
  color: ${(props) =>
    props.selected ? props.theme.colors.tertiary1 : props.theme.colors.purewhite};
  transition: color 200ms;
  z-index: ${(props) => props.theme.zIndex.raised};

  & > *:not(:last-child) {
    margin-right: ${(props) => props.theme.spacing['24']};
  }
`;

export const IconContainer = styled(motion.span)`
  width: ${(props) => props.theme.spacing['40']};
  height: ${(props) => props.theme.spacing['40']};
  position: relative;

  & > :first-child {
    width: 100%;
    height: 100%;
  }

  &:hover + & > * {
    opacity: 1;
  }
`;

export const AvatarImage = styled(motion.img)`
  border-radius: ${(props) => props.theme.borderRadius.circle};
  width: ${(props) => props.theme.spacing['40']};
  height: ${(props) => props.theme.spacing['40']};
`;
