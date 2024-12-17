export const unexpectedErrorMessage = '予期しないエラーが発生しました。もう一度やり直してください。';

export async function handleError(
  error: unknown,
): Promise<{ message: string; status: number }> {
  if (error instanceof Error && 'response' in error) {
    const responseError = error as { response: Response };
    const errorResponseBody = await responseError.response.text();
    const parsedBody = JSON.parse(errorResponseBody);
    return {
      message: parsedBody.error_message,
      status: responseError.response.status,
    };
  }
  return { message: '予期しないエラーが発生しました。', status: 500 };
}
