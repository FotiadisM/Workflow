import { useEffect, useState } from "react";
import MessagingBox from "@/components/conversation/MessagingBox";
import MessagingUser from "@/components/conversation/MessagingUser";
import Navbar from "@/components/navbar/Navbar";
import { AuthRoute, useAuth } from "@/components/auth/AuthRoute";
import { Conversation } from "@/src/types/conversation";
import { serverURI, serverWSURI } from "@/src/api/url";
import { User } from "@/src/types/user";

export default function Conversations() {
  const auth = useAuth();
  const [curUser, setCurUser] = useState<number>(0);
  const [convs, setConvs] = useState<Conversation[]>([]);
  const [currPerp, setCurrPerp] = useState<User | null>(null);

  useEffect(() => {
    if (auth !== null) {
      if (auth.user !== null) {
        fetch(serverURI + "/users/connections/" + auth.user.id)
          .then((res) => res.json())
          .then((data) => {
            setConvs(data.connections);
          })
          .catch((err) => console.log("failed to fetch connections: ", err));
      }
    }
  }, []);

  const [wsInstance, setWsInstance] = useState<WebSocket | null>(null);
  useEffect(() => {
    if (window !== undefined) {
      if (auth !== null) {
        if (auth.user !== null) {
          const ws = new WebSocket(serverWSURI + "/ws/" + auth.user.id);
          setWsInstance(ws);
        }
      }
    }

    return () => {
      if (wsInstance?.readyState !== 3) {
        wsInstance?.close();
      }
    };
  }, []);

  const tmpFunc = () => {
    if (convs === undefined || convs.length === 0 || wsInstance === null) {
      return <div>You have no connections to message</div>;
    } else {
      return (
        <div className="flex" style={{ width: "70vw", maxWidth: "1120px" }}>
          <div
            className="pr-7 border-r space-y-1"
            style={{ flexBasis: "100%", maxWidth: "270px" }}
          >
            {convs.map((c, i) => (
              <MessagingUser
                key={c.conn_id}
                user_id={c.user_id}
                current={convs[curUser].user_id === c.user_id}
                setCurUser={() => setCurUser(i)}
                setCurrPerp={setCurrPerp}
              />
            ))}
          </div>
          <div style={{ flexShrink: 2 }}>
            <MessagingBox
              perp={currPerp}
              conn_id={convs[curUser].conn_id}
              user_id={convs[curUser].user_id}
              wsInstance={wsInstance}
            />
          </div>
        </div>
      );
    }
  };

  return (
    <AuthRoute>
      <Navbar />
      <main className="flex justify-center py-10">{tmpFunc()}</main>
    </AuthRoute>
  );
}
