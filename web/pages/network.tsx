import { useEffect, useState } from "react";
import Navbar from "@/components/navbar/Navbar";
import { useRouter } from "next/router";
import { AuthRoute, useAuth } from "@/components/auth/AuthRoute";
import { Conversation } from "@/src/types/conversation";
import { User } from "@/src/types/user";
import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";

interface NetworkUserProps {
  user: User;
}

const NetworkUser: React.FC<NetworkUserProps> = ({ user }) => {
  const router = useRouter();

  return (
    <div className="flex items-center justify-between px-3 py-2 border rounded-lg">
      <div className="flex items-center space-x-8 ml-5">
        <img
          className="h-12 w-12 rounded-full ring ring-offset-4 ring-purple-800 m-1"
          src={serverURI + "/static/" + user.profile_pic}
          alt="profile-picture"
        />
        <div className="text-xl">
          {user.f_name} {user.l_name}
        </div>
        <div className="flex items-center justify-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            className="text-gray-600 h-6 w-6"
            viewBox="0 0 16 16"
          >
            <path d="M6.5 1A1.5 1.5 0 0 0 5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v8A1.5 1.5 0 0 0 1.5 14h13a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 14.5 3H11v-.5A1.5 1.5 0 0 0 9.5 1h-3zm0 1h3a.5.5 0 0 1 .5.5V3H6v-.5a.5.5 0 0 1 .5-.5zm1.886 6.914L15 7.151V12.5a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5V7.15l6.614 1.764a1.5 1.5 0 0 0 .772 0zM1.5 4h13a.5.5 0 0 1 .5.5v1.616L8.129 7.948a.5.5 0 0 1-.258 0L1 6.116V4.5a.5.5 0 0 1 .5-.5z" />
          </svg>
          <div className="ml-2">{user.position}</div>
        </div>
        <div className="flex items-center justify-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            className="text-gray-600 h-6 w-6"
            viewBox="0 0 16 16"
          >
            <path
              fillRule="evenodd"
              d="M14.763.075A.5.5 0 0 1 15 .5v15a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5V14h-1v1.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5V10a.5.5 0 0 1 .342-.474L6 7.64V4.5a.5.5 0 0 1 .276-.447l8-4a.5.5 0 0 1 .487.022zM6 8.694 1 10.36V15h5V8.694zM7 15h2v-1.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5V15h2V1.309l-7 3.5V15z"
            />
            <path d="M2 11h1v1H2v-1zm2 0h1v1H4v-1zm-2 2h1v1H2v-1zm2 0h1v1H4v-1zm4-4h1v1H8V9zm2 0h1v1h-1V9zm-2 2h1v1H8v-1zm2 0h1v1h-1v-1zm2-2h1v1h-1V9zm0 2h1v1h-1v-1zM8 7h1v1H8V7zm2 0h1v1h-1V7zm2 0h1v1h-1V7zM8 5h1v1H8V5zm2 0h1v1h-1V5zm2 0h1v1h-1V5zm0-2h1v1h-1V3z" />
          </svg>
          <div className="ml-2">{user.company}</div>
        </div>
      </div>
      <button
        className="btn px-2 py-1 mr-5 bg-purple-800 hover:bg-purple-700 active:bg-purple-900"
        onClick={() => {
          router.push(`/user/${user.id}`);
        }}
      >
        <span className="sr-only">Visit user's profile</span>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="currentColor"
          className="text-white h-8 w-8"
          viewBox="0 0 16 16"
        >
          <path
            fillRule="evenodd"
            d="M4 8a.5.5 0 0 1 .5-.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5A.5.5 0 0 1 4 8z"
          />
        </svg>
      </button>
    </div>
  );
};

interface NetworkUserWraperProps {
  user_id: string;
}

const NetworkUserWraper: React.FC<NetworkUserWraperProps> = ({ user_id }) => {
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

  return <NetworkUser user={user} />;
};

export default function Network() {
  const auth = useAuth();

  const [connections, setConnections] = useState<Conversation[]>([]);
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    if (auth !== null) {
      if (auth.user !== null) {
        fetch(serverURI + "/users/connections/" + auth.user.id)
          .then((res) => res.json())
          .then((data) => {
            setConnections(data.connections);
          })
          .catch((err) => console.log("failed to fetch connections: ", err));
      }
    }

    fetch(serverURI + "/users")
      .then((res) => res.json())
      .then((data) => {
        setUsers(data.users);
      })
      .catch((err) => console.log(err));
  }, []);

  const [searchText, setSearchText] = useState<string>("");
  useEffect(() => {
    if (searchText === "") {
    } else {
      // fetch users that match text
      // setNetworkUsers();
    }
  }, [searchText]);

  const evalUser = (user_id: string): boolean => {
    if (auth !== null)
      if (auth.user !== null) if (auth.user.id === user_id) return false;

    if (connections !== undefined) {
      for (let i = 0; i < connections.length; i++) {
        if (connections[i].user_id === user_id) return false;
      }
    }

    return true;
  };

  return (
    <AuthRoute>
      <Navbar />
      <main className="flex justify-center">
        <div className="my-10">
          <input
            type="text"
            className="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
            placeholder="Search Workflow"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
          />
          <div
            className="space-y-2 mt-2 overflow-y-auto"
            style={{ width: "70vw", maxWidth: "1120px" }}
          >
            {connections === undefined
              ? null
              : [
                  connections.map((c) => (
                    <NetworkUserWraper key={c.conn_id} user_id={c.user_id} />
                  )),
                ]}
            {users.length === 0 ? null : (
              <div>
                <hr className="my-8 mx-12 border-2 border-purple-800" />
              </div>
            )}
            {users.map((u) => {
              if (evalUser(u.id)) {
                return <NetworkUser key={u.id} user={u} />;
              }
              return null;
            })}
          </div>
        </div>
      </main>
    </AuthRoute>
  );
}
