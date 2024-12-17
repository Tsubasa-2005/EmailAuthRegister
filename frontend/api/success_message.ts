import { SuccessMessageParams } from '@/api/types';

export function successMessageHelper({
  setSuccessMessage,
  message,
  timeout = 3000,
}: SuccessMessageParams): void {
  setSuccessMessage(message);

  setTimeout(() => {
    setSuccessMessage(null);
  }, timeout);
}
