import { GetAllUsersRequest, User } from '@/generated';
import { Dispatch, SetStateAction } from 'react';
import { unexpectedErrorMessage } from '@/api/error_message';
import { getUsersAPI } from '@/api/users/api';

export async function getUsersHelper(
  setUsers: Dispatch<SetStateAction<User[]>>,
  setTotalPage: Dispatch<SetStateAction<number>>,
  params: GetAllUsersRequest,
  setError: Dispatch<SetStateAction<string | null>>
): Promise<void> {
  try {
    const data = await getUsersAPI(params);

    if (data.success) {
      setUsers(data.users);
      setTotalPage(data.totalPage);
    } else {
      setError(data.message);
    }
  } catch {
    setError(unexpectedErrorMessage);
  }
}
