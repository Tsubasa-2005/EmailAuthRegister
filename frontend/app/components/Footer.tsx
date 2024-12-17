import { GithubIcon, LinkedinIcon, MailIcon, TwitterIcon } from 'lucide-react';

export default function Footer() {
  return (
    <footer className="bg-white border-t border-gray-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="md:flex md:justify-between md:items-center">
          {/* Logo and Description */}
          <div className="mb-4 md:mb-0 flex items-center">
            <MailIcon size={32} className="text-blue-500 mr-3" />
            <div>
              <h3 className="text-xl font-bold text-gray-800">VerifyApp</h3>
              <p className="text-sm text-gray-500">
                Simple and secure email verification
              </p>
            </div>
          </div>

          {/* Social Links */}
          <div className="flex space-x-4">
            <a
              href="#"
              className="text-gray-400 hover:text-gray-600
              transition-colors group"
            >
              <GithubIcon size={24} className="group-hover:text-black" />
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-gray-600
              transition-colors group"
            >
              <TwitterIcon size={24} className="group-hover:text-blue-400" />
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-gray-600
              transition-colors group"
            >
              <LinkedinIcon size={24} className="group-hover:text-blue-600" />
            </a>
          </div>
        </div>

        {/* Copyright */}
        <div className="mt-6 pt-4 border-t border-gray-200 text-center">
          <p className="text-sm text-gray-500">
            © {new Date().getFullYear()} VerifyApp. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
}
