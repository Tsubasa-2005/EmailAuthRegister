import {LoginOperationRequest} from "@/generated";

export async function postLoginAPI(
    requestParams: LoginOperationRequest
): Promise<{
    success: boolean;
    message: string;
}> {
    try{
        const res = await fetch("api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: requestParams.loginRequest.email, password: requestParams.loginRequest.password }),
        });

        if (res.ok) {
            return {
                success: true,
                message: '',
            }
        } else {
            const errorData = await res.json();
            return {
                success: false,
                message: errorData.message,
            }
        }
    } catch {
        return {
            success: false,
            message: 'ログイン中にエラーが発生しました',
        }
    }
}