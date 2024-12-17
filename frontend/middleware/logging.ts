import { Middleware, RequestContext, ResponseContext } from '@/generated';

interface LogData {
  headers?: HeadersInit;
  body?: unknown;
  status?: number;
  responseHeaders?: HeadersInit;
  error?: string;
}

async function sendLogToServer(level: string, message: string, data: LogData) {
  try {
    await fetch('/api/log', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ level, message, data }),
    });
  } catch (error) {
    console.error('Failed to send log to server:', error);
  }
}

export const loggingMiddleware: Middleware = {
  pre: async (context: RequestContext) => {
    const message = `[HTTP] <-- ${context.init.method} ${context.url}`;
    let parsedBody: unknown = context.init.body;

    if (typeof parsedBody === 'string') {
      try {
        const parsedJson = JSON.parse(parsedBody);

        if (typeof parsedJson === 'object' && parsedJson !== null) {
          delete parsedJson.token;
        }
        parsedBody = parsedJson;
      } catch (error) {
        console.warn('Failed to parse body as JSON:', error);
      }
    }

    const data: LogData = {
      body: parsedBody,
    };

    await sendLogToServer('info', message, data);

    return context;
  },

  post: async (context: ResponseContext) => {
    const message = `[HTTP] --> ${context.init.method} ${context.url}`;
    const data: LogData = {
      status: context.response.status,
      body: context.response.body,
    };

    await sendLogToServer('info', message, data);

    return context.response;
  },

  onError: async (context) => {
    const error = context.error as Error;
    const message = `[HTTP] FAIL ${context.init.method} ${context.url}`;
    const data: LogData = {
      error: error.message,
      status: context.response ? context.response.status : undefined,
      responseHeaders: context.response?.headers,
    };

    await sendLogToServer('error', message, data);

    return context.response;
  },
};
