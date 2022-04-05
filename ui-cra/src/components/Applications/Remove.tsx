import React, { FC } from 'react';
import { PageTemplate } from '../Layout/PageTemplate';
import { SectionHeader } from '../Layout/SectionHeader';
import { ContentWrapper } from '../Layout/ContentWrapper';
import { ApplicationRemove } from '@weaveworks/weave-gitops';
import { useApplicationsCount } from './utils';
import styled from 'styled-components';

const ApplicationRemoveWrapper = styled(ApplicationRemove)`
  div[role='alert'] {
    width: 100%;
  }
`;

const WGApplicationRemove: FC = () => {
  const applicationsCount = useApplicationsCount();

  const queryParams = new URLSearchParams(window.location.search);
  const name = queryParams.get('name');

  return (
    <PageTemplate documentTitle="WeGO · Application Detail">
      <SectionHeader
        path={[
          {
            label: 'Applications',
            url: '/applications',
            count: applicationsCount,
          },
          { label: `${name}` },
        ]}
      />
      <ContentWrapper type="WG">
        <ApplicationRemoveWrapper name={name || ''} />
      </ContentWrapper>
    </PageTemplate>
  );
};

export default WGApplicationRemove;