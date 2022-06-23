import { FC } from 'react';
import {
  PageContainer,
  Logo,
  ShowcaseContainer,
  FormContainer,
  ShowcaseImg,
  Showcase,
  ShowcaseTitle,
} from './LoginPage.styles';
import mockup from '../../assets/img/mockup.png';
import { transition } from '../../utils/animations';
import LoginForm from './LoginForm/LoginForm';

const formVariant = {
  initial: { x: 0, y: '5rem', opacity: 0 },
  animate: { x: 0, y: 0, opacity: 1 },
  exit: { x: '-100%' },
};
const showcaseVariant = {
  ...formVariant,
  initial: { x: 0, y: 0, opacity: 1 },
  exit: { x: '100%' },
};

const LoginPage: FC = () => {
  return (
    <PageContainer key="loginPage">
      <ShowcaseContainer
        variants={showcaseVariant}
        initial="initial"
        animate="animate"
        exit="exit"
        transition={transition}
        key="loginShowcaseContainer"
      >
        <Logo>Good Reads</Logo>
        <Showcase>
          <ShowcaseImg src={mockup} />
          <ShowcaseTitle>The future of reading</ShowcaseTitle>
        </Showcase>
      </ShowcaseContainer>
      <FormContainer
        variants={formVariant}
        initial="initial"
        animate="animate"
        exit="exit"
        transition={transition}
        key="loginFormContainer"
      >
        <LoginForm />
      </FormContainer>
    </PageContainer>
  );
};

export default LoginPage;
