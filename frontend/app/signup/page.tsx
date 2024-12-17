'use client';

import Link from "next/link";
import React, { useState, useEffect } from "react";
import { useRouter,useSearchParams } from "next/navigation";
import useEmailVerificationAdapter from "./page.adapter";
import {ArrowRightIcon, CheckCircleIcon, MailIcon, TimerIcon, UserPlusIcon} from 'lucide-react';

export default function EmailVerificationPage() {
    const router = useRouter();

    const { error, isSubmitting, sendVerificationEmail, register, verifyToken, email, expired, resetError } = useEmailVerificationAdapter();
    const searchParams = useSearchParams();

    const[registerEmail, setRegisterEmail] = useState("");
    const [fullName, setFullName] = useState("");
    const [password, setPassword] = useState("");
    const [passError, setPassError] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [stage, setStage] = useState("email");

    useEffect(() => {
        const token = searchParams.get("id");
        if (token) {
            verifyToken({verifyEmailRequest: {token}}).then(r =>r);
            setStage("complete");
        }
    }, [searchParams, verifyToken]);

    const handleEmailSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        await sendVerificationEmail({sendEmailVerificationRequest: {email: registerEmail}});
        setStage("success");
    };

    const handleCompleteRegistration = async (e: React.FormEvent) => {
        e.preventDefault();

        if (password !== confirmPassword) {
            setPassError('Passwords do not match');
            return;
        }

        const token = searchParams.get("id");
        if (!token) return;
        await register({completeUserRegistrationRequest: {token: token, email: email, name: fullName, password: password}});

        if (!error) {
            router.push("/users");
        }
    };

    const handleNavigateToSignup = () => {
        setStage("email");
        resetError();
    };

    if (stage === "success") {
        return (
            <div
                className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-blue-100 p-4">
                <div className="w-full max-w-md bg-white rounded-2xl shadow-2xl border border-gray-100 overflow-hidden">
                    <div className="p-8 text-center">
                        <div className="flex justify-center mb-6">
                            <CheckCircleIcon
                                size={64}
                                className="text-green-500 animate-bounce"
                            />
                        </div>
                        <h2 className="text-3xl font-extrabold text-gray-900 mb-4 tracking-tight">
                            確認メール送信完了
                        </h2>
                        <p className="text-gray-600 mb-6 leading-relaxed">
                            ご登録のメールアドレスに確認メールをお送りしました。メール内のリンクをクリックして、サインアップ手続きを完了してください。
                        </p>
                        <Link
                            href="/"
                            className="w-full inline-block px-6 py-3 text-white font-semibold
            bg-gradient-to-r from-blue-500 to-purple-600
            rounded-lg shadow-md hover:from-blue-600 hover:to-purple-700
            transition-all duration-300 ease-in-out
            transform hover:-translate-y-1 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                        >
                            ホームに戻る
                        </Link>
                    </div>
                </div>
            </div>
        );
    }

    if (stage === "email") {
        return (
            <div
                className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-50 to-gray-100 p-4">
                <div className="w-full max-w-md bg-white rounded-2xl shadow-2xl border border-gray-100 overflow-hidden">
                    <div className="p-8">
                        <div className="flex items-center justify-center mb-6">
                            <MailIcon
                                size={48}
                                className="text-blue-500 mr-3 animate-pulse"
                            />
                            <h2 className="text-3xl font-bold text-gray-800">
                                Email Verification
                            </h2>
                        </div>

                        <form onSubmit={handleEmailSubmit} className="space-y-6">
                            <div>
                                <label
                                    htmlFor="email"
                                    className="block text-sm font-medium text-gray-700 mb-2"
                                >
                                    Email Address
                                </label>
                                <input
                                    type="email"
                                    id="email"
                                    value={registerEmail}
                                    onChange={(e) => setRegisterEmail(e.target.value)}
                                    required
                                    className="w-full px-4 py-3 border-2 border-gray-300 rounded-xl
                  focus:outline-none focus:ring-2 focus:ring-blue-500
                  focus:border-transparent
                  transition-all duration-300
                  placeholder-gray-400"
                                    placeholder="you@example.com"
                                />
                            </div>

                            {error && (
                                <p className="text-red-500 text-sm flex items-center">
                                    <span className="mr-2">❌</span>
                                    {error}
                                </p>
                            )}

                            <button
                                type="submit"
                                disabled={isSubmitting}
                                className="w-full py-3 px-4
                bg-gradient-to-r from-blue-500 to-purple-600
                text-white font-semibold rounded-xl
                hover:from-blue-600 hover:to-purple-700
                focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500
                transition-all duration-300
                flex items-center justify-center
                disabled:opacity-50 disabled:cursor-not-allowed
                transform hover:-translate-y-1 hover:scale-105"
                            >
                                {isSubmitting ? (
                                    <span className="animate-pulse">Sending...</span>
                                ) : (
                                    <>
                                        Send Verification
                                        <ArrowRightIcon size={20} className="ml-2"/>
                                    </>
                                )}
                            </button>
                        </form>

                        <div className="mt-6 text-center">
                            <p className="text-sm text-gray-600">
                                Already verified?{" "}
                                <Link
                                    href="/signin"
                                    className="text-blue-500 font-semibold hover:underline transition-colors"
                                >
                                    Sign In
                                </Link>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    if (expired) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-gray-50 p-4">
                <div className="w-full max-w-md bg-white rounded-2xl shadow-2xl p-8 text-center">
                    <div className="flex justify-center mb-6">
                        <TimerIcon size={64} className="text-yellow-500 animate-pulse" />
                    </div>
                    <h2 className="text-3xl font-extrabold text-gray-900 mb-4">有効期限が切れました</h2>
                    <p className="text-gray-600 mb-6">
                        認証リンクの有効期限が切れています。再度認証メールを送信して手続きを進めてください。
                    </p>
                    <Link
                        href="/signup"
                        onClick={handleNavigateToSignup}
                        className="w-full inline-block px-6 py-3 text-white font-semibold
                          bg-gradient-to-r from-blue-500 to-purple-600
                          rounded-lg shadow-md hover:from-blue-600 hover:to-purple-700
                          transition-all duration-300 ease-in-out
                          transform hover:-translate-y-1 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                    >
                        サインアップ画面に移動
                    </Link>
                </div>
            </div>
        );
    }

    // Complete Registration
    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-50 to-gray-100 p-4">
            <div className="w-full max-w-md bg-white rounded-2xl shadow-2xl border border-gray-100 overflow-hidden">
                <div className="p-8">
                    <div className="flex items-center justify-center mb-6">
                        <UserPlusIcon
                            size={48}
                            className="text-blue-500 mr-3 animate-pulse"
                        />
                        <h2 className="text-3xl font-bold text-gray-800">
                            Complete Registration
                        </h2>
                    </div>

                    <form onSubmit={handleCompleteRegistration} className="space-y-6">
                        {/* Email Field */}
                        <div>
                            <label
                                htmlFor="email"
                                className="block text-sm font-medium text-gray-700 mb-2"
                            >
                                Email Address
                            </label>
                            <input
                                type="email"
                                id="email"
                                value={email}
                                readOnly
                                className="w-full px-4 py-3 border-2 border-gray-300 rounded-xl
                  bg-gray-100 text-gray-500 cursor-not-allowed"
                                placeholder="you@example.com"
                            />
                        </div>

                        {/* Full Name Field */}
                        <div>
                            <label
                                htmlFor="fullName"
                                className="block text-sm font-medium text-gray-700 mb-2"
                            >
                                Full Name
                            </label>
                            <input
                                type="text"
                                id="fullName"
                                value={fullName}
                                onChange={(e) => setFullName(e.target.value)}
                                required
                                className="w-full px-4 py-3 border-2 border-gray-300 rounded-xl
                  focus:outline-none focus:ring-2 focus:ring-blue-500
                  focus:border-transparent
                  transition-all duration-300
                  placeholder-gray-400"
                                placeholder="John Doe"
                            />
                        </div>

                        {/* Password Field */}
                        <div>
                            <label
                                htmlFor="password"
                                className="block text-sm font-medium text-gray-700 mb-2"
                            >
                                Password
                            </label>
                            <input
                                type="password"
                                id="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                                className="w-full px-4 py-3 border-2 border-gray-300 rounded-xl
                  focus:outline-none focus:ring-2 focus:ring-blue-500
                  focus:border-transparent
                  transition-all duration-300
                  placeholder-gray-400"
                                placeholder="Create a strong password"
                            />
                        </div>

                        {/* Confirm Password Field */}
                        <div>
                            <label
                                htmlFor="confirmPassword"
                                className="block text-sm font-medium text-gray-700 mb-2"
                            >
                                Confirm Password
                            </label>
                            <input
                                type="password"
                                id="confirmPassword"
                                value={confirmPassword}
                                onChange={(e) => setConfirmPassword(e.target.value)}
                                required
                                className="w-full px-4 py-3 border-2 border-gray-300 rounded-xl
                  focus:outline-none focus:ring-2 focus:ring-blue-500
                  focus:border-transparent
                  transition-all duration-300
                  placeholder-gray-400"
                                placeholder="Repeat your password"
                            />
                        </div>

                        {passError && (
                            <p className="text-red-500 text-sm flex items-center">
                                <span className="mr-2">❌</span>
                                {passError}
                            </p>
                        )}

                        {/* Error Message */}
                        {error && (
                            <p className="text-red-500 text-sm flex items-center">
                                <span className="mr-2">❌</span>
                                {error}
                            </p>
                        )}

                        {/* Terms Checkbox */}
                        <div className="flex items-center">
                            <input
                                id="terms"
                                type="checkbox"
                                required
                                className="h-4 w-4 text-blue-500 focus:ring-blue-500 border-gray-300 rounded"
                            />
                            <label
                                htmlFor="terms"
                                className="ml-2 block text-sm text-gray-600"
                            >
                                I agree to the{' '}
                                <Link
                                    href="/terms"
                                    className="text-blue-500 hover:underline"
                                >
                                    Terms and Conditions
                                </Link>
                            </label>
                        </div>

                        {/* Submit Button */}
                        <button
                            type="submit"
                            disabled={isSubmitting}
                            className="w-full py-3 px-4
                bg-gradient-to-r from-blue-500 to-purple-600
                text-white font-semibold rounded-xl
                hover:from-blue-600 hover:to-purple-700
                focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500
                transition-all duration-300
                flex items-center justify-center
                disabled:opacity-50 disabled:cursor-not-allowed
                transform hover:-translate-y-1 hover:scale-105"
                        >
                            {isSubmitting ? (
                                <span className="animate-pulse">Creating Account...</span>
                            ) : (
                                <>
                                    Create Account
                                    <ArrowRightIcon size={20} className="ml-2"/>
                                </>
                            )}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}