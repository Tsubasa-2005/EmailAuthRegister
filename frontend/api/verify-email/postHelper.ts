import { VerifyEmailOperationRequest} from "@/generated";
import {Dispatch, SetStateAction} from "react";
import {SuccessMessageParams} from "@/api/types";
import {successMessageHelper} from "@/api/success_message";
import {postVerifyEmailAPI} from "@/api/verify-email/api";
import {unexpectedErrorMessage} from "@/api/error_message";
import {unauthorized} from "next/navigation";
import {HTTP_STATUS} from "@/constants/httpStatus";

export async function postVerificationHelper(
    params: VerifyEmailOperationRequest,
    email: Dispatch<SetStateAction<string>>,
    expired: Dispatch<SetStateAction<boolean>>,
    setError: Dispatch<SetStateAction<string | null>>,
    successParams: SuccessMessageParams,
): Promise<void> {
    try {
        const data = await postVerifyEmailAPI(params);

        if (data.success) {
            email(data.email);
            successMessageHelper(successParams);
        } else {
            setError(data.message);
            if (data.status === HTTP_STATUS.UNAUTHORIZED) {
                expired(true);
            }
        }
    } catch {
        setError(unexpectedErrorMessage);
    }
}
