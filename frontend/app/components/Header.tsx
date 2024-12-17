'use client';

import Link from 'next/link';
import { MailIcon, UserIcon, UserPlusIcon } from 'lucide-react';
import React from 'react';

export default function Header() {
  return (
    <header className="bg-white shadow-md border-b border-gray-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center py-4">
          {/* Logo */}
          <Link href="/" className="flex items-center">
            <MailIcon size={32} className="text-blue-500 mr-2" />
            <span className="text-2xl font-bold text-gray-800">VerifyApp</span>
          </Link>

          {/* Navigation */}
          <nav className="flex items-center space-x-4">
            <Link
              href="/signup"
              className="text-gray-600 hover:text-blue-500
              flex items-center transition-colors group"
            >
              <UserPlusIcon
                size={20}
                className="mr-2 group-hover:text-blue-500 text-gray-400"
              />
              Sign Up
            </Link>
            <Link
              href="/signin"
              className="text-gray-600 hover:text-blue-500
              flex items-center transition-colors group"
            >
              <UserIcon
                size={20}
                className="mr-2 group-hover:text-blue-500 text-gray-400"
              />
              Sign In
            </Link>
          </nav>
        </div>
      </div>
    </header>
  );
}
