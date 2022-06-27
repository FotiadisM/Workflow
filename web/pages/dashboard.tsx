import { AuthRoute } from "@/components/auth/AuthRoute";
import Navbar from "@/components/navbar/Navbar";
import { serverURI } from "@/src/api/url";
import { Job } from "@/src/types/job";
import { User } from "@/src/types/user";
import { classNames } from "@/src/util";
import { useEffect, useState } from "react";

const onDownloadBtn = (jsobj: any, name: string) => {
  const dataStr =
    "data:text/json;charset=utf-8," +
    encodeURIComponent(JSON.stringify({ jsobj }));
  const a = document.getElementById("downloada");
  if (a !== null) {
    a.setAttribute("href", dataStr);
    a.setAttribute("download", name);
    a.click();
  }
};

const ShowJobs: React.FC = () => {
  const [jobs, setJobs] = useState<Job[]>([]);
  useEffect(() => {
    (async () => {
      try {
        const res = await fetch(serverURI + "/jobs/");
        if (!res.ok) {
          const text = await res.text();
          console.log("failed to fetch jobs:", text);
        }
        const data = await res.json();
        setJobs(data.jobs);
      } catch (err) {
        console.log("failed to fetch jobs:", err);
      }
    })();
  }, []);

  if (jobs === undefined) {
    return null;
  }

  return (
    <span className="flex flex-col items-center justify-center">
      <button
        className="btn px-3 py-2 bg-purple-800 text-white hover:bg-purple-900 mb-4 self-end mr-15"
        onClick={() => {
          onDownloadBtn(jobs, "jobs.json");
        }}
      >
        Download
      </button>
      <table className="text-lg">
        <tr className="text-purple-800">
          <th>id</th>
          <th>UserID</th>
          <th>Title</th>
          <th>Type</th>
          <th>Location</th>
          <th>Company</th>
          <th>Created</th>
        </tr>
        {jobs.map((j) => {
          return (
            <tr key={j.id}>
              <th>{j.id}</th>
              <th>{j.user_id}</th>
              <th>{j.title}</th>
              <th>{j.location}</th>
              <th>{j.company}</th>
              <th>{j.created}</th>
            </tr>
          );
        })}
      </table>
    </span>
  );
};

const ShowUsers: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  useEffect(() => {
    (async () => {
      try {
        const res = await fetch(serverURI + "/users");
        if (!res.ok) {
          const text = await res.text();
          console.log("failed to fetch users:", text);
        }
        const data = await res.json();
        setUsers(data.users);
      } catch (err) {
        console.log("failed to fetch users:", err);
      }
    })();
  }, []);

  if (users === undefined) {
    return null;
  }

  return (
    <span className="flex flex-col items-center justify-center">
      <button
        className="btn px-3 py-2 bg-purple-800 text-white hover:bg-purple-900 mb-4 self-end mr-15"
        onClick={() => {
          onDownloadBtn(users, "users.json");
        }}
      >
        Download
      </button>
      <table className="text-lg">
        <tr className="text-purple-800">
          <th>id</th>
          <th>First Name</th>
          <th>Last Name</th>
          <th>Email</th>
          <th>Company</th>
          <th>Position</th>
        </tr>
        {users.map((u) => {
          return (
            <tr key={u.id}>
              <th>{u.id}</th>
              <th>{u.f_name}</th>
              <th>{u.l_name}</th>
              <th>{u.email}</th>
              <th>{u.company}</th>
              <th>{u.position}</th>
            </tr>
          );
        })}
      </table>
    </span>
  );
};

export default function Dashboard() {
  const [activePanel, setActivePanel] = useState<"users" | "jobs" | "posts">(
    "users"
  );

  const returnTable = () => {
    switch (activePanel) {
      case "users":
        return <ShowUsers />;
      case "jobs":
        return <ShowJobs />;
      default:
        return <ShowUsers />;
    }
  };

  return (
    <>
      <AuthRoute>
        <Navbar />
        <main className="flex flex-col justify-center">
          <div className="my-10 flex flex-col items-center">
            <hr
              className="border-2 border-purple-800"
              style={{ width: "80vw" }}
            />
            <nav className="flex items-center space-x-20 my-2 text-xl">
              <button
                className={classNames(
                  activePanel === "users"
                    ? "bg-purple-800 text-white hover:bg-purple-700"
                    : "bg-purple-50 text-purple-900 hover:bg-purple-800 hover:text-white",
                  "btn p-2 rounded-md border border-purple-800"
                )}
                onClick={() => setActivePanel("users")}
              >
                Users
              </button>
              <button
                className={classNames(
                  activePanel === "jobs"
                    ? "bg-purple-800 text-white hover:bg-purple-700"
                    : "bg-purple-50 text-purple-900 hover:bg-purple-800 hover:text-white",
                  "btn p-2 rounded-md border border-purple-800"
                )}
                onClick={() => setActivePanel("jobs")}
              >
                Jobs
              </button>
              <button
                className={classNames(
                  activePanel === "posts"
                    ? "bg-purple-800 text-white hover:bg-purple-700"
                    : "bg-purple-50 text-purple-900 hover:bg-purple-800 hover:text-white",
                  "btn p-2 rounded-md border border-purple-800"
                )}
                onClick={() => setActivePanel("posts")}
              >
                Posts
              </button>
            </nav>
            <hr
              className="border-2 border-purple-800"
              style={{ width: "80vw" }}
            />
          </div>
          <div className="my-10 flex flex-col justify-center items-center">
            <a id="downloada" style={{ display: "none" }} />
            {returnTable()}
          </div>
        </main>
      </AuthRoute>
    </>
  );
}
