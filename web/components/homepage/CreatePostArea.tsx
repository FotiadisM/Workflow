import { serverURI } from "@/src/api/url";
import { Feed } from "@/src/types/posts";
import { Dispatch, SetStateAction, useState } from "react";
import { useAuth } from "../auth/AuthRoute";
import { CreatePostModal } from "./CreatePostModal";

interface CreateButtonsProps {
  text: string;
  setIsModalOpen: Dispatch<SetStateAction<boolean>>;
}

const CreateButtons: React.FC<CreateButtonsProps> = ({
  text,
  setIsModalOpen,
  children,
}) => {
  return (
    <button
      className="flex-1 flex items-center justify-center py-3 rounded-lg hover:bg-gray-100 text-center cursor-pointer focus:outline-none"
      onClick={() => setIsModalOpen(true)}
    >
      {children}
      <div className="ml-2 text-gray-800">{text}</div>
    </button>
  );
};

interface CreatePostAreaProps {
  setFeed: Dispatch<SetStateAction<Feed[]>>;
}

export const CreatePostArea: React.FC<CreatePostAreaProps> = ({ setFeed }) => {
  const auth = useAuth();
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const onPostSubmit = (post_id: string) => {
    setFeed((oldFeed) => {
      if (oldFeed === undefined) {
        return [
          {
            id: "-2",
            perpetator_id: auth!.user!.id,
            post_id: post_id,
            type: "post",
          },
        ];
      }

      if (oldFeed.length === 0) {
        return [
          {
            id: "-2",
            perpetator_id: auth!.user!.id,
            post_id: post_id,
            type: "post",
          },
        ];
      }

      return [
        {
          id: "-2",
          perpetator_id: auth!.user!.id,
          post_id: post_id,
          type: "post",
        },
        ...oldFeed,
      ];
    });
    setIsModalOpen(false);
  };

  return (
    <>
      <div
        className="py-2 px-3 rounded-lg shadow-md bg-gray-50"
        style={{ minWidth: "556px" }}
      >
        <div className="flex items-start">
          <img
            className="h-10 w-10 rounded-full"
            src={serverURI + "/static/" + auth?.user?.profile_pic}
            alt="profile-picture"
          />
          <form className="flex-1 ml-4" onClick={() => setIsModalOpen(true)}>
            <textarea
              className="w-full px-4 py-1 rounded-xl bg-gray-100 focus:outline-none"
              placeholder={`What's on your mind, ${auth?.user?.f_name}?`}
            />
          </form>
        </div>
        <hr className="my-3 mx-3" />
        <div className="flex items-center">
          <CreateButtons text="Photo/Video" setIsModalOpen={setIsModalOpen}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="h-7 w-7 text-purple-800"
              viewBox="0 0 16 16"
            >
              <path d="M4.502 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3z" />
              <path d="M14.002 13a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V5A2 2 0 0 1 2 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v8a2 2 0 0 1-1.998 2zM14 2H4a1 1 0 0 0-1 1h9.002a2 2 0 0 1 2 2v7A1 1 0 0 0 15 11V3a1 1 0 0 0-1-1zM2.002 4a1 1 0 0 0-1 1v8l2.646-2.354a.5.5 0 0 1 .63-.062l2.66 1.773 3.71-3.71a.5.5 0 0 1 .577-.094l1.777 1.947V5a1 1 0 0 0-1-1h-10z" />
            </svg>
          </CreateButtons>
          <CreateButtons text="Create an event" setIsModalOpen={setIsModalOpen}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="h-7 w-7 text-purple-800"
              viewBox="0 0 16 16"
            >
              <path d="M3.5 0a.5.5 0 0 1 .5.5V1h8V.5a.5.5 0 0 1 1 0V1h1a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1V.5a.5.5 0 0 1 .5-.5zM2 2a1 1 0 0 0-1 1v1h14V3a1 1 0 0 0-1-1H2zm13 3H1v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V5z" />
              <path d="M11 7.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1z" />
            </svg>
          </CreateButtons>
        </div>
      </div>
      <CreatePostModal {...{ isModalOpen, setIsModalOpen, onPostSubmit }} />
    </>
  );
};
