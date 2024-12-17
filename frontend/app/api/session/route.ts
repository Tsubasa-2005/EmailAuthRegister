import {cookies} from "next/headers";
import {NextResponse} from "next/server";

export async function GET() {
    try {
        const cookieStore = await cookies(); // `cookies()` を呼び出し
        const sessionToken = cookieStore.get('session'); // `await` を使用

        if (!sessionToken) {
            return NextResponse.json(
                { error: 'No session found' },
                { status: 401 }
            );
        }

        return NextResponse.json(
            { token: sessionToken.value },
            { status: 200 }
        );
    } catch (error) {
        console.error('Failed to retrieve session:', error);
        return NextResponse.json(
            { error: 'Failed to retrieve session' },
            { status: 500 },
        );
    }
}