import { createStyles, makeStyles } from '@material-ui/styles';
import { theme } from '@weaveworks/weave-gitops';
import styled from 'styled-components';

export const useCanaryStyle = makeStyles(() =>
  createStyles({
    rowHeaderWrapper: {
      margin: `${theme.spacing.small} 0`,
      display: 'flex',
      justifyContent: 'start',
      alignItems: 'center',
    },
    cardTitle: {
      fontWeight: 600,
      fontSize: theme.fontSizes.normal,
      color: theme.colors.neutral30,
    },
    body1: {
      fontWeight: 400,
      fontSize: theme.fontSizes.normal,
      color: theme.colors.black,
      marginLeft: theme.spacing.xs,
      display: 'flex',
      justifyContent: 'start',
      alignItems: 'center',
    },
    colorGreen: {
      color: theme.colors.success,
    },
    statusWrapper: {
      display: 'flex',
      gap: theme.spacing.xs,
      justifyContent: 'start',
      alignItems: 'center',
    },
    statusMessage: {
      color: '#9E9E9E', // add natural25 to core
    },
    statusReady: {
      color: theme.colors.success,
    },
    statusWaiting: {
      color: '#F2994A',
    },
    statusFailed: {
      color: theme.colors.alert,
    },
    sectionHeaderWrapper: {
      background: theme.colors.neutral10,
      padding: `${theme.spacing.base} ${theme.spacing.xs}`,
      margin: `${theme.spacing.base} 0`,
    },
    straegyIcon: {
      marginLeft: theme.spacing.small,
    },
    barroot: {
      backgroundColor: theme.colors.success,
    },
    root: {
      backgroundColor: theme.colors.neutral20,
      width: '100%',
      height: 8,
      marginRight: '4px',
      borderRadius: 5,
    },
    code: {
      wordBreak: 'break-word',
    },
  }),
);

export const TableWrapper = styled.div`
  margin-top: ${theme.spacing.medium};
  div[class*='FilterDialog__SlideContainer'],
  div[class*='SearchField'] {
    overflow: hidden;
  }
  div[class*='FilterDialog'] {
    .Mui-checked {
      color: ${theme.colors.primary};
    }
  }
  tr {
    vertical-align:'center';
  }
  max-width: calc(100vw - 220px);
`;

export const OnBoardingMessageWrapper = styled.div`
  background: rgba(255, 255, 255, 0.85);
  box-shadow: 5px 10px 50px 3px rgb(0 0 0 / 10%);
  border-radius: 10px;
  padding: ${theme.spacing.large} ${theme.spacing.xxl};
  max-width: 560px;
  margin: auto;
`;

export const Header4 = styled.div`
  font-size: ${theme.fontSizes.large};
  font-weight: 600;
  color: ${theme.colors.neutral30};
  margin-bottom: ${theme.spacing.small};
`;

export const TextWrapper = styled.p`
  font-size: ${theme.fontSizes.normal};
  color: ${theme.colors.neutral30};
  font-weight: 400;
`;

export const FlexCenter = styled.div`
  display: flex;
  lign-items: center;
  justify-content: center;
`;

export const LinkTag = styled.a`
  color: ${theme.colors.primary};
`;