import React from 'react';
import {
    Alert,
    Bullseye,
    Card,
    CardBody,
    DescriptionListDescription,
    DescriptionListGroup,
    DescriptionListTerm,
    Flex,
    Spinner,
} from '@patternfly/react-core';

import useFeatureFlags from 'hooks/useFeatureFlags';
import { getAxiosErrorMessage } from 'utils/responseErrorUtils';
import NotifierConfigurationView from 'Components/NotifierConfiguration/NotifierConfigurationView';
import {
    getBodyDefault,
    getSubjectDefault,
    getTimeWithHourMinuteFromISO8601,
} from '../compliance.scanConfigs.utils';
import ScanConfigParametersView from './ScanConfigParametersView';
import ScanConfigClustersTable from './ScanConfigClustersTable';
import ScanConfigProfilesView from './ScanConfigProfilesView';

const headingLevel = 'h2';

function ConfigDetails({ isLoading, error, scanConfig }) {
    const { isFeatureFlagEnabled } = useFeatureFlags();
    const isComplianceReportingEnabled = isFeatureFlagEnabled('ROX_COMPLIANCE_REPORTING');

    if (isLoading) {
        return (
            <Bullseye>
                <Spinner />
            </Bullseye>
        );
    }

    if (error) {
        return (
            <Alert variant="warning" title="Unable to fetch scan schedule" component="p" isInline>
                {getAxiosErrorMessage(error)}
            </Alert>
        );
    }

    if (scanConfig) {
        return (
            <Card>
                <CardBody>
                    <Flex
                        direction={{ default: 'column' }}
                        spaceItems={{ default: 'spaceItemsLg' }}
                    >
                        <ScanConfigParametersView
                            headingLevel={headingLevel}
                            scanName={scanConfig.scanName}
                            description={scanConfig.scanConfig.description}
                            scanSchedule={scanConfig.scanConfig.scanSchedule}
                        >
                            <DescriptionListGroup>
                                <DescriptionListTerm>Last run</DescriptionListTerm>
                                <DescriptionListDescription>
                                    {scanConfig.lastExecutedTime
                                        ? getTimeWithHourMinuteFromISO8601(
                                              scanConfig.lastExecutedTime
                                          )
                                        : 'Scan is in progress'}
                                </DescriptionListDescription>
                            </DescriptionListGroup>
                            <DescriptionListGroup>
                                <DescriptionListTerm>Last updated</DescriptionListTerm>
                                <DescriptionListDescription>
                                    {getTimeWithHourMinuteFromISO8601(scanConfig.lastUpdatedTime)}
                                </DescriptionListDescription>
                            </DescriptionListGroup>
                        </ScanConfigParametersView>
                        <ScanConfigClustersTable
                            headingLevel={headingLevel}
                            clusterScanStatuses={scanConfig.clusterStatus}
                        />
                        <ScanConfigProfilesView
                            headingLevel={headingLevel}
                            profiles={scanConfig.scanConfig.profiles}
                        />
                        {isComplianceReportingEnabled && (
                            <NotifierConfigurationView
                                customBodyDefault={getBodyDefault(scanConfig.scanConfig.profiles)}
                                customSubjectDefault={getSubjectDefault(
                                    scanConfig.scanName,
                                    scanConfig.scanConfig.profiles
                                )}
                                notifierConfigurations={scanConfig.scanConfig.notifiers}
                            />
                        )}
                    </Flex>
                </CardBody>
            </Card>
        );
    }
}

export default ConfigDetails;