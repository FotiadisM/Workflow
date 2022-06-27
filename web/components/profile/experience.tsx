import { useEffect, useState } from "react";

interface Exp {
  company: string;
  position: string;
  from: string;
  to: string;
}

interface ExperienceProps {
  exp_id: string;
}

const Experience: React.FC<ExperienceProps> = ({ exp_id }) => {
  const [exp, setExp] = useState<Exp | null>(null);

  useEffect(() => {
    setExp({
      company: "Whisper",
      position: "UI Manager",
      from: "06/2018",
      to: "12/2021",
    });
  }, [exp_id]);

  if (exp === null) {
    return null;
  }

  return (
    <div className="flex items-center">
      <div className="text-center px-4 py-2 text-sm text-gray-600">
        <div>{exp.to}</div>
        <div>-</div>
        <div>{exp.from}</div>
      </div>
      <div
        className="flex-1 px-4 py-2 ml-5 text-center"
        style={{ minWidth: "400px" }}
      >
        <span className="mr-2 text-purple-800">{exp.position}</span>
        <span className="text-gray-600 mr-2">{"@"}</span>
        <span className="text-purple-800">{exp.company}</span>
      </div>
    </div>
  );
};

interface UserExperienceProps {
  user_id: string;
}

export const UserExperience: React.FC<UserExperienceProps> = ({ user_id }) => {
  const [exp, setExp] = useState<string[]>([]);

  useEffect(() => {
    setExp([""]);
  }, [user_id]);

  if (exp === undefined) {
    return null;
  }

  return (
    <div className="px-4 border rounded-md space-y-2">
      {exp.map((e, i) => (
        <>
          <Experience key={e} exp_id={e} />
          {i === exp.length - 1 ? null : <hr />}
        </>
      ))}
    </div>
  );
};
