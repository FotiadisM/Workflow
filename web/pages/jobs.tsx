import React, { useEffect, useState } from "react";
import JobListItem from "@/components/jobs/JobListItem";
import Navbar from "@/components/navbar/Navbar";
import JobCreationForm from "@/components/jobs/JobCreationForm";
import { classNames } from "@/src/util";
import { Job, SideBarType } from "@/src/types/job";
import { fetchJobs, postJob } from "@/src/api/jobs";
import { AuthRoute, useAuth } from "@/components/auth/AuthRoute";
import { serverURI } from "@/src/api/url";

const sideBar: { name: string; type: SideBarType }[] = [
  { name: "Career Opportunities", type: SideBarType.SEARCH },
  { name: "Interested in", type: SideBarType.INTERESTED },
  { name: "Applied", type: SideBarType.APPLIED },
  { name: "My Job Postings", type: SideBarType.CREATE },
];

export default function Jobs() {
  const auth = useAuth();

  const [curPage, setCurPage] = useState<number>(0);
  const changePage = (i: number) => {
    if (sideBar[i].type == SideBarType.SEARCH) {
      setCurPage(i);
      setCurrJobs(jobs.all);
    } else if (sideBar[i].type == SideBarType.INTERESTED) {
      setCurPage(i);
      setCurrJobs(jobs.interested);
    } else if (sideBar[i].type == SideBarType.APPLIED) {
      setCurPage(i);
      setCurrJobs(jobs.applied);
    } else {
      setCurPage(i);
      setCurrJobs(jobs.user);
    }
  };

  const [currJobs, setCurrJobs] = useState<Job[]>([]);
  const [jobs, setJobs] = useState<{
    all: Job[];
    interested: Job[];
    applied: Job[];
    user: Job[];
  }>({
    all: [],
    interested: [],
    applied: [],
    user: [],
  });

  useEffect(() => {
    fetchJobs()
      .then((jobs) => {
        if (jobs !== null) {
          let int: Job[] = [];
          let appl: Job[] = [];
          let usr: Job[] = [];

          for (let i = 0; i < jobs.length; i++) {
            if (auth !== null && auth.user !== null) {
              if (jobs[i].user_id === auth.user.id) usr.push(jobs[i]);
              if (jobs[i].interested.indexOf(auth.user.id) !== -1)
                int.push(jobs[i]);
              if (jobs[i].applied.indexOf(auth.user.id) !== -1)
                appl.push(jobs[i]);
            }
          }

          setJobs({ all: jobs, interested: int, applied: appl, user: usr });
          setCurrJobs(jobs);
        }
      })
      .catch((err) => console.error(err));
  }, []);

  const onJobButtonPress = (job: Job, action: SideBarType) => {
    if (action === SideBarType.CREATE) {
      setJobFormState({ open: true, mode: "edit", job: job });
      return;
    }

    // remove from interested
    if (action === SideBarType.SEARCH) {
      fetch(serverURI + "/jobs/interested", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          user_id: auth?.user?.id,
          job_id: job.id,
        }),
      })
        .then(() => {
          setJobs((old) => {
            old.interested.splice(old.interested.indexOf(job), 1);
            return { ...old };
          });
        })
        .catch((err) => console.log(err));
      return;
    }

    // add to interested
    if (action === SideBarType.INTERESTED) {
      fetch(serverURI + "/jobs/interested", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          user_id: auth?.user?.id,
          job_id: job.id,
        }),
      })
        .then(() => {
          setJobs((old) => ({
            ...old,
            interested: [...old.interested, job],
          }));
        })
        .catch((err) => console.log(err));
      return;
    }
  };

  const [jobFormState, setJobFormState] = useState<{
    open: boolean;
    mode: "create" | "edit";
    job: Job | null;
  }>({
    open: false,
    mode: "create",
    job: null,
  });

  const onJobEdit = (j: Job) => {};

  const onJobCreate = (j: Job) => {
    if (auth !== null)
      if (auth.user !== null)
        postJob(j, auth.user.id)
          .then((newJ) => {
            if (newJ !== null) {
              setJobFormState({ open: false, mode: "create", job: null });
              setJobs((o) => {
                setCurrJobs([...o.user, newJ]);
                return { ...o, all: [...o.all, newJ], user: [...o.user, newJ] };
              });
            }
          })
          .catch((err) => console.error(err));
  };

  const onJobApply = (j: Job) => {
    fetch(serverURI + "/jobs/apply", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        user_id: auth?.user?.id,
        job_id: j.id,
      }),
    })
      .then(() => {
        setJobs((old) => ({
          ...old,
          applied: [...old.applied, j],
        }));
      })
      .catch((err) => console.log(err));
  };

  return (
    <AuthRoute>
      <Navbar />
      <main className="flex justify-center py-10 px-44">
        <div className="px-7 border-r space-y-1 flex flex-col">
          {sideBar.map((s, i) => (
            <button
              key={s.type}
              className={classNames(
                sideBar[curPage].type === sideBar[i].type
                  ? " bg-gray-200 text-purple-800"
                  : "text-gray-800 hover:bg-gray-100",
                "text-lg rounded-md px-2 py-3 cursor-pointer focus:outline-none"
              )}
              onClick={() => changePage(i)}
            >
              {s.name}
            </button>
          ))}
        </div>
        <div className="pl-20 flex-1">
          <div className="flex justify-between">
            <h1 className="text-4xl font-semibold">{sideBar[curPage].name}</h1>
            <button
              className={classNames(
                sideBar[curPage].type === SideBarType.CREATE ? "" : "invisible",
                "btn bg-purple-800 text-white rounded-md py-2 px-3 hover:bg-purple-900 flex justify-center items-center"
              )}
              onClick={() =>
                setJobFormState({ open: true, mode: "create", job: null })
              }
            >
              <div className="mr-2">Create new</div>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                viewBox="0 0 16 16"
              >
                <path d="M8 0a1 1 0 0 1 1 1v6h6a1 1 0 1 1 0 2H9v6a1 1 0 1 1-2 0V9H1a1 1 0 0 1 0-2h6V1a1 1 0 0 1 1-1z" />
              </svg>
            </button>
          </div>
          <hr className="mt-2 mb-4 border shadow-md" />
          <div className="space-y-1">
            {currJobs.map((j, i) => {
              return (
                <>
                  {i !== 0 ? <hr key={j.id} /> : null}
                  <JobListItem
                    key={i !== 0 ? undefined : j.id}
                    currJob={j}
                    jobs={jobs}
                    currPage={sideBar[curPage]}
                    onJobButtonPress={onJobButtonPress}
                    onJobApply={onJobApply}
                  />
                </>
              );
            })}
          </div>
        </div>
      </main>

      <JobCreationForm
        {...{ jobFormState, setJobFormState, onJobEdit, onJobCreate }}
      />
    </AuthRoute>
  );
}
