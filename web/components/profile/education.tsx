import { useEffect, useState } from "react";

interface Ed {
  time: string;
  title: string;
}

interface EducationProps {
  ed_id: string;
}

const Education: React.FC<EducationProps> = ({ ed_id }) => {
  const [ed, setEd] = useState<Ed | null>(null);

  useEffect(() => {
    setEd({
      time: "2017",
      title:
        "Information and Telecommunication @ Natioanl and Kapodistrian University of Athens",
    });
  }, [ed_id]);

  if (ed === null) {
    return null;
  }

  return (
    <div className="flex items-center">
      <div className="px-4 py-2 mr-4 text-gray-600">{ed.time}</div>
      <div
        className="flex-1 px-4 py-2 text-purple-800"
        style={{ minWidth: "400px" }}
      >
        {ed.title}
      </div>
    </div>
  );
};

interface UserEducationProps {
  user_id: string;
}

export const UserEducation: React.FC<UserEducationProps> = ({ user_id }) => {
  const [ed, setEd] = useState<string[]>([]);

  useEffect(() => {
    setEd([""]);
  }),
    [user_id];

  if (ed === undefined) {
    return null;
  }

  return (
    <div className="px-4 border rounded-md space-y-2">
      {ed.map((e, i) => (
        <>
          <Education ed_id={e} />
          {i === ed.length - 1 ? null : <hr />}
        </>
      ))}
    </div>
  );
};
