import { FC} from 'react';
import { PageContainer, Logo } from './LoadingPage.styles';

const LoadingPage: FC = () => (
  <PageContainer initial={{ opacity: 1 }} exit={{ opacity: 0 }}>
    <Logo>Good Reads</Logo>
  </PageContainer>
);

export default LoadingPage;
