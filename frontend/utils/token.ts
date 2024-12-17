import jwt from 'jsonwebtoken';

export interface UserTokenPayload extends jwt.JwtPayload {
    id: number;
    name: string;
}

export function parseToken(token: string): UserTokenPayload | null {
    try {
        const decoded = jwt.decode(token);

        if (!decoded || typeof decoded === 'string') {
            console.error("Invalid token");
            return null;
        }

        return decoded as UserTokenPayload;
    } catch (error) {
        console.error("Error decoding token:", error);
        return null;
    }
}