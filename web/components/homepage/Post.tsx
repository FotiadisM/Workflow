import { fetchPost } from "@/src/api/posts";
import { serverURI } from "@/src/api/url";
import { Feed, Post as hPost } from "@/src/types/posts";
import { classNames } from "@/src/util";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { useAuth } from "../auth/AuthRoute";
import { PostComments } from "./PostComments";
import { PostHead } from "./PostHead";

interface PostBodyProps {
  text: string;
  images: string[];
  videos: string[];
}

const PostBody: React.FC<PostBodyProps> = ({ text, images, videos }) => {
  const [curMedia, setCurMedia] = useState<{
    type: "image" | "video";
    index: number;
  }>(() => {
    if (images.length === 0) return { type: "video", index: 0 };
    return { type: "image", index: 0 };
  });

  const onBack = () => {
    if (curMedia.type === "image") {
      if (curMedia.index === 0) return;
      setCurMedia((m) => ({ type: "image", index: m.index - 1 }));
      return;
    }

    if (curMedia.index === 0) {
      if (images.length === 0) return;
      setCurMedia({ type: "image", index: images.length - 1 });
      return;
    }
    setCurMedia((m) => ({ type: "video", index: m.index - 1 }));
  };

  const onForward = () => {
    if (curMedia.type === "image") {
      if (curMedia.index === images.length - 1) {
        if (videos.length === 0) return;
        setCurMedia({ type: "video", index: 0 });
        return;
      }
      setCurMedia((m) => ({ type: "image", index: m.index + 1 }));
      return;
    }

    if (curMedia.index === videos.length - 1) return;
    setCurMedia((m) => ({ type: "video", index: m.index + 1 }));
  };

  return (
    <div className="px-3 pb-1">
      <div className="mx-2 mt-2 text-gray-800">{text}</div>
      {images.length === 0 && videos.length === 0 ? null : (
        <div className="relative mt-4">
          {curMedia.type === "image" ? (
            <img
              className="rounded-md"
              src={serverURI + "/static/" + images[curMedia.index]}
              alt="post picture"
            />
          ) : (
            <video
              width="100%"
              className="rounded-md"
              src={serverURI + "/static/" + videos[curMedia.index]}
              controls
            />
          )}
          <button
            className="absolute top-1/2 left-4 transform -translate-y-1/2 text-gray-600"
            onClick={onBack}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="h-8 w-8"
              viewBox="0 0 16 16"
            >
              <path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3.5 7.5a.5.5 0 0 1 0 1H5.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L5.707 7.5H11.5z" />
            </svg>
          </button>
          <button
            className="absolute top-1/2 right-4 transform -translate-y-1/2 text-gray-600"
            onClick={onForward}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="h-8 w-8"
              viewBox="0 0 16 16"
            >
              <path d="M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0zM4.5 7.5a.5.5 0 0 0 0 1h5.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 0-.708l-3-3a.5.5 0 1 0-.708.708L10.293 7.5H4.5z" />
            </svg>
          </button>
        </div>
      )}
    </div>
  );
};

interface PostStatsProps {
  liked: boolean;
  likesNum: number;
  commentsNum: number;
  setOpenComments: Dispatch<SetStateAction<boolean>>;
}

const PostStats: React.FC<PostStatsProps> = ({
  liked,
  likesNum,
  commentsNum,
  setOpenComments,
}) => {
  return (
    <div className="flex items-center justify-between text-gray-600">
      <div className="flex items-center">
        <svg
          viewBox="0 0 16 16"
          fill="currentColor"
          className={classNames(liked ? "text-purple-800" : "", "h-5 w-5")}
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="m0 1v8c0 .552246.447693 1 1 1h3v-10h-3c-.552307 0-1 .447693-1 1z"
            transform="translate(0 5)"
          />
          <path
            d="m9.15332 5.02979h-2.9541c-.258301 0-.387695-.172363-.431152-.246582-.043457-.0737305-.131348-.270508-.0063477-.496094l1.0415-1.87549c.228516-.410645.251953-.893555.0649414-1.32471-.187012-.43164-.556152-.744629-1.0127-.858398l-.734375-.183594c-.178711-.0449219-.368164.0122071-.492676.150391l-3.9873 4.42969c-.413574.460449-.641113 1.0542-.641113 1.67236v5.23242c0 1.37842 1.12158 2.5 2.5 2.5l4.97412-.0004883c1.12305 0 2.11475-.756348 2.41113-1.83887l1.06738-4.89844c.03125-.13623.0473633-.275879.0473633-.415527 0-1.01807-.828613-1.84668-1.84668-1.84668z"
            transform="translate(5 .97)"
          />
        </svg>
        <div className={classNames(liked ? "text-purple-800" : "", "ml-2")}>
          {liked ? "You and " : ""}
          {likesNum}
        </div>
      </div>
      <button
        className="cursor-pointer hover:text-purple-800 focus:outline-none"
        onClick={() => setOpenComments((o) => !o)}
      >
        {commentsNum} Comments
      </button>
    </div>
  );
};

interface PostActionsProps {
  liked: boolean;
  onToggleLike: () => void;
  setOpenComments: Dispatch<SetStateAction<boolean>>;
}

const PostActions: React.FC<PostActionsProps> = ({
  liked,
  onToggleLike,
  setOpenComments,
}) => {
  return (
    <div className="flex items-center justify-between">
      <button
        className={classNames(
          liked ? "font-semibold text-purple-700 bg-gray-100" : "",
          "flex-1 btn py-2 text-purple-800 hover:bg-gray-200"
        )}
        onClick={onToggleLike}
      >
        Like
      </button>
      <button
        className="flex-1 btn py-2 text-purple-800 hover:bg-gray-200"
        onClick={() => setOpenComments(true)}
      >
        Comment
      </button>
      <button className="flex-1 btn py-2 text-purple-800 hover:bg-gray-200">
        Share
      </button>
    </div>
  );
};

interface PostProps {
  feed: Feed;
}

export const Post: React.FC<PostProps> = ({ feed }) => {
  const auth = useAuth();
  const [post, setPost] = useState<hPost | null>(null);
  const [liked, setLiked] = useState<boolean>(false);

  useEffect(() => {
    if (feed.post_id !== undefined) {
      fetchPost(feed.post_id).then((p) => {
        if (p !== null) {
          if (auth !== null)
            if (auth.user !== null)
              if (p?.likes.indexOf(auth.user.id) !== -1) setLiked(true);
          setPost(p);
        }
      });
    }
  }, [feed.post_id]);

  const onToggleLike = () => {
    (async () => {
      if (auth !== null)
        if (auth.user !== null) {
          try {
            const res = await fetch(serverURI + "/posts/like", {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({
                user_id: auth.user.id,
                post_id: feed.post_id,
              }),
            });

            if (res.status === 200) {
              setLiked((l) => !l);
            }
          } catch (err) {
            console.log(err);
          }
        }
    })();
  };

  const [openComments, setOpenComments] = useState<boolean>(false);

  if (post === null) {
    return null;
  }

  return (
    <div className="border rounded-lg shadow-lg" style={{ width: "600px" }}>
      <PostHead
        user_id={post.user_id}
        type={feed.type}
        perpetaror_id={feed.perpetator_id}
        created={post.created}
      />
      <PostBody text={post.text} images={post.images} videos={post.videos} />
      <div className="py-3 px-5">
        <PostStats
          {...{ liked, setOpenComments }}
          likesNum={post.likes.length}
          commentsNum={post.comments.length}
        />
        <hr className="my-3" />
        <PostActions
          liked={liked}
          onToggleLike={onToggleLike}
          setOpenComments={setOpenComments}
        />
        {openComments ? (
          <>
            <hr className="my-3" />
            <PostComments
              post_id={post.id}
              comments={post.comments}
              setPost={setPost}
            />
          </>
        ) : null}
      </div>
    </div>
  );
};
