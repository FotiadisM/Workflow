import { useEffect, useRef, useState } from "react";
import { classNames } from "@/src/util";
import { useLayoutEffect } from "@/src/useIsomorphicLayoutEffect";
import { User } from "@/src/types/user";
import { serverURI } from "@/src/api/url";
import { Message } from "@/src/types/conversation";
import { useAuth } from "../auth/AuthRoute";
import Link from "next/link";

interface MessagingBoxProps {
  conn_id: string;
  user_id: string;
  perp: User | null;
  wsInstance: WebSocket;
}

export default function MessagingBox({
  conn_id,
  user_id,
  perp,
  wsInstance,
}: MessagingBoxProps) {
  const auth = useAuth();
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    fetch(serverURI + "/conversations/messages/" + conn_id)
      .then((res) => res.json())
      .then((data) => setMessages(data.messages))
      .catch((err) => console.log(err));
  }, [conn_id]);

  useEffect(() => {
    wsInstance.addEventListener("message", (e) => {
      let msg = JSON.parse(e.data);
      if (msg.conn_id === conn_id) {
        setMessages((m) => {
          return [
            ...m,
            {
              id: msg.id,
              senter_id: msg.senter_id,
              text: msg.text,
              time: msg.time,
            },
          ];
        });
      }
    });

    return () => {
      wsInstance.removeEventListener("message", () => {});
    };
  }, []);

  const [curMsg, setCurMsg] = useState<string>("");
  const msgBox = useRef<HTMLDivElement>(null);
  const onMessageSent = (e: React.FormEvent) => {
    e.preventDefault();

    if (curMsg !== undefined) {
      if (auth !== null) {
        if (auth.user !== null) {
          fetch(serverURI + "/conversations/messages", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              conv_id: conn_id,
              senter_id: auth.user.id,
              text: curMsg,
            }),
          })
            .then((res) => res.json())
            .then((data) => {
              setMessages((m) => {
                return [
                  ...m,
                  {
                    id: data.id,
                    senter_id: auth.user!.id,
                    text: curMsg,
                    time: data.time,
                  },
                ];
              });
              setCurMsg("");
            })
            .catch((err) => console.log(err));
        }
      }
    }
  };

  useLayoutEffect(() => {
    if (msgBox !== null) {
      if (msgBox.current !== null) {
        msgBox.current.scrollTop = msgBox.current.scrollHeight;
      }
    }
  }, [messages]);

  if (perp === null) {
    return null;
  }

  return (
    <div className="ml-7" style={{ minWidth: "800px" }}>
      <div>
        <div className="flex items-center justify-between">
          <div className="flex items-center pb-4">
            <img
              className="rounded-full h-14 w-14 mr-4"
              src={serverURI + "/static/" + perp.profile_pic}
            />
            <Link href={"/user/" + perp.id}>
              <a className="text-2xl font-semibold hover:text-purple-800">
                {perp.f_name} {perp.l_name}
              </a>
            </Link>
          </div>
        </div>
        <hr />
      </div>

      <div className="pt-4">
        <div
          ref={msgBox}
          className="flex flex-col space-y-2 px-5 overflow-y-auto"
          style={{ height: "420px" }}
        >
          {messages.map((m) => (
            <div
              key={m.id}
              className={classNames(
                m.senter_id === user_id
                  ? "bg-gray-100 text-gray-800"
                  : "bg-purple-800 text-white",
                "px-2 py-1 rounded-md"
              )}
              style={{
                maxWidth: "40%",
                alignSelf: m.senter_id === user_id ? "flex-start" : "flex-end",
              }}
            >
              {m.text}
            </div>
          ))}
        </div>
        <form className="flex items-center pt-4" onSubmit={onMessageSent}>
          <input
            className="w-full px-3 py-1 bg-gray-100 rounded-md focus:outline-none"
            value={curMsg}
            onChange={(e) => setCurMsg(e.target.value)}
            placeholder="Send a message.."
            autoFocus
          />
          <button type="submit" className="focus:outline-none">
            <span className="sr-only">Send message</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="h-7 w-7 m-2 text-purple-800"
              viewBox="0 0 16 16"
            >
              <path d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z" />
            </svg>
          </button>
        </form>
      </div>
    </div>
  );
}
