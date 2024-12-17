import { NextResponse } from 'next/server';

export async function POST(request: Request) {
  try {
    const { level, message, data } = await request.json();

    switch (level) {
      case 'info':
        console.info(`[Server Log] ${message}`, data);
        break;
      case 'warn':
        console.warn(`[Server Log] ${message}`, data);
        break;
      case 'error':
        console.error(`[Server Log] ${message}`, data);
        break;
      default:
        console.log(`[Server Log] ${message}`, data);
    }

    return NextResponse.json({ status: 'Logged' }, { status: 200 });
  } catch (error) {
    console.error('[Server Log] Error processing log request', {
      error: error instanceof Error ? error.message : String(error),
    });
    return NextResponse.json(
      { status: 'Error logging message' },
      { status: 500 },
    );
  }
}
