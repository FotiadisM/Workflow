import { Fragment } from "react";
import { Menu, Transition } from "@headlessui/react";
import { classNames } from "@/src/util";
import { useRouter } from "next/router";
import { useAuth } from "../auth/AuthRoute";
import { serverURI } from "@/src/api/url";

export const ProfileOptions: React.FC = () => {
  const router = useRouter();
  const auth = useAuth();

  return (
    <Menu as="div" className="ml-5">
      {({ open }) => (
        <>
          <Menu.Button className="bg-purple-800 flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-purple-100 focus:ring-white">
            <span className="sr-only">Open user menu</span>
            <img
              className="h-8 w-8 rounded-full"
              src={serverURI + "/static/" + auth?.user?.profile_pic}
              alt="profile-picture"
            />
          </Menu.Button>
          <Transition
            show={open}
            as={Fragment}
            enter="transition ease-out duration-100"
            enterFrom="transform opacity-0 scale-95"
            enterTo="transform opacity-100 scale-100"
            leave="transition ease-in duration-75"
            leaveFrom="transform opacity-100 scale-100"
            leaveTo="transform opacity-0 scale-95"
          >
            <Menu.Items
              static
              className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none z-10"
            >
              <Menu.Item>
                {({ active }) => (
                  <div
                    className={classNames(
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-700 hover:bg-purple-800 hover:text-white"
                    )}
                    onClick={() => {
                      router.push(`/user/${auth?.user?.id}`);
                    }}
                  >
                    Profile
                  </div>
                )}
              </Menu.Item>
              <Menu.Item>
                {({ active }) => (
                  <div
                    className={classNames(
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-700 hover:bg-purple-800 hover:text-white"
                    )}
                    onClick={() => {
                      router.push("/settings");
                    }}
                  >
                    Settings
                  </div>
                )}
              </Menu.Item>
              <Menu.Item>
                {({ active }) => (
                  <div
                    className={classNames(
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-700 hover:bg-purple-800 hover:text-white"
                    )}
                    onClick={() => {
                      auth?.signOut();
                    }}
                  >
                    Sign Out
                  </div>
                )}
              </Menu.Item>
            </Menu.Items>
          </Transition>
        </>
      )}
    </Menu>
  );
};
