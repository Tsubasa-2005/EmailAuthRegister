import Link from 'next/link';

export default function AuthLanding() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)] bg-gray-50 dark:bg-gray-900">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start text-center sm:text-left">
        <h1 className="text-2xl sm:text-3xl font-bold text-gray-800 dark:text-gray-200">
          Welcome to Our Platform
        </h1>

        <p className="text-gray-600 dark:text-gray-400 max-w-md">
          Join our community and start your journey. Whether you&apos;re new or
          returning, we&apos;re excited to have you!
        </p>

        <div className="flex gap-4 items-center flex-col sm:flex-row">
          <Link
            href="/signin"
            className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-gray-100 dark:hover:bg-gray-800 text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
          >
            Sign In
          </Link>

          <Link
            href="/signup"
            className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-black text-white dark:bg-white dark:text-black gap-2 hover:bg-gray-800 dark:hover:bg-gray-200 text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5"
          >
            Create Account
          </Link>
        </div>
      </main>
    </div>
  );
}
