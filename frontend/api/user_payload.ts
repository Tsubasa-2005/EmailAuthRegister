import {parseToken, UserTokenPayload} from "@/utils/token";

export async function fetchCurrentUser(): Promise<{
  success: boolean;
  user: UserTokenPayload;
  message: string;
}> {
  const response = await fetch('/api/session');
  if (!response.ok) {
    return {
      success: false,
      user: {} as UserTokenPayload,
      message: 'Failed to fetch user.',
    };
  }

  const data = await response.json();
  const token = data.token;

  const user = parseToken(token);
    if (!user) {
        return {
        success: false,
        user: {} as UserTokenPayload,
        message: 'Failed to parse token.',
        };
    }

  return {
    success: true,
    user: user,
    message: '',
  };
}
