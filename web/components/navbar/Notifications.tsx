import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";
import { Conversation } from "@/src/types/conversation";
import { User } from "@/src/types/user";
import { Menu, Transition } from "@headlessui/react";
import Link from "next/link";
import { Fragment, useEffect, useState } from "react";
import { useAuth } from "../auth/AuthRoute";

// interface Notification {
//   id: string;
//   type: "friend-request";
//   value: {
//     user_id: string;
//     time: string;
//   };
// }

interface FriendRequestProps {
  user_id: string;
  conn_id: string;
}

const FriendRequest: React.FC<FriendRequestProps> = ({ user_id, conn_id }) => {
  const [user, setUser] = useState<User | null>(null);
  useEffect(() => {
    fetchPerpetrator(user_id)
      .then((u) => {
        setUser(u);
      })
      .catch((err) => console.log(err));
  }, [user_id]);

  if (user === null) {
    return null;
  }

  const onRequestAccept = () => {
    (async () => {
      try {
        const res = await fetch(serverURI + "/users/connectionRequests", {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            conn_id: conn_id,
            accept: true,
          }),
        });

        if (res.status !== 200) {
          const data = await res.text();
          console.log("faled to accept request:", data);
        }
      } catch (err) {
        console.log(err);
      }
    })();
  };
  const onRequestReject = () => {
    (async () => {
      try {
        const res = await fetch(serverURI + "/users/connectionRequests", {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            conn_id: conn_id,
            accept: false,
          }),
        });

        if (res.status !== 200) {
          const data = await res.text();
          console.log("faled to accept request:", data);
        }
      } catch (err) {
        console.log(err);
      }
    })();
  };

  return (
    <div className="px-2 flex items-center">
      <div className="border-r pr-2">
        <img
          className="h-10 w-10 rounded-full"
          src={serverURI + "/static/" + user.profile_pic}
          alt="profile-picture"
        />
      </div>
      <div className="ml-2 px-3">
        <h4 className="text-sm">
          <Link href="/users/">
            <a className="text-purple-800 hover:underline">
              {user.f_name + " " + user.l_name}
            </a>
          </Link>{" "}
          Sent you a <br /> Connection request
        </h4>
        <div className="flex items-center space-x-3">
          <button
            className="btn px-2 py-1 text-sm bg-purple-800 text-white hover:bg-purple-900"
            onClick={onRequestAccept}
          >
            Accept
          </button>
          <button
            className="btn px-2 py-1 text-sm border border-purple-800 text-purple-800 hover:bg-purple-50"
            onClick={onRequestReject}
          >
            Reject
          </button>
        </div>
      </div>
    </div>
  );
};

interface NotificationsProps {}

export const Notifications: React.FC<NotificationsProps> = () => {
  const auth = useAuth();
  const [requests, setRequests] = useState<Conversation[]>([]);

  useEffect(() => {
    const interval = setInterval(() => {
      if (auth !== null)
        if (auth.user !== null)
          fetch(serverURI + "/users/connectionRequests/" + auth.user.id)
            .then((res) => res.json())
            .then((data) => setRequests(data.connections))
            .catch((err) =>
              console.log("faile to get connections requests:", err)
            );
    }, 3000);

    return () => {
      clearInterval(interval);
    };
  });

  return (
    <Menu as="div" className="ml-5">
      {({ open }) => (
        <>
          <Menu.Button className="relative p-1 text-gray-100 bg-purple-800 rounded-full hover:text-white focus:outline-none focus:ring-1 focus:ring-offset-2 focus:ring-offset-purple-100 focus:ring-white">
            <span className="sr-only">View notifications</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              viewBox="0 0 16 16"
            >
              <path d="M8 16a2 2 0 0 0 2-2H6a2 2 0 0 0 2 2zm.995-14.901a1 1 0 1 0-1.99 0A5.002 5.002 0 0 0 3 6c0 1.098-.5 6-2 7h14c-1.5-1-2-5.902-2-7 0-2.42-1.72-4.44-4.005-4.901z" />
            </svg>
            {requests === undefined
              ? null
              : [
                  requests.length === 0 ? null : (
                    <div className="absolute -top-4 -right-4 text-white rounded-full bg-red-400 h-6 w-6 text-center">
                      {requests.length}
                    </div>
                  ),
                ]}
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
              className="origin-top-right absolute right-0 mt-2 rounded-md space-y-3 shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none z-10"
            >
              {requests === undefined
                ? null
                : [
                    requests.map((r) => (
                      <Menu.Item key={r.conn_id}>
                        <FriendRequest
                          user_id={r.user_id}
                          conn_id={r.conn_id}
                        />
                      </Menu.Item>
                    )),
                  ]}
            </Menu.Items>
          </Transition>
        </>
      )}
    </Menu>
  );
};
