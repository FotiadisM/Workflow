import { Job } from "@/src/types/job";
import { serverURI } from "./url";

export const fetchJobs = async (): Promise<Job[] | null> => {
  const res = await fetch(serverURI + "/jobs/");

  if (!res.ok) {
    const text = await res.text();
    console.log("ERROR fetching jobs:", text);
    return null;
  }

  const data = await res.json();

  let jobs: Job[] = [];

  data.jobs.forEach((j: any) => {
    jobs.push({
      id: j.id,
      user_id: j.user_id,
      title: j.title,
      type: j.type,
      location: j.location,
      company: { company_name: j.company },
      description: j.description,
      salary: { min: j.min_salary, max: j.max_salary },
      skills: j.skills,
      interested: j.interested,
      applied: j.applied,
      created: j.created,
    });
  });

  return jobs;
};

export const postJob = async (j: Job, user_id: string): Promise<Job | null> => {
  const res = await fetch(serverURI + "/jobs/", {
    method: "POST",
    headers: {},
    body: JSON.stringify({
      user_id: user_id,
      title: j.title,
      type: j.type,
      location: j.location,
      company: j.company.company_name,
      min_salary: j.salary.min,
      max_salary: j.salary.max,
      description: j.description,
      skills: j.skills,
    }),
  });

  if (!res.ok) {
    const text = await res.text();
    console.log("ERROR posting job:", text);
    return null;
  }

  const data = await res.json();

  j.id = data.id;
  j.created = data.created;

  return j;
};
