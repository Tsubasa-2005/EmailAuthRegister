import { Dispatch, SetStateAction } from 'react';

/**
 * If no timeout parameter is specified, 3000 ms (3 seconds) is adopted as default.
 */
export interface SuccessMessageParams {
  setSuccessMessage: Dispatch<SetStateAction<string | null>>;
  message: string;
  timeout?: number;
}
