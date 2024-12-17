import {CompleteUserRegistrationOperationRequest} from "@/generated";
import apiClient from "@/api/client";
import {handleError} from "@/api/error_message";

export async function postCompleteRegistrationAPI(
    requestParams: CompleteUserRegistrationOperationRequest
): Promise<{
    success: boolean;
    token: string;
    message: string;
}> {
    try {
        const response = await apiClient.completeUserRegistration(requestParams);

        return {
            success: true,
            token: response.token,
            message: "",
        };
    } catch (error) {
        const { message } = await handleError(error);
        return {
            success: false,
            token: "",
            message,
        };
    }
}