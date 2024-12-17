import { postSendVerificationAPI } from '@/api/send-verification/api';
import { SuccessMessageParams } from '@/api/types';
import { Dispatch, SetStateAction } from 'react';
import { SendEmailVerificationOperationRequest } from '@/generated';
import { successMessageHelper } from '@/api/success_message';
import { unexpectedErrorMessage } from '@/api/error_message';

export async function postSendVerificationHelper(
  params: SendEmailVerificationOperationRequest,
  setError: Dispatch<SetStateAction<string | null>>,
  successParams: SuccessMessageParams
): Promise<void> {
  try {
    const data = await postSendVerificationAPI(params);

    if (data.success) {
      successMessageHelper(successParams);
    } else {
      setError(data.message);
    }
  } catch {
    setError(unexpectedErrorMessage);
  }
}
