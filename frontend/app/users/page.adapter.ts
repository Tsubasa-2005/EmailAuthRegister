import { useState, useEffect, useCallback } from 'react';
import { User, GetAllUsersRequest } from '@/generated';
import { getUsersHelper } from '@/api/users/getHelper';

export default function useUsersAdapter() {
  const [users, setUsers] = useState<User[]>([]);
  const [totalPages, setTotalPages] = useState<number>(1);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const fetchUsers = useCallback(async (params: GetAllUsersRequest) => {
    setIsLoading(true);
    setError(null);

    await getUsersHelper(setUsers, setTotalPages, params, setError);

    setIsLoading(false);
  }, []);

  const changePage = (page: number) => {
    if (page > 0 && page <= totalPages) {
      setCurrentPage(page);
      fetchUsers({ page }).then((r) => r);
    }
  };

  // 初回データ取得
  useEffect(() => {
    fetchUsers({ page: 1 }).then((r) => r);
  }, [fetchUsers]);

  // リロード用
  const refetch = () => {
    fetchUsers({ page: currentPage }).then((r) => r);
  };

  return {
    users,
    error,
    isLoading,
    totalPages,
    currentPage,
    changePage,
    refetch,
  };
}
