import { useState, useCallback } from "react";
import { LoginOperationRequest } from "@/generated";
import { postLoginHelper } from "@/api/login/postHelper";

export default function useLoginAdapter() {
    const [error, setError] = useState<string | null>(null);
    const [isSubmitting, setIsSubmitting] = useState(false);

    const login = useCallback(
        async (params: LoginOperationRequest) => {
            setError(null);
            setIsSubmitting(true);

            await postLoginHelper(params, setError);

            setIsSubmitting(false);
        },
        []
    );

    return {
        error,
        isSubmitting,
        login
    };
}