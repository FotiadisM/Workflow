import Link from "next/link";
import { useRouter } from "next/router";
import { classNames } from "@/src/util";
import { useLayoutEffect } from "@/src/useIsomorphicLayoutEffect";
import { ProfileOptions } from "./ProfileOptions";
import { Notifications } from "./Notifications";
import { useAuth } from "../auth/AuthRoute";

const navigation: { name: string; href: string }[] = [
  { name: "Home", href: "/home" },
  { name: "Network", href: "/network" },
  { name: "Messages", href: "/conversations" },
  { name: "Job Openings", href: "/jobs" },
];

interface NavBarProps {
  getHeight?: (h: number) => void;
}

export default function Navbar({ getHeight }: NavBarProps) {
  const router = useRouter();
  const auth = useAuth();

  useLayoutEffect(() => {
    if (getHeight !== undefined) {
      const nav = document.getElementById("navbar");
      if (nav !== null) {
        getHeight(nav.offsetHeight);
      }
    }
  }, []);

  const returnAdminDashboardLink = () => {
    if (auth !== null)
      if (auth.user !== null)
        if (auth.user.role === "admin")
          return (
            <Link href="/dashboard">
              <a
                className={classNames(
                  router.route === "/dashboard"
                    ? "bg-purple-300 text-gray-700"
                    : "text-gray-100 hover:bg-purple-600",
                  "px-3 py-2 rounded-md text-sm font-medium"
                )}
                aria-current={
                  router.route === "/dashboard" ? "page" : undefined
                }
              >
                Admin Dashboard
              </a>
            </Link>
          );
  };

  return (
    <nav id="navbar" className="bg-purple-900">
      <div className="px-10 py-4">
        <div className="flex items-center justify-between">
          <div className="flex space-x-4">
            <img
              className="block h-8 w-auto texg-white"
              src="https://tailwindui.com/img/logos/workflow-logo-indigo-500-mark-white-text.svg"
              alt="Workflow"
            />
            {navigation.map((item) => (
              <Link key={item.name} href={item.href}>
                <a
                  key={item.name}
                  className={classNames(
                    router.route === item.href
                      ? "bg-purple-300 text-gray-700"
                      : "text-gray-100 hover:bg-purple-600",
                    "px-3 py-2 rounded-md text-sm font-medium"
                  )}
                  aria-current={router.route === item.href ? "page" : undefined}
                >
                  {item.name}
                </a>
              </Link>
            ))}
            {returnAdminDashboardLink()}
          </div>
          <div className="flex items-center">
            {auth === null || auth?.user === null ? null : (
              <>
                <Notifications />
                <ProfileOptions />
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}
