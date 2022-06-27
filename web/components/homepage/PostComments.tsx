import { fetchComment, postComment } from "@/src/api/posts";
import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";
import { Post, Comment } from "@/src/types/posts";
import { User } from "@/src/types/user";
import { useLayoutEffect } from "@/src/useIsomorphicLayoutEffect";
import { classNames } from "@/src/util";
import React, {
  Dispatch,
  SetStateAction,
  useEffect,
  useRef,
  useState,
} from "react";
import { useAuth } from "../auth/AuthRoute";

interface CommentUserInputProps {
  onComment: (text: string) => void;
}

const CommentUserInput: React.FC<CommentUserInputProps> = ({ onComment }) => {
  const auth = useAuth();
  const inputRef = useRef<HTMLTextAreaElement | null>(null);
  const [commentText, setCommentText] = useState<string>("");

  useLayoutEffect(() => {
    if (inputRef !== null) {
      if (inputRef.current !== null) {
        inputRef.current.focus();
      }
    }
  }, []);

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onComment(commentText);
    setCommentText("");
  };

  return (
    <div className="flex item-center">
      <img
        className="h-8 w-8 rounded-full"
        src={serverURI + "/static/" + auth?.user?.profile_pic}
        alt="profile-picture"
      />
      <form className="flex-1 flex items-center ml-3" onSubmit={onSubmit}>
        <textarea
          ref={inputRef}
          value={commentText}
          onChange={(e) => setCommentText(e.target.value)}
          className="flex-1 px-3 py-1 mr-1 rounded-lg bg-purple-100 focus:outline-none"
          placeholder="Write a comment.."
        ></textarea>
        <button
          type="submit"
          className="focus:outline-none hover:bg-gray-200 rounded-md"
        >
          <span className="sr-only">Post the comment</span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            className="h-10 w-10 text-purple-800"
            viewBox="0 0 16 16"
          >
            <path d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z" />
          </svg>
        </button>
      </form>
    </div>
  );
};

interface PostCommentProps {
  comment_id: string;
}

const PostComment: React.FC<PostCommentProps> = ({ comment_id }) => {
  const auth = useAuth();
  const [liked, setLiked] = useState<boolean>(false);
  const [comment, setComment] = useState<Comment | null>(null);
  const [commentator, setCommentator] = useState<User | null>(null);

  useEffect(() => {
    fetchComment(comment_id)
      .then((c) => {
        if (c !== null) {
          if (auth !== null)
            if (auth.user !== null)
              if (c.likes.indexOf(auth.user.id) !== -1) setLiked(true);
          fetchPerpetrator(c.user_id)
            .then((u) => {
              setCommentator(u);
              setComment(c);
            })
            .catch((err) => console.log(err));
        }
      })
      .catch((err) => console.log(err));
  }, [comment_id]);

  const onCommentLike = () => {
    // TODO: toogleCommentLike
    setLiked((l) => !l);
  };

  if (comment == null || commentator == null) {
    return null;
  }

  return (
    <div className="flex">
      <img
        className="mt-1 h-8 w-8 rounded-full"
        src={serverURI + "/static/" + commentator.profile_pic}
        alt="profile-picture"
      />
      <div className="ml-3 flex items-center" style={{ maxWidth: "100%" }}>
        <div>
          <div className="px-3 py-2 rounded-lg bg-gray-100">
            <div className="font-semibold text-purple-500 hover:underline cursor-pointer">
              {commentator.f_name} {commentator.l_name}
            </div>
            <div className="text-gray-700">{comment.text}</div>
          </div>
          <div className="flex items-center justify-between text-sm px-3">
            <div className="flex items-center">
              <div
                className={classNames(
                  liked
                    ? "text-purple-800 font-medium"
                    : "text-gray-600 hover:text-purple-800",
                  "cursor-pointer"
                )}
                onClick={onCommentLike}
              >
                Like
              </div>
              <div className="ml-2 text-gray-400">{comment.created}</div>
            </div>
            <div className="float-right flex items-center">
              <svg
                viewBox="0 0 16 16"
                xmlns="http://www.w3.org/2000/svg"
                fill="currentColor"
                className="h-3 w-3 text-purple-800"
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
              <div className="ml-1 text-gray-600">{comment.likes.length}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

interface PostCommentsProps {
  post_id: string;
  comments: string[];
  setPost: Dispatch<SetStateAction<Post | null>>;
}

export const PostComments: React.FC<PostCommentsProps> = ({
  post_id,
  comments,
  setPost,
}) => {
  const auth = useAuth();

  const onComment = (text: string) => {
    if (auth !== null)
      if (auth.user !== null)
        postComment(post_id, auth.user.id, text)
          .then((c) => {
            if (c !== null)
              setPost((p) => {
                if (p !== null)
                  return { ...p, comments: [c.id, ...p.comments] };
                return p;
              });
          })
          .catch((err) => console.log(err));
  };

  return (
    <div>
      <CommentUserInput onComment={onComment} />
      <div className="mt-3 space-y-3">
        {comments.map((c) => (
          <PostComment key={c} comment_id={c} />
        ))}
      </div>
    </div>
  );
};
