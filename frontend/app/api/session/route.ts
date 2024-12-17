import {NextResponse} from "next/server";
import {setCookie} from "@/utils/cookie";

export async function POST(request: Request) {
    const { token } = await request.json();

    console.log('token:', token);

    try {
        await setCookie(token);
        return NextResponse.json({success: true}, {status: 200});
    } catch {
        return NextResponse.json({error: 'Failed to set session'}, {status: 500});
    }
}