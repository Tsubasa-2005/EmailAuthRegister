import { DefaultApi, Configuration } from '@/generated';
import { loggingMiddleware } from '@/middleware/logging';

const getTokenAsync = async (): Promise<string> => {
  const response = await fetch('/api/session');
  const data = await response.json();

  if (response.ok) {
    return data.token;
  } else {
    return '';
  }
};

const apiClient = new DefaultApi(
  new Configuration({
    basePath: process.env.API_BASE_URL,
    accessToken: async () => await getTokenAsync(),
    middleware: [loggingMiddleware],
  })
);

export default apiClient;
