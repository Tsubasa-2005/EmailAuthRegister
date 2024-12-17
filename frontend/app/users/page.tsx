'use client';

import useUsersAdapter from '@/app/users/page.adapter';

export default function UsersPage() {
  const {
    users,
    error,
    isLoading,
    totalPages,
    currentPage,
    changePage,
    refetch,
  } = useUsersAdapter();

  return (
    <div className="min-h-screen bg-gray-50 p-4">
      <div className="max-w-4xl mx-auto bg-white shadow-md rounded-lg p-6">
        <h1 className="text-2xl font-bold mb-4 text-center text-gray-800">
          User List
        </h1>

        {isLoading && <p className="text-gray-500 text-center">Loading...</p>}
        {error && <p className="text-red-500 text-center">{error}</p>}

        <button
          onClick={refetch}
          className="w-full mb-4 py-2 px-4 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600 transition"
        >
          Refresh Users
        </button>

        {!isLoading && !error && (
          <>
            <ul className="space-y-3">
              {users.map((user) => (
                <li
                  key={user.id}
                  className="p-3 border rounded-md shadow-sm hover:bg-gray-100 transition"
                >
                  <p className="text-gray-700">
                    <strong>ID:</strong> {user.id}
                  </p>
                  <p className="text-gray-700">
                    <strong>Name:</strong> {user.name}
                  </p>
                  <p className="text-gray-700">
                    <strong>Email:</strong> {user.email}
                  </p>
                </li>
              ))}
            </ul>

            {/* ページネーション */}
            <div className="flex justify-between items-center mt-6">
              <button
                onClick={() => changePage(currentPage - 1)}
                disabled={currentPage === 1}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded disabled:opacity-50"
              >
                Previous
              </button>

              <span className="text-gray-700">
                Page {currentPage} of {totalPages}
              </span>

              <button
                onClick={() => changePage(currentPage + 1)}
                disabled={currentPage === totalPages}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded disabled:opacity-50"
              >
                Next
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
