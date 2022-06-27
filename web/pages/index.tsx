import { useAuth } from "@/components/auth/AuthRoute";
import Link from "next/link";
import { useRouter } from "next/router";
import React, { useState } from "react";

const navigation: { name: string; href: string }[] = [
  { name: "Product", href: "#" },
  { name: "Features", href: "#" },
  { name: "Marketplace", href: "#" },
  { name: "About", href: "#" },
];

function Navbar() {
  const router = useRouter();

  return (
    <nav className="flex items-center justify-between px-10 py-8">
      <div className="flex items-center space-x-10">
        <img
          className="h-10 w-auto"
          src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg"
          alt="company logo"
        />
        {navigation.map((i) => (
          <Link key={i.name} href={i.href}>
            <a className="text-lg text-gray-500 hover:text-gray-800">
              {i.name}
            </a>
          </Link>
        ))}
      </div>
      <button
        className="btn px-2 py-1 border-2 border-purple-800 text-lg text-purple-800 hover:text-white hover:bg-purple-800 hover:shadow-xl active:bg-purple-900"
        onClick={() => router.push("/signup")}
      >
        Sign Up
      </button>
    </nav>
  );
}

export default function LandingPage() {
  const router = useRouter();
  const auth = useAuth();

  const [form, setForm] = useState<{ email: string; password: string }>({
    email: "",
    password: "",
  });

  const onSignIn = (e: React.FormEvent) => {
    e.preventDefault();
    (async () => {
      if (auth !== null) {
        const res = await auth.signIn(form.email, form.password);
        if (res === null) {
          router.push("/home");
        } else {
          window.alert("Credentials don't match");
        }
      }
    })();
  };

  return (
    <>
      <Navbar />
      <main className="flex pt-32">
        <div className="flex justify-center" style={{ flexBasis: "50%" }}>
          <div>
            <h1 className="text-5xl font-bold">
              Welcome to your
              <span className="block text-purple-800">
                professional community
              </span>
            </h1>
            <p className="text-xl text-gray-700 max-w-md mt-5">
              Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui
              lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat
              fugiat aliqua.
            </p>
            <div className="flex space-x-5 mt-6">
              <button className="btn px-4 py-3 bg-purple-800 text-white hover:bg-purple-700 hover:shadow-xl active:bg-purple-900">
                Search for a job
              </button>
              <button className="btn px-4 py-3 bg-purple-100 text-purple-800 hover:bg-purple-200 hover:shadow-xl active:bg-purple-300">
                Find a person you know
              </button>
            </div>
          </div>
        </div>
        <div
          className="flex flex-col justify-center items-center"
          style={{ flexBasis: "50%" }}
        >
          <h2 className="text-4xl font-semibold text-purple-800">Sign in</h2>
          <hr className="border" />
          <form className="mt-6 space-y-6 w-2/4" onSubmit={onSignIn}>
            <div className="rounded-md shadow-sm -space-y-px text-md">
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
                  value={form.email}
                  onChange={(e) =>
                    setForm((f) => ({ ...f, email: e.target.value }))
                  }
                />
              </div>
              <div>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
                  placeholder="Password"
                  value={form.password}
                  onChange={(e) =>
                    setForm((f) => ({ ...f, password: e.target.value }))
                  }
                />
              </div>
            </div>

            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <input
                  id="remember_me"
                  name="remember_me"
                  type="checkbox"
                  className="h-4 w-4 text-purple-800 focus:ring-indigo-500 border-gray-300 rounded"
                />
                <label
                  htmlFor="remember_me"
                  className="ml-2 block text-sm text-gray-900"
                >
                  Remember me
                </label>
              </div>

              <div className="text-sm">
                <Link href="#">
                  <a className="font-medium text-purple-800 hover:text-purple-500">
                    Forgot your password?
                  </a>
                </Link>
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="w-full btn py-2 px-4 border border-transparent text-sm font-medium text-white bg-purple-800 hover:bg-purple-700 focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
              >
                Sign in
              </button>
            </div>
          </form>
        </div>
      </main>
    </>
  );
}
