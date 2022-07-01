import React, { useState } from 'react';
import { useLocation } from 'react-router-dom';
import {
  MdArticle,
  MdCategory,
  MdSettings,
  MdChevronLeft,
} from 'react-icons/md';
import {
  Container,
  Link,
  ListItem,
  IconContainer,
  LinkText,
  AvatarImage,
  TopContainer,
  Username,
  ToggleExpandedButton,
} from './Sidebar.styles';
import defaultAvatar from '../../assets/img/defaultAvatar.jpg';
import { RoutesDict } from '../../utils/enums';
import { useAuth } from '../../contexts/AuthContext';

interface SidebarLink {
  to: string
  text: string
  icon: JSX.Element
}

const linksInTheMiddle: SidebarLink[] = [
  {
    to: RoutesDict.ARTICLES,
    text: 'Articles',
    icon: <MdArticle />,
  },
  {
    to: RoutesDict.CATEGORIES,
    text: 'Categories',
    icon: <MdCategory />,
  }
]

const linksAtTheEnd: SidebarLink[] = [
  {
    to: RoutesDict.SETTINGS,
    text: 'Settings',
    icon: <MdSettings />,
  },
]

const containerVariants = {
  expanded: {
    width: '20vw',
    transition: {
      staggerChildren: 0.1,
      staggerDirection: -1,
      when: 'beforeChildren',
      width: { type: 'spring', stiffness: 500, damping: 200 },
    },
  },
  retracted: {
    width: '140px',
    transition: {
      staggerChildren: 0.1,
      when: 'afterChildren',
      width: { type: 'spring', stiffness: 500, damping: 200 },
    },
  },
};

const toggleButtonVariants = {
  expanded: {
    transform: 'translate(-50%, -50%) rotate(0deg)',
  },
  retracted: {
    transform: 'translate(-50%, -50%) rotate(180deg)',
  },
};


const Sidebar: React.FC = () => {
  const [expanded, setExpanded] = useState(true)
  const { pathname } = useLocation();
  const { user } = useAuth()

  const renderLink = (link: SidebarLink) => (
    <ListItem key={link.text}>
      <Link
        key={link.to}
        to={link.to}
        selected={pathname === link.to}
      >
        <IconContainer layout>{link.icon}</IconContainer>
        {expanded && (
          <LinkText selected={pathname === link.to}>
            {link.text}
          </LinkText>
        )}
      </Link>
    </ListItem>
  )

  const renderList = (links: SidebarLink[]) =>(
    <ul>
        {links.map(renderLink)}
      </ul>
  )

  return (
    <Container
      expanded={expanded}
      variants={containerVariants}
      initial={expanded ? 'expanded' : 'retracted'}
      animate={expanded ? 'expanded' : 'retracted'}
    >
      <ToggleExpandedButton
        onClick={() => setExpanded(prev => !prev)}
        variants={toggleButtonVariants}
      >
        <MdChevronLeft />
      </ToggleExpandedButton>
      <TopContainer expanded={expanded}>
        <AvatarImage layout src={defaultAvatar} />
        {expanded && <Username>{user?.username}</Username>}
      </TopContainer>
      {renderList(linksInTheMiddle)}
      {renderList(linksAtTheEnd)}
    </Container>
  );
};

export default Sidebar;
