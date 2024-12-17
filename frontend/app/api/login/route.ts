import { NextResponse } from 'next/server';
import {LoginOperationRequest} from "@/generated";
import {setCookie} from "@/utils/cookie";
import apiClient from "@/app/api/client";

export async function POST(request: Request) {
    try {
        const { email, password } = await request.json();

        const requestBody: LoginOperationRequest = {
            loginRequest: { email, password }
        };

        // TODO: おそらくheaderのtokenをセットする処理がcode genの方であるはず
        const response = await apiClient.loginRaw(requestBody);

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
    } catch {
        return NextResponse.json({ message: "error" }, { status: 500 });
    }
}