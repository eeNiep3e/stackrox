import { JobContextTab, jobContextTabs } from './types';

export function ensureJobContextTab(value: unknown): JobContextTab {
    if (typeof value === 'string' && jobContextTabs.includes(value as JobContextTab)) {
        return value as JobContextTab;
    }
    return 'CONFIGURATION_DETAILS';
}