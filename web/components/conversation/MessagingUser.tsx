import { serverURI } from "@/src/api/url";
import { fetchPerpetrator } from "@/src/api/user";
import { User } from "@/src/types/user";
import { classNames } from "@/src/util";
import { Dispatch, SetStateAction, useEffect, useState } from "react";

interface MessagingUserProps {
  user_id: string;
  current: boolean;
  setCurUser: () => void;
  setCurrPerp: Dispatch<SetStateAction<User | null>>;
}

export default function MessagingUser({
  user_id,
  current,
  setCurUser,
  setCurrPerp,
}: MessagingUserProps) {
  const [perp, setPerp] = useState<User | null>(null);

  useEffect(() => {
    fetchPerpetrator(user_id)
      .then((u) => {
        setPerp(u);
      })
      .catch((err) => console.log(err));
  }, [user_id]);

  useEffect(() => {
    if (current) if (perp !== null) setCurrPerp(perp);
  }, [current, perp]);

  if (perp === null) {
    return null;
  }

  return (
    <div
      className={classNames(
        current ? "bg-gray-200 text-purple-800" : "hover:bg-gray-100",
        "flex items-center border rounded-md py-2 px-3 cursor-pointer"
      )}
      onClick={setCurUser}
    >
      <img
        className="rounded-full h-10 w-10 mr-4"
        src={serverURI + "/static/" + perp.profile_pic}
      />
      <div>
        {perp.f_name} {perp.l_name}
      </div>
    </div>
  );
}
