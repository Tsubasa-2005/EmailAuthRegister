import { VerifyEmailOperationRequest } from '@/generated';
import apiClient from '@/api/client';
import { handleError } from '@/api/error_message';

export async function postVerifyEmailAPI(
  requestParams: VerifyEmailOperationRequest
): Promise<{
  success: boolean;
  email: string;
  message: string;
  status: number;
}> {
  try {
    const response = await apiClient.verifyEmail(requestParams);

    return {
      success: true,
      email: response.email,
      message: '',
      status: 200,
    };
  } catch (error) {
    const { message, status } = await handleError(error);
    return {
      success: false,
      email: '',
      message,
      status,
    };
  }
}
