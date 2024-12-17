import { CompleteUserRegistrationOperationRequest } from '@/generated';

export async function postCompleteRegistrationAPI(
  requestParams: CompleteUserRegistrationOperationRequest
): Promise<{
  success: boolean;
  message: string;
}> {
  try {
    const res = await fetch('api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        token: requestParams.completeUserRegistrationRequest.token,
        email: requestParams.completeUserRegistrationRequest.email,
        password: requestParams.completeUserRegistrationRequest.password,
        name: requestParams.completeUserRegistrationRequest.name,
      }),
    });

    if (res.ok) {
      return {
        success: true,
        message: '',
      };
    } else {
      const errorData = await res.json();
      return {
        success: false,
        message: errorData.message,
      };
    }
  } catch {
    return {
      success: false,
      message: 'エラーが発生しました',
    };
  }
}
