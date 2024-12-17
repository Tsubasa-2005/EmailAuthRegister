import apiClient from '@/api/client';

export async function pingAPI(): Promise<{
  success: boolean;
}> {
  try {
    await apiClient.ping();

    return {
      success: true,
    };
  } catch {
    return {
      success: false,
    };
  }
}
