import {CompleteUserRegistrationOperationRequest} from "@/generated";
import {Dispatch, SetStateAction} from "react";
import {SuccessMessageParams} from "@/api/types";
import {successMessageHelper} from "@/api/success_message";
import {postCompleteRegistrationAPI} from "@/api/complete-registration/api";
import {unexpectedErrorMessage} from "@/api/error_message";

export async function postCompleteRegistrationHelper(
    params: CompleteUserRegistrationOperationRequest,
    setError: Dispatch<SetStateAction<string | null>>,
    successParams: SuccessMessageParams,
): Promise<void> {
    try {
        const data = await postCompleteRegistrationAPI(params);

        if (data.success) {
            successMessageHelper(successParams);
        } else {
            setError(data.message);
        }
    } catch {
        setError(unexpectedErrorMessage);
    }
}
