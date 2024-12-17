import apiClient from '@/api/client';
import { SendEmailVerificationOperationRequest } from '@/generated';
import { handleError } from '@/api/error_message';

export async function postSendVerificationAPI(
  requestParams: SendEmailVerificationOperationRequest
): Promise<{
  success: boolean;
  message: string;
}> {
  try {
    await apiClient.sendEmailVerification(requestParams);

    return {
      success: true,
      message: '',
    };
  } catch (error) {
    const { message } = await handleError(error);
    return {
      success: false,
      message,
    };
  }
}
