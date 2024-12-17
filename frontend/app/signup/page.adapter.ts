import { useState, useCallback } from 'react';
import {
  CompleteUserRegistrationOperationRequest,
  SendEmailVerificationOperationRequest,
  VerifyEmailOperationRequest,
} from '@/generated';
import { postSendVerificationHelper } from '@/api/send-verification/postHelper';
import { postCompleteRegistrationHelper } from '@/api/complete-registration/postHelper';
import { postVerificationHelper } from '@/api/verify-email/postHelper';

export default function useEmailVerificationAndRegisterAdapter() {
  const [expired, setExpired] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);
  const [email, setEmail] = useState<string>('');

  const sendVerificationEmail = useCallback(
    async (params: SendEmailVerificationOperationRequest) => {
      setError(null);
      setIsSubmitting(true);

      await postSendVerificationHelper(params, setError, {
        setSuccessMessage,
        message: 'Verification email sent successfully',
      });
      setIsSubmitting(false);
    },
    []
  );

  const verifyToken = useCallback(
    async (params: VerifyEmailOperationRequest) => {
      setError(null);
      setIsSubmitting(true);

      await postVerificationHelper(params, setEmail, setExpired, setError, {
        setSuccessMessage,
        message: 'Email verified successfully',
      });
      setIsSubmitting(false);
    },
    []
  );

  const register = useCallback(
    async (params: CompleteUserRegistrationOperationRequest) => {
      setError(null);
      setIsSubmitting(true);

      await postCompleteRegistrationHelper(params, setError, {
        setSuccessMessage,
        message: 'Registration completed successfully',
      });
      setIsSubmitting(false);
    },
    []
  );

  const resetError = useCallback(() => {
    setError(null);
  }, []);

  return {
    error,
    isSubmitting,
    sendVerificationEmail,
    register,
    successMessage,
    verifyToken,
    email,
    expired,
    resetError,
  };
}
