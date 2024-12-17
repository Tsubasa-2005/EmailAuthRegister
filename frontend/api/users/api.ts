import { GetAllUsersRequest, User } from '@/generated';
import apiClient from '@/api/client';
import { handleError } from '@/api/error_message';

export async function getUsersAPI(requestParams: GetAllUsersRequest): Promise<{
  success: boolean;
  users: User[];
  totalPage: number;
  message: string;
}> {
  try {
    const response = await apiClient.getAllUsers(requestParams);

    return {
      success: true,
      users: response.users,
      totalPage: response.totalPage,
      message: '',
    };
  } catch (error: any) {
    const { message } = await handleError(error);
    return {
      success: false,
      users: [],
      totalPage: 0,
      message: message,
    };
  }
}
