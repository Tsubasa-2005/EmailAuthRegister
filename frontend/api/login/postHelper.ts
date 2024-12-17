import { LoginOperationRequest } from '@/generated';
import { Dispatch, SetStateAction } from 'react';
import { unexpectedErrorMessage } from '@/api/error_message';
import { postLoginAPI } from '@/api/login/api';

export async function postLoginHelper(
  params: LoginOperationRequest,
  setError: Dispatch<SetStateAction<string | null>>
): Promise<void> {
  try {
    const data = await postLoginAPI(params);

    if (!data.success) {
      setError(data.message);
    }
  } catch {
    setError(unexpectedErrorMessage);
  }
}
