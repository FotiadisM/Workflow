import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";
import { User } from "@/src/types/user";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface PostTypeProps {
  type: string;
  perpetaror_id: string;
}

const PostType: React.FC<PostTypeProps> = ({ type, perpetaror_id }) => {
  const [postPerpetrator, setPostPerpetrator] = useState<User | null>(null);

  useEffect(() => {
    fetchPerpetrator(perpetaror_id)
      .then((u) => {
        if (u !== null) {
          setPostPerpetrator(u);
        }
      })
      .catch((err) => console.log(err));
  }, [perpetaror_id]);

  if (postPerpetrator === null) {
    return null;
  }

  switch (type) {
    case "like":
      return (
        <div className="px-3">
          <span className="text-purple-800 hover:underline cursor-pointer">
            {postPerpetrator.f_name} {postPerpetrator.l_name}
          </span>{" "}
          <span className="font-semibold text-purple-700">liked</span> a post.
          <hr className="my-3" />
        </div>
      );
    case "share":
      return (
        <div className="px-3">
          <span className="text-purple-800 hover:underline cursor-pointer">
            {postPerpetrator.f_name} {postPerpetrator.l_name}
          </span>{" "}
          <span className="font-semibold text-purple-700">shared</span> a post.
          <hr className="my-3" />
        </div>
      );
    case "comment":
      return (
        <div className="px-3">
          <span className="text-purple-800 hover:underline cursor-pointer">
            {postPerpetrator.f_name} {postPerpetrator.l_name}
          </span>{" "}
          <span className="font-semibold text-purple-700">comment</span> on a
          post.
          <hr className="my-3" />
        </div>
      );
    default:
      return null;
  }

  //   if (type === "like") {
  //     return (
  //       <div className="px-3">
  //         <span className="text-purple-800 hover:underline cursor-pointer">
  //           {postPerpetrator.f_name} {postPerpetrator.l_name}
  //         </span>{" "}
  //         <span className="font-semibold text-purple-700">liked</span> a post.
  //         <hr className="my-3" />
  //       </div>
  //     );
  //   }

  //   if (type === "share") {
  //     return (
  //       <div className="px-3">
  //         <span className="text-purple-800 hover:underline cursor-pointer">
  //           {postPerpetrator.f_name} {postPerpetrator.l_name}
  //         </span>{" "}
  //         <span className="font-semibold text-purple-700">shared</span> a post.
  //         <hr className="my-3" />
  //       </div>
  //     );
  //   }

  //   return (
  //     <div className="px-3">
  //       <span className="text-purple-800 hover:underline cursor-pointer">
  //         {postPerpetrator.f_name} {postPerpetrator.l_name}
  //       </span>{" "}
  //       <span className="font-semibold text-purple-700">comment</span> on a post.
  //       <hr className="my-3" />
  //     </div>
  //   );
};

interface PostHeadProps {
  user_id: string;
  type: string;
  perpetaror_id: string;
  created: string;
}

export const PostHead: React.FC<PostHeadProps> = ({
  user_id,
  type,
  perpetaror_id,
  created,
}) => {
  const router = useRouter();
  const [postUser, setPostUser] = useState<User | null>(null);

  useEffect(() => {
    fetchPerpetrator(user_id)
      .then((u) => {
        if (u !== null) {
          setPostUser(u);
        }
      })
      .catch((err) => console.log(err));
  }, [user_id]);

  if (postUser === null) {
    return null;
  }

  console.log("TYPE", type);

  return (
    <div className="pt-4 px-3">
      <PostType type={type} perpetaror_id={perpetaror_id} />
      <div className="flex items-center justify-between mx-3">
        <div className="flex items-center">
          <img
            className="h-12 w-12 rounded-full"
            src={serverURI + "/static/" + postUser.profile_pic}
            alt="profile-picture"
          />
          <button
            className="ml-3 font-semibold text-purple-700 cursor-pointer hover:underline focus:outline-none"
            onClick={() => router.push("/user/" + postUser.id)}
          >
            {postUser.f_name} {postUser.l_name}
          </button>
        </div>
        <div className="text-gray-600">{created}</div>
      </div>
    </div>
  );
};
