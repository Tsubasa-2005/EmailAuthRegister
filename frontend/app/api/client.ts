import { DefaultApi, Configuration } from '@/generated';

const apiClient = new DefaultApi(
  new Configuration({
    basePath: process.env.API_BASE_URL,
  })
);

export default apiClient;
