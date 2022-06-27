import { useAuth } from "@/components/auth/AuthRoute";
import { useRouter } from "next/router";
import React, { useState } from "react";

export default function SignIn() {
  const router = useRouter();
  const auth = useAuth();

  const [signUpInfo, setSignUpInfo] = useState<{
    first_name: string;
    last_name: string;
    email: string;
    company: string;
    position: string;
    password: string;
    password_repeat: string;
    profile: File | null;
  }>({
    first_name: "",
    last_name: "",
    email: "",
    company: "",
    position: "",
    password: "",
    password_repeat: "",
    profile: null,
  });

  const onSignUp = (e: React.FormEvent) => {
    e.preventDefault();

    if (signUpInfo.password !== signUpInfo.password_repeat) {
      window.alert("Passwords don't match");
      setSignUpInfo((o) => ({ ...o, password: "", password_repeat: "" }));
      return;
    }

    if (signUpInfo.profile !== null) {
      auth
        ?.signUp(
          signUpInfo.first_name,
          signUpInfo.last_name,
          signUpInfo.email,
          signUpInfo.company,
          signUpInfo.position,
          signUpInfo.password,
          signUpInfo.profile
        )
        .then((v) => {
          if (v !== null) {
            window.alert(v);
          }
          router.push("/home");
        })
        .catch((err) => console.log(err));
    }
  };

  return (
    <main className="h-screen flex flex-col items-center justify-center">
      <button
        className="rounded-md bg-purple-800 absolute transform -translate-x-1/2 -translate-y-1/2 p-3 hover:bg-purple-600 focus:outline-none"
        style={{ top: "15%", left: "15%" }}
        onClick={() => router.push("/")}
      >
        <span className="sr-only">Go back to home page</span>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          fill="currentColor"
          className="text-white h-8 w-8"
          viewBox="0 0 16 16"
        >
          <path
            fillRule="evenodd"
            d="M14.5 1.5a.5.5 0 0 1 .5.5v4.8a2.5 2.5 0 0 1-2.5 2.5H2.707l3.347 3.346a.5.5 0 0 1-.708.708l-4.2-4.2a.5.5 0 0 1 0-.708l4-4a.5.5 0 1 1 .708.708L2.707 8.3H12.5A1.5 1.5 0 0 0 14 6.8V2a.5.5 0 0 1 .5-.5z"
          />
        </svg>
      </button>
      <h1 className="text-6xl text-gray-700 mb-7">Sign Up</h1>
      <form
        className="rounded-xl border border-purple-800 space-y-6 px-7 py-4 shadow-xl"
        onSubmit={onSignUp}
      >
        <div>
          <label htmlFor="first_name" className="sr-only">
            First Name
          </label>
          <input
            id="first_name"
            name="first_name"
            type="text"
            autoComplete="given-name"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="First Name"
            value={signUpInfo.first_name}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, first_name: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="last_name" className="sr-only">
            Last Name
          </label>
          <input
            id="last_name"
            name="last_name"
            type="text"
            autoComplete="family-name"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Last address"
            value={signUpInfo.last_name}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, last_name: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="email-address" className="sr-only">
            Email address
          </label>
          <input
            id="email-address"
            name="email"
            type="email"
            autoComplete="email"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Email address"
            value={signUpInfo.email}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, email: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="company" className="sr-only">
            Company name
          </label>
          <input
            id="company"
            name="company"
            type="text"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Company name"
            value={signUpInfo.company}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, company: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="position" className="sr-only">
            Position name
          </label>
          <input
            id="position"
            name="position"
            type="text"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Position name"
            value={signUpInfo.position}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, position: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="password" className="sr-only">
            New Password
          </label>
          <input
            id="email-password"
            name="password"
            type="password"
            autoComplete="new-password"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Password"
            value={signUpInfo.password}
            onChange={(e) => {
              setSignUpInfo((o) => ({ ...o, password: e.target.value }));
            }}
          />
        </div>
        <div>
          <label htmlFor="password_repeat" className="sr-only">
            Repeat Password
          </label>
          <input
            id="password_repeat"
            name="password_repeat"
            type="password"
            autoComplete="new-password"
            required
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Repeat Password"
            value={signUpInfo.password_repeat}
            onChange={(e) => {
              setSignUpInfo((o) => ({
                ...o,
                password_repeat: e.target.value,
              }));
            }}
          />
        </div>
        <div>
          <label htmlFor="profile" className="sr-only">
            Image for profile picture
          </label>
          <input
            id="profile"
            name="profile"
            type="file"
            autoComplete="photo"
            required
            accept="image/*"
            className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            onChange={(e) => {
              setSignUpInfo((o) => ({
                ...o,
                profile: e.target.files !== null ? e.target.files[0] : null,
              }));
            }}
          />
        </div>
        <button
          type="submit"
          className="w-full btn px-2 py-2 text-white bg-purple-800"
        >
          Get started
        </button>
      </form>
    </main>
  );
}
