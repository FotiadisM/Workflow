import { useAuth } from "@/components/auth/AuthRoute";
import { Post } from "@/components/homepage/Post";
import Navbar from "@/components/navbar/Navbar";
import { UserEducation } from "@/components/profile/education";
import { UserExperience } from "@/components/profile/experience";
import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";
import { Conversation } from "@/src/types/conversation";
import { Feed } from "@/src/types/posts";
import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";

export default function UserProfile() {
  const router = useRouter();
  const { id } = router.query;

  const auth = useAuth();
  const [profUser, setProfUser] = useState({
    id: "",
    f_name: "",
    l_name: "",
    email: "",
    company: "",
    position: "",
    profile_pic: "",
  });

  useEffect(() => {
    if (id !== undefined) {
      if (typeof id === "string") {
        if (id !== auth?.user?.id) {
          fetchPerpetrator(id).then((u) => {
            if (u !== null) {
              setProfUser(u);
            }
          });
        }
      }
    }
  }, [id]);

  const [isFriend, setIsFriend] = useState<{ conn_id: string } | false>(false);
  const [connections, setConnections] = useState<Conversation[]>([]);
  useEffect(() => {
    if (id !== undefined)
      fetch(serverURI + "/users/connections/" + id)
        .then((res) => res.json())
        .then((data) => {
          const conss: Conversation[] = data.connections;
          if (conss === undefined) return;
          if (auth !== null) {
            if (auth.user !== null) {
              if (auth.user.id !== id) {
                conss.forEach((c) => {
                  if (c.user_id === auth?.user?.id) {
                    setIsFriend({ conn_id: c.conn_id });
                  }
                });
              }
            }
          }
          setConnections(conss);
        })
        .catch((err) => console.log("failed to fetch connections: ", err));
  }, [id]);

  const [feed, setFeed] = useState<Feed[]>([]);
  useEffect(() => {
    (async () => {
      try {
        if (id !== undefined) {
          let from_id = "";
          if (auth !== null) if (auth.user !== null) from_id = auth.user.id;
          const res = await fetch(serverURI + "/posts/user", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              user_id: id,
              from_user_id: from_id,
            }),
          });

          if (res.status === 200) {
            const data = await res.json();
            console.log("DATA:", data);
            const posts = data.posts;
            if (posts === undefined) return;
            console.log("POSTS:", posts);
            let fe: Feed[] = [];
            for (let i = 0; i < posts.length; i++) {
              fe.push({
                id: "",
                type: "post",
                perpetator_id: "1",
                post_id: posts[i].id,
              });
            }
            setFeed(fe);
          }
        }
      } catch (err) {
        console.log(err);
      }
    })();
  }, [id]);

  const onConnectionPost = () => {
    if (auth !== null)
      if (auth.user !== null)
        fetch(serverURI + "/users/connections", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            user_id: auth.user.id,
            receiver_id: id,
          }),
        })
          .then((res) => res.json())
          .then((data) => console.log(data))
          .catch((err) =>
            console.log("failed to post connections request:", err)
          );
  };
  const onConnectionRemove = () => {
    console.log("remove friend");
  };

  const utilUserCheck = (v1: any, v2: any) => {
    return id === auth?.user?.id ? v1 : v2;
  };

  if (id === undefined) {
    return null;
  }

  if (Array.isArray(id)) {
    return null;
  }

  return (
    <>
      <Navbar />
      <main className="flex justify-center items-center flex-col py-10 px-44">
        <div className="w-full border-2 rounded-full p-4 flex pr-28">
          <img
            className="h-40 w-40 rounded-full mr-14"
            src={
              serverURI +
              "/static/" +
              utilUserCheck(auth?.user?.profile_pic, profUser.profile_pic)
            }
            alt="profile-picture"
          />
          <div className="pt-4 flex-1">
            <div className="flex items-end">
              <h1 className="text-4xl font-medium italic">
                {utilUserCheck(auth?.user?.f_name, profUser.f_name)}{" "}
                {utilUserCheck(auth?.user?.l_name, profUser.l_name)}
              </h1>
              <h3 className="pl-8 pb-1 text-xl text-gray-800 italic">
                {utilUserCheck(auth?.user?.position, profUser.position)}{" "}
                <span className="text-gray-600">@</span>{" "}
                {utilUserCheck(auth?.user?.company, profUser.company)}
              </h3>
            </div>
            <hr />
            <div className="pt-2 flex items-center text-gray-700">
              <p className="mr-2">{connections?.length} connections</p>
            </div>
            {isFriend === false ? (
              [
                auth?.user?.id === id ? null : (
                  <div className="pt-4 flex items-center">
                    <button
                      className="btn px-3 py-2 text-purple-800 border border-purple-800 hover:text-white hover:bg-purple-800 flex items-center"
                      onClick={onConnectionPost}
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        className="mr-2"
                        viewBox="0 0 16 16"
                      >
                        <path
                          fillRule="evenodd"
                          d="M15.854 5.146a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 0 1 .708-.708L12.5 7.793l2.646-2.647a.5.5 0 0 1 .708 0z"
                        />
                        <path d="M1 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H1zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z" />
                      </svg>
                      Send conenction request
                    </button>
                  </div>
                ),
              ]
            ) : (
              <div className="pt-4 flex items-center justify-between">
                <div className="flex items-center">
                  <button
                    className="btn px-3 mr-3 py-2 text-purple-800 border border-purple-800 hover:text-white hover:bg-purple-800 flex items-center"
                    onClick={() => {
                      router.push("/conversations");
                    }}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      className="h-5 w-5 mr-2"
                      viewBox="0 0 16 16"
                    >
                      <path d="M8 15c4.418 0 8-3.134 8-7s-3.582-7-8-7-8 3.134-8 7c0 1.76.743 3.37 1.97 4.6-.097 1.016-.417 2.13-.771 2.966-.079.186.074.394.273.362 2.256-.37 3.597-.938 4.18-1.234A9.06 9.06 0 0 0 8 15z" />
                    </svg>
                    Message
                  </button>
                  <button
                    className="btn px-3 py-2 text-purple-800 border border-purple-800 hover:text-white hover:bg-purple-800 flex items-center"
                    onClick={onConnectionRemove}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5 mr-2"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                        clipRule="evenodd"
                      />
                    </svg>
                    Remove friend
                  </button>
                </div>
                <button className="btn px-3 py-2 text-purple-800 border border-purple-800 hover:text-white hover:bg-purple-800 flex items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    className="mr-2"
                    viewBox="0 0 16 16"
                  >
                    <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-1 4v3.793l1.146-1.147a.5.5 0 0 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-2-2a.5.5 0 0 1 .708-.708L7.5 11.293V7.5a.5.5 0 0 1 1 0z" />
                  </svg>
                  Resume
                </button>
              </div>
            )}
          </div>
        </div>
        <div className="mt-8">
          <h2 className="text-3xl mb-2">Experience</h2>
          <hr className="mb-4" />
          <UserExperience user_id={id} />
        </div>
        <div className="mt-8">
          <h2 className="text-3xl mb-2">Education</h2>
          <hr className="mb-4" />
          <UserEducation user_id={id} />
        </div>
        <div className="mt-8 space-y-4">
          {feed === undefined
            ? null
            : [feed.map((f) => <Post key={f.post_id} feed={f} />)]}
        </div>
      </main>
    </>
  );
}
