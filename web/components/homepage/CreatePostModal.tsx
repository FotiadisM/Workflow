import { serverURI } from "@/src/api/url";
import { classNames } from "@/src/util";
import { Dialog, Listbox, Transition } from "@headlessui/react";
import React, { Dispatch, Fragment, SetStateAction, useState } from "react";
import { useAuth } from "../auth/AuthRoute";

const publishOptions: string[] = ["Only you", "Only friends", "All"];

interface ListProps {
  selected: string;
  setSelected: React.Dispatch<React.SetStateAction<string>>;
}

const List: React.FC<ListProps> = ({ selected, setSelected }) => {
  return (
    <Listbox value={selected} onChange={setSelected}>
      <div className="relative" style={{ minWidth: "150px" }}>
        <Listbox.Button className="btn px-3 py-2 w-full text-purple-800 border-2 border-purple-800">
          <div className="flex items-center justify-between">
            <span className="mr-3">{selected}</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              viewBox="0 0 16 16"
            >
              <path
                fillRule="evenodd"
                d="M3.646 13.854a.5.5 0 0 0 .708 0L8 10.207l3.646 3.647a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 0 0 0 .708zm0-11.708a.5.5 0 0 1 .708 0L8 5.793l3.646-3.647a.5.5 0 0 1 .708.708l-4 4a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 0-.708z"
              />
            </svg>
          </div>
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options className="absolute w-full overflow-auto text-base py-1 mt-1 bg-white rounded-lg shadow-md ring-1 ring-black ring-opacity-5 focus:outline-none">
            {publishOptions.map((o) => (
              <Listbox.Option key={o} value={o} as={Fragment}>
                {({ active }) => (
                  <li
                    key={o}
                    className={classNames(
                      active ? "bg-purple-100" : "",
                      "cursor-default truncate px-2 py-1 flex"
                    )}
                  >
                    <span>{o}</span>
                  </li>
                )}
              </Listbox.Option>
            ))}
          </Listbox.Options>
        </Transition>
      </div>
    </Listbox>
  );
};

interface FilesPreviewerProps {
  files: File[];
  setFiles: React.Dispatch<React.SetStateAction<File[]>>;
}

const FilesPreviewer: React.FC<FilesPreviewerProps> = ({ files, setFiles }) => {
  const onFileRemove = (i: number) => {
    setFiles((oldFiles) => {
      let tmpFiles = [...oldFiles];
      tmpFiles.splice(i, 1);
      return tmpFiles;
    });
  };

  if (files.length === 0) {
    return null;
  }

  return (
    <div className="space-y-6">
      {files.map((f, i) => {
        if (f.type.match("image.*")) {
          return (
            <div key={i} className="relative rounded-lg">
              <img
                key={i}
                src={URL.createObjectURL(f)}
                className="rounded-lg"
              />
              <button
                className="absolute -top-2 -right-2 rounded-full p-2 bg-purple-800 focus:outline-none shadow-xl"
                onClick={() => onFileRemove(i)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="white"
                  viewBox="0 0 16 16"
                >
                  <path d="M1.293 1.293a1 1 0 0 1 1.414 0L8 6.586l5.293-5.293a1 1 0 1 1 1.414 1.414L9.414 8l5.293 5.293a1 1 0 0 1-1.414 1.414L8 9.414l-5.293 5.293a1 1 0 0 1-1.414-1.414L6.586 8 1.293 2.707a1 1 0 0 1 0-1.414z" />
                </svg>
              </button>
            </div>
          );
        } else if (f.type.match("video.*")) {
          return (
            <div key={i} className="relative rounded-lg">
              <video
                width="100%"
                src={URL.createObjectURL(f)}
                className="rounded-lg"
                controls
              />
              <button
                className="absolute -top-2 -right-2 rounded-full p-2 bg-purple-800 focus:outline-none shadow-xl"
                onClick={() => onFileRemove(i)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="white"
                  viewBox="0 0 16 16"
                >
                  <path d="M1.293 1.293a1 1 0 0 1 1.414 0L8 6.586l5.293-5.293a1 1 0 1 1 1.414 1.414L9.414 8l5.293 5.293a1 1 0 0 1-1.414 1.414L8 9.414l-5.293 5.293a1 1 0 0 1-1.414-1.414L6.586 8 1.293 2.707a1 1 0 0 1 0-1.414z" />
                </svg>
              </button>
            </div>
          );
        }

        return null;
      })}
    </div>
  );
};

interface CreatePostModalProps {
  isModalOpen: boolean;
  setIsModalOpen: Dispatch<SetStateAction<boolean>>;
  onPostSubmit: (post_id: string) => void;
}

export const CreatePostModal: React.FC<CreatePostModalProps> = ({
  isModalOpen,
  setIsModalOpen,
  onPostSubmit,
}) => {
  const auth = useAuth();
  const [selected, setSelected] = useState<string>(publishOptions[2]);

  const [text, setText] = useState<string>("");

  const onFilesButtonClick = () => {
    const inpt = document.getElementById("filesInput");
    if (inpt !== null) {
      inpt.click();
    }
  };

  const [files, setFiles] = useState<File[]>([]);
  const onFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFiles((f) => {
      let tmpFiles = [...f];
      if (e.target.files !== null) {
        for (let i = 0; i < e.target.files.length; i++) {
          let tmpFile = e.target.files.item(i);
          if (tmpFile !== null) {
            tmpFiles.push(tmpFile);
          }
        }
        e.target.value = "";
      }

      return tmpFiles;
    });
  };

  const onPostCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData();

    if (auth !== null) {
      if (auth?.user !== null) {
        formData.append("user_id", auth.user.id);
        formData.append("text", text);
        switch (selected) {
          case "Only you":
            formData.append("visibility", "you");
            break;
          case "Only friends":
            formData.append("visibility", "friends");
            break;
          case "All":
            formData.append("visibility", "all");
            break;
        }
        files.forEach((f) => {
          if (f.type.match("video.*")) {
            formData.append("videos", f);
          }
          if (f.type.match("image.")) {
            formData.append("images", f);
          }
        });
      }
    }

    (async function () {
      const res = await fetch(serverURI + "/posts/", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        const text = await res.text();
        console.log("error creating post:" + text);
      }

      const data = await res.json();
      onPostSubmit(data.post.id);
      console.log(data);
    })();
  };

  return (
    <Transition appear show={isModalOpen} as={Fragment}>
      <Dialog
        as="div"
        className="fixed inset-0 z-10 overflow-y-auto"
        onClose={() => setIsModalOpen(false)}
      >
        <div className="min-h-screen px-4 text-center">
          <Dialog.Overlay className="fixed inset-0 bg-black opacity-40" />

          {/* This element is to trick the browser into centering the modal contents. */}
          <span
            className="inline-block h-screen align-middle"
            aria-hidden="true"
          >
            &#8203;
          </span>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <div className="inline-block w-full max-w-3xl p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl">
              <Dialog.Title
                as="h3"
                className="text-2xl font-medium  text-gray-900"
              >
                What's on your mind, Mike?
              </Dialog.Title>
              <form className="mt-3" onSubmit={(e) => onPostCreate(e)}>
                <div className="mt-4 flex items-center justify-between">
                  <button
                    className="btn py-2 px-3 flex items-center hover:bg-purple-100 active:bg-purple-50"
                    type="button"
                    onClick={onFilesButtonClick}
                  >
                    <input
                      id="filesInput"
                      type="file"
                      className="hidden"
                      accept="image/*,video/*"
                      multiple
                      onChange={(e) => {
                        onFileChange(e);
                      }}
                    />
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      className="h-5 w-5 text-purple-800"
                      viewBox="0 0 16 16"
                    >
                      <path d="M4.502 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3z" />
                      <path d="M14.002 13a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V5A2 2 0 0 1 2 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v8a2 2 0 0 1-1.998 2zM14 2H4a1 1 0 0 0-1 1h9.002a2 2 0 0 1 2 2v7A1 1 0 0 0 15 11V3a1 1 0 0 0-1-1zM2.002 4a1 1 0 0 0-1 1v8l2.646-2.354a.5.5 0 0 1 .63-.062l2.66 1.773 3.71-3.71a.5.5 0 0 1 .577-.094l1.777 1.947V5a1 1 0 0 0-1-1h-10z" />
                    </svg>
                    <span className="text-gray-600 ml-3 font-medium">
                      Photo/Video
                    </span>
                  </button>
                  <div className="flex justify-end items-center">
                    <p className="mr-5 text-gray-600">
                      Choose who can see your post:
                    </p>
                    <List {...{ selected, setSelected }} />
                    <button
                      type="submit"
                      className="ml-3 btn py-2 px-4 text-lg text-white bg-purple-800 rounded-md hover:bg-purple-900 shadow-md"
                    >
                      Publish
                    </button>
                  </div>
                </div>
                <div className="flex mt-4">
                  <img
                    className="h-12 w-12 rounded-full"
                    src={serverURI + "/static/" + auth?.user?.profile_pic}
                    alt="profile-picture"
                  />
                  <div className="flex-1 ml-4">
                    <textarea
                      className="w-full px-4 py-1 rounded-xl bg-gray-100 focus:outline-none"
                      placeholder="Share your thoughts"
                      value={text}
                      onChange={(e) => setText(e.target.value)}
                      rows={4}
                    />
                  </div>
                </div>
                <div className="mt-8 px-8">
                  <FilesPreviewer {...{ files, setFiles }} />
                </div>
              </form>
            </div>
          </Transition.Child>
        </div>
      </Dialog>
    </Transition>
  );
};
