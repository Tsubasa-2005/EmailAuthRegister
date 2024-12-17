import 'server-only';
import { cookies } from 'next/headers';

export async function setCookie(token: string) {
  const expiresAt = new Date(Date.now() + 1 * 60 * 60 * 1000); // 1時間後の有効期限

  const cookieStore = await cookies(); // `cookies()` を取得
  cookieStore.set('session', token, {
    httpOnly: true,
    secure: true,
    expires: expiresAt,
  });
}
