import { NextResponse } from 'next/server';
import {CompleteUserRegistrationOperationRequest, LoginOperationRequest} from "@/generated";
import {setCookie} from "@/utils/cookie";
import apiClient from "@/app/api/client";

export async function POST(request: Request) {
    try {
        const { token, password, email, name } = await request.json();

        const requestBody: CompleteUserRegistrationOperationRequest = {
            completeUserRegistrationRequest: { token, password, email, name }
        };

        // TODO: おそらくheaderのtokenをセットする処理がcode genの方であるはず
        const response = await apiClient.completeUserRegistrationRaw(requestBody);

        if (response.raw.ok) {
            const token = response.raw.headers.get("set-cookie");

            if (token) {
                await setCookie(token);
                return NextResponse.json({ success: true }, { status: 200 });
            } else {
                return NextResponse.json({ error: 'トークンが存在しません。' }, { status: 500 });
            }
        } else {
            const errorData = await response.raw.json();
            return NextResponse.json({ error: errorData.error || 'ログインに失敗しました' }, { status: response.raw.status });
        }
    } catch (error: any) {
        return NextResponse.json({ message: "error" }, { status: error.response.status });
    }
}